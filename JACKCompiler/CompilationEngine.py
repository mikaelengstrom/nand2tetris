from collections import namedtuple
from Tokenizer import TokenizerReachedEndOfFileException, Token

SyntaxNode = namedtuple('SyntaxNode', ['name', 'value'])

SyntaxNode.from_token = lambda t: SyntaxNode(name=t.type, value=t.token)


class CompilationEngine:
    def __init__(self, tokenizer):
        self.tokenizer = tokenizer

    def compile(self):
        token = self.tokenizer.advance()
        if token.type != 'keyword' and token.name != 'class':
            raise Exception("JACK programs must start with a keyword class")

        return SyntaxNode(name='class', value=self.compile_class())

    def compile_class(self):
        _r = [SyntaxNode(name='keyword', value='class')]

        token = self.tokenizer.advance()
        _assert_identifier(token)

        _r.append(SyntaxNode(name='identifier', value=token.token))

        token = self.tokenizer.advance()
        _assert_symbol('{', token)
        _r.append(SyntaxNode.from_token(token))

        while True:
            try:
                token = self.tokenizer.peak_next()
                if token.token == '}':
                    _r.append(SyntaxNode.from_token(self.tokenizer.advance()))
                    return _r
                elif token.token in ['field', 'static']:
                    _r.append(self.compile_class_var_dec())
                elif token.token in ['constructor', 'function', 'method']:
                    _r.append(self.compile_subroutine_dec())
                else:
                    raise Exception("Unexpected token '{}'".format(token.token))

            except TokenizerReachedEndOfFileException:
                raise Exception("Expected '}' for class")

    def compile_class_var_dec(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        token = self.tokenizer.advance()
        _assert_type(token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_identifier(token)
        _r.append(SyntaxNode.from_token(token))

        while True:
            last_token = token
            token = self.tokenizer.advance()
            _r.append(SyntaxNode.from_token(token))

            if last_token.type == Token.TYPE_IDENTIFIER:
                if token.token not in [',', ';']:
                    raise Exception('Expected "," or ";"')
                if token.token == ';':
                    return SyntaxNode(name="classVarDec", value=_r)
            else:
                _assert_identifier(token)

    def compile_subroutine_dec(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        token = self.tokenizer.advance()
        _assert_type(token)
        _r.append(SyntaxNode(name=token.type, value=token.token))

        token = self.tokenizer.advance()
        _assert_identifier(token)
        _r.append(SyntaxNode(name=token.type, value=token.token))

        token = self.tokenizer.advance()
        _assert_symbol('(', token)
        _r.append(SyntaxNode(name=token.type, value=token.token))

        _r.append(self.compile_parameter_list())

        token = self.tokenizer.advance()
        _assert_symbol(')', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_subroutine_body())

        return SyntaxNode(name='subroutineDec', value=_r)

    def compile_parameter_list(self):
        _r = []

        while True:
            next_token = self.tokenizer.peak_next()
            if next_token.token == ')':
                return SyntaxNode(name="parameterList", value=_r)
            elif next_token.token == ',':
                _r.append(SyntaxNode.from_token(self.tokenizer.advance()))
                continue

            _r += self.compile_parameter()

    def compile_parameter(self):
        _r = []

        token = self.tokenizer.advance()
        _assert_type(token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_identifier(token)
        _r.append(SyntaxNode.from_token(token))

        return _r

    def compile_subroutine_body(self):
        _r = []

        token = self.tokenizer.advance()
        _assert_symbol('{', token)
        _r.append(SyntaxNode.from_token(token))

        while self.tokenizer.peak_next().token == 'var':
            _r.append(self.compile_var_dec())

        _r.append(self.compile_statements())

        token = self.tokenizer.advance()
        _assert_symbol('}', token)
        _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name="subroutineBody", value=_r)

    def compile_var_dec(self):
        _r = []

        token = self.tokenizer.advance()
        _assert_token('var', token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_type(token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_identifier(token)
        _r.append(SyntaxNode.from_token(token))

        while self.tokenizer.peak_next().token != ';':
            _r.append(SyntaxNode.from_token(self.tokenizer.advance()))

        token = self.tokenizer.advance()
        _assert_symbol(';', token)
        _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name="varDec", value=_r)

    def compile_statements(self):
        _r = []
        while True:
            token = self.tokenizer.peak_next()

            if token.token == '}':
                return SyntaxNode(name="statements", value=_r)

            elif token.token == 'return':
                _r.append(self.compile_return_statement())

            elif token.token == 'let':
                _r.append(self.compile_let_statement())

            elif token.token == 'do':
                _r.append(self.compile_do_statement())

            elif token.token == 'if':
                _r.append(self.compile_if_statement())

            elif token.token == 'while':
                _r.append(self.compile_while_statement())

            else:
                raise Exception('subroutine: NOTIMPLEMENTED: {}'.format(token))

    def compile_return_statement(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        if self.tokenizer.peak_next().token == ';':
            _r.append(SyntaxNode.from_token(self.tokenizer.advance()))
        else:
            _r.append(self.compile_expression())

            token = self.tokenizer.advance()
            _assert_symbol(';', token)
            _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='returnStatement', value=_r)

    def compile_expression(self):
        _r = [self.compile_term()]

        while self._is_op(self.tokenizer.peak_next()):
            _r.append(SyntaxNode.from_token(self.tokenizer.advance()))
            _r.append(self.compile_term())

        return SyntaxNode(name='expression', value=_r)

    def compile_term(self):
        _r = []
        token = self.tokenizer.advance()
        next_token = self.tokenizer.peak_next()

        while token.token == '.' or next_token.token == '.':
            _r.append(SyntaxNode.from_token(token))
            token = self.tokenizer.advance()
            next_token = self.tokenizer.peak_next()

        _r.append(SyntaxNode.from_token(token))

        if self._is_unary_op(token):
            _r.append(self.compile_term())

        elif token.token == '(':
            _r.append(self.compile_expression())

            token = self.tokenizer.advance()
            _assert_symbol(')', token)
            _r.append(SyntaxNode.from_token(token))

        elif next_token.token == '(':
            _r += self.compile_subroutine_call()

        elif next_token.token == '[':
            _r.append(SyntaxNode.from_token(self.tokenizer.advance()))

            _r.append(self.compile_expression())

            token = self.tokenizer.advance()
            _assert_symbol(']', token)
            _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='term', value=_r)

    def _is_op(self, token):
        return token.token in ['+', '-', '*', '/', '&', '|', '<', '>', '=']

    def _is_unary_op(self, token):
        return token.token in ['-', '~']

    def compile_let_statement(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        token = self.tokenizer.advance()
        _assert_identifier(token)
        _r.append(SyntaxNode.from_token(token))

        if self.tokenizer.peak_next().token == '[':
            token = self.tokenizer.advance()
            _assert_symbol('[', token)
            _r.append(SyntaxNode.from_token(token))

            _r.append(self.compile_expression())

            token = self.tokenizer.advance()
            _assert_symbol(']', token)
            _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_symbol('=', token)
        _r.append(SyntaxNode.from_token(token))

        if self.tokenizer.peak_next().token != ';':
            _r.append(self.compile_expression())

        token = self.tokenizer.advance()
        _assert_symbol(';', token)
        _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='letStatement', value=_r)

    def compile_do_statement(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        _r += self.compile_subroutine_call()

        token = self.tokenizer.advance()
        _assert_symbol(';', token)
        _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='doStatement', value=_r)

    def compile_subroutine_call(self):
        _r = []

        while self.tokenizer.peak_next().token != '(':
            token = self.tokenizer.advance()
            _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_symbol('(', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_expression_list())

        token = self.tokenizer.advance()
        _assert_symbol(')', token)
        _r.append(SyntaxNode.from_token(token))

        return _r

    def compile_expression_list(self):
        _r = []
        while self.tokenizer.peak_next().token != ')':
            _r.append(self.compile_expression())
            if self.tokenizer.peak_next().token == ',':
                _r.append(SyntaxNode.from_token(self.tokenizer.advance()))

        return SyntaxNode(name='expressionList', value=_r)

    def compile_if_statement(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        token = self.tokenizer.advance()
        _assert_symbol('(', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_expression())

        token = self.tokenizer.advance()
        _assert_symbol(')', token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_symbol('{', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_statements())

        token = self.tokenizer.advance()
        _assert_symbol('}', token)
        _r.append(SyntaxNode.from_token(token))

        if self.tokenizer.peak_next().token == 'else':
            _r.append(SyntaxNode.from_token(self.tokenizer.advance()))

            token = self.tokenizer.advance()
            _assert_symbol('{', token)
            _r.append(SyntaxNode.from_token(token))

            _r.append(self.compile_statements())

            token = self.tokenizer.advance()
            _assert_symbol('}', token)
            _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='ifStatement', value=_r)

    def compile_while_statement(self):
        _r = [SyntaxNode.from_token(self.tokenizer.advance())]

        token = self.tokenizer.advance()
        _assert_symbol('(', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_expression())

        token = self.tokenizer.advance()
        _assert_symbol(')', token)
        _r.append(SyntaxNode.from_token(token))

        token = self.tokenizer.advance()
        _assert_symbol('{', token)
        _r.append(SyntaxNode.from_token(token))

        _r.append(self.compile_statements())

        token = self.tokenizer.advance()
        _assert_symbol('}', token)
        _r.append(SyntaxNode.from_token(token))

        return SyntaxNode(name='whileStatement', value=_r)


def _assert_type(token):
    if token.type not in [Token.TYPE_KEYWORD, Token.TYPE_IDENTIFIER]:
        return Exception("Expected type declartion, got '{}'".format(token))


def _assert_identifier(token):
    if token.type != Token.TYPE_IDENTIFIER:
        raise Exception("Expected identifier, got '{}' class identifier".format(token))


def _assert_symbol(symbol, token):
    if token.type != Token.TYPE_SYMBOL or token.token != symbol:
        raise Exception("Expected symbol '{}', got '{}'".format(symbol, token))


def _assert_token(token_name, token):
    if token.token != token_name:
        raise Exception("Expected token '{}', got '{}'".format(token_name, token))
