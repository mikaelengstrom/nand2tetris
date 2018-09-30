import os
from io import BytesIO

KEYWORDS = [
    b'class', b'constructor', b'function', b'method', b'field',
    b'static', b'var', b'int', b'char', b'boolean',
    b'void', b'true', b'false', b'null', b'this',
    b'let', b'do', b'if', b'else', b'while', b'return'
]
SYMBOLS = [
    b'{', b'}', b'(', b')', b'[', b']',
    b'.', b',', b';', b'+', b'-', b'*',
    b'/', b'&', b'|', b'<', b'>', b'=',
    b'~'
]


class Token:
    token = None
    type = None

    TYPE_SYMBOL = 'symbol'
    TYPE_KEYWORD = 'keyword'
    TYPE_INTEGER_CONSTANT = 'integerConstant'
    TYPE_STRING_CONSTANT = 'stringConstant'
    TYPE_IDENTIFIER = 'identifier'

    def __init__(self, token, type=None):
        self.token = token.decode('utf-8')

        # Set type
        if type:
            self.type = type
        elif _is_symbol(token):
            self.type = Token.TYPE_SYMBOL
        elif _is_keyword(token):
            self.type = Token.TYPE_KEYWORD
        elif _is_integer_constant(token):
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

    def dump(self, x):
        pass

    def peak_next(self):
        _cursor = self.in_file.tell()
        _next = self._get_next_token()
        self.in_file.seek(_cursor, os.SEEK_SET)
        return _next

    def advance(self):
        while True:
            self.last_token = self._get_next_token()
            return self.last_token

    def _get_next_token(self):
        while True:
            head = self._take_to_next_char()

            # If head is the beginning of a string we take till the string end
            if head == b'"':
                return Token(self._take_to_string_end(), Token.TYPE_STRING_CONSTANT)

            # If head is the beginning of comment we drop the rest of the comment and fetch a new char
            if head == b'/':
                slash_plus_next_char = head + self._peak_next_char()
                if _is_multi_line_comment(slash_plus_next_char):
                    self._drop_to_comment_end()
                    continue

                elif _is_single_line_comment(slash_plus_next_char):
                    self._drop_row()
                    continue

            # If it is a symbol, it is a token in itself and we
            # therefore need no tail
            if _is_symbol(head):
                return Token(head)

            tail = self._take_to_next_white_space_or_symbol()

            token = head + tail
            return Token(token)

    def _take_to_next_char(self):
        while True:
            x = self._read()
            if not x.isspace():
                return x

    def _take_to_next_white_space_or_symbol(self):
        _return = BytesIO()
        while True:
            x = self._read()
            if x.isspace():
                _return.seek(0)
                return _return.read()
            elif _is_symbol(x):
                self._unread()
                _return.seek(0)
                return _return.read()
            else:
                _return.write(x)

    def _drop_row(self):
        self.in_file.readline()

    def _drop_to_comment_end(self):
        while True:
            char = self._read()

            if char == b'*':
                if self._peak_next_char() == b'/':
                    self._read()
                    break

    def _take_to_string_end(self):
        take = b""
        while True:
            char = self._read()
            if char == b'"':
                return take

            take += char

    def _read(self):
        s = self.in_file.read(1)
        if s == b'':
            raise TokenizerReachedEndOfFileException("Tokenizer reached end of file")

        return s

    def _peak_next_char(self):
        x = self._read()
        self._unread()
        return x

    def _unread(self):
        return self.in_file.seek(self.in_file.tell() - 1, os.SEEK_SET)


def _is_single_line_comment(string):
    return string[:2] == b"//"


def _is_multi_line_comment(string):
    return string[:2] == b"/*"


def _is_symbol(x):
    return x in SYMBOLS


def _is_keyword(x):
    return x in KEYWORDS


def _is_integer_constant(x):
    return x.isdigit()
