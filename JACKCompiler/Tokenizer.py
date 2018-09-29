import os

KEYWORDS = [
    'class', 'constructor', 'function', 'method', 'field', 'static', 'var', 'int', 'char', 'boolean',
    'void', 'true', 'false', 'null', 'this', 'let', 'do', 'if', 'else', 'while', 'return'
]
SYMBOLS = ['{', '}', '(', ')', '[', ']', '.', ',', ';', '+', '-', '*', '/', '&', '|', '<', '>', '=', '~']


class Token:
    token = None
    type = None

    TYPE_SYMBOL = 'symbol'
    TYPE_KEYWORD = 'keyword'
    TYPE_INTEGER_CONSTANT = 'integerConstant'
    TYPE_STRING_CONSTANT = 'stringConstant'
    TYPE_IDENTIFIER = 'identifier'

    def __init__(self, token, type=None):
        self.token = token

        # Set type
        if type:
            self.type = type
        elif _is_symbol(self.token):
            self.type = Token.TYPE_SYMBOL
        elif _is_keyword(self.token):
            self.type = Token.TYPE_KEYWORD
        elif _is_integer_constant(self.token):
            self.type = Token.TYPE_INTEGER_CONSTANT
        else:
            self.type = Token.TYPE_IDENTIFIER

    def __str__(self):
        return "Token({}): {}".format(self.type, self.token)


class TokenizerReachedEndOfFileException(Exception):
    pass


class Tokenizer:
    is_done = False
    last_token = None

    def __init__(self, file_stream_in):
        self.in_file = file_stream_in
        self.last_token = self._get_next_token()

    def is_done(self):
        return self._peak_next_char() == ''

    def dump(self, x):
        pass

    def advance(self):
        _return = self.last_token
        while True:
            self.last_token = self._get_next_token()
            return _return

    def _get_next_token(self):
        head = self._take_to_next_char()

        # If head is the beginning of comment we drop the rest of the file and return recursive result
        if head == '/':
            slash_plus_next_char = ''.join([head, self._peak_next_char()])
            if _is_multi_line_comment(slash_plus_next_char):
                self._drop_to_comment_end()
                return self._get_next_token()

            elif _is_single_line_comment(slash_plus_next_char):
                self._drop_row()
                return self._get_next_token()

        # If head is the beginning of a string we take till the string end
        if head == '"':
            return Token(self._take_to_string_end(), Token.TYPE_STRING_CONSTANT)

        # If it is a symbol, it is a token in itself and we
        # therefore need no tail
        if _is_symbol(head):
            return Token(head)

        tail = self._take_to_next_white_space_or_symbol()

        return Token(''.join([head] + tail))

    def _take_to_next_char(self):
        while True:
            x = self._read()
            if not x.isspace():
                return x

    def _take_to_next_white_space_or_symbol(self):
        _return = []
        while True:
            x = self._read()
            if x.isspace():
                return _return
            elif _is_symbol(x):
                self._unread()
                return _return
            else:
                _return.append(x)

    def _drop_row(self):
        print("DEBUG: Dropping rest of line: {}".format(self.in_file.readline()))

    def _drop_to_comment_end(self):
        drop = []
        while True:
            char = self._read()
            drop.append(char)

            if char == '*':
                if self._peak_next_char() == '/':
                    drop.append(self._read())
                    print('DEBUG: Dropping {}'.format(''.join(drop)))
                    break

    def _take_to_string_end(self):
        take = []
        while True:
            char = self._read()
            if char == '"':
                return ''.join(take)

            take.append(char)

    def _read(self):
        s = self.in_file.read(1)
        if s == '':
            raise TokenizerReachedEndOfFileException("Tokenizer reached end of file")

        return s

    def _peak_next_char(self):
        x = self._read()
        self._unread()
        return x

    def _unread(self):
        return self.in_file.seek(self.in_file.tell() - 1, os.SEEK_SET)


def _is_single_line_comment(string):
    return string[:2] == "//"


def _is_multi_line_comment(string):
    return string[:2] == "/*"


def _is_symbol(x):
    return x in SYMBOLS


def _is_keyword(x):
    return x in KEYWORDS


def _is_integer_constant(x):
    return x.isdigit()
