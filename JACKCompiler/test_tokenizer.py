from io import BytesIO

from Tokenizer import Tokenizer


def test_base():
    fake_file = BytesIO(b"method void dispose();")
    t = Tokenizer(fake_file)

    token = t.advance()
    assert(token.token == 'method')
    assert(token.type == 'keyword')

    token = t.advance()
    assert(token.token == 'void')
    assert(token.type == 'keyword')

    token = t.advance()
    assert(token.token == 'dispose')
    assert(token.type == 'identifier')

    token = t.advance()
    assert(token.token == '(')
    assert(token.type == 'symbol')

    token = t.advance()
    assert(token.token == ')')
    assert(token.type == 'symbol')

    token = t.advance()
    assert(token.token == ';')
    assert(token.type == 'symbol')


def test_issues__semicolon_missing_from_outfile():
    fake_file = BytesIO(b"""
      var char key;  // the key currently pressed by the user
      var boolean exit;
      let exit = false;
    """)

    t = Tokenizer(fake_file)
    assert(t.advance().token == 'var')
    assert(t.advance().token == 'char')
    assert(t.advance().token == 'key')
    assert(t.advance().token == ';')

    assert(t.advance().token == 'var')
    assert(t.advance().token == 'boolean')
    assert(t.advance().token == 'exit')
    assert(t.advance().token == ';')

    assert(t.advance().token == 'let')
    assert(t.advance().token == 'exit')
    assert(t.advance().token == '=')
    assert(t.advance().token == 'false')
    assert(t.advance().token == ';')

    fake_file = BytesIO(b"        return();")
    t = Tokenizer(fake_file)
    assert(t.advance().token == 'return')
    assert(t.advance().token == '(')
    assert(t.advance().token == ')')
    assert(t.advance().token == ';')
