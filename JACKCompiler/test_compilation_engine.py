from io import BytesIO

from Tokenizer import Tokenizer
from CompilationEngine import CompilationEngine


def test_compile_empty_class():
    fake_file = BytesIO(b"""
    class X{
    }""")
    t = Tokenizer(fake_file)
    c = CompilationEngine(t)
    node = c.compile()
    assert(node.name == 'class')

    assert(node.value[0].name == 'keyword')
    assert(node.value[0].value == 'class')

    assert(node.value[1].name == 'identifier')
    assert(node.value[1].value == 'X')

    assert(node.value[2].name == 'symbol')
    assert(node.value[2].value == '{')

    assert(node.value[3].name == 'symbol')
    assert(node.value[3].value == '}')


def test_compile_class_with_var_description():
    fake_file = BytesIO(b"""
    class X{
        static String x1, x2;
        field String y;
    }""")
    t = Tokenizer(fake_file)
    c = CompilationEngine(t)
    node = c.compile()

    vd_1 = node.value[3]
    vd_2 = node.value[4]

    # Test static

    assert(vd_1.value[0].name == 'keyword')
    assert(vd_1.value[0].value == 'static')

    assert(vd_1.value[1].name == 'identifier')
    assert(vd_1.value[1].value == 'String')

    assert(vd_1.value[2].name == 'identifier')
    assert(vd_1.value[2].value == 'x1')

    assert(vd_1.value[3].name == 'symbol')
    assert(vd_1.value[3].value == ',')

    assert(vd_1.value[4].name == 'identifier')
    assert(vd_1.value[4].value == 'x2')

    # Test field
    assert(vd_2.value[0].name == 'keyword')
    assert(vd_2.value[0].value == 'field')

    assert(vd_2.value[1].name == 'identifier')
    assert(vd_2.value[1].value == 'String')

    assert(vd_2.value[2].name == 'identifier')
    assert(vd_2.value[2].value == 'y')

    assert(vd_2.value[3].name == 'symbol')
    assert(vd_2.value[3].value == ';')

    # Test that class closes properly
    assert(node.value[5].name == 'symbol')
    assert(node.value[5].value == '}')


def test_compile_subroutine_desc():
    fake_file = BytesIO(b"""
        {
            return this;
        }
    NEXTTOKEN
    """)
    t = Tokenizer(fake_file)
    c = CompilationEngine(t)

    body = c.compile_subroutine_body()
    assert body.name == 'subroutineBody'

    assert body.value[0].value == '{'
    statements = body.value[1].value
    assert body.value[2].value == '}'

    assert statements[0].name == 'returnStatement'
    assert statements[0].value[0].value == 'return'
    assert statements[0].value[1].name == 'expression'
    assert statements[0].value[2].value == ';'

    assert t.advance().token == 'NEXTTOKEN'


def test_compile_class_with_subroutine_description():
    fake_file = BytesIO(b"""
    class X{
        constructor X new(){
            return this;
        }
    
        method void new(){
            return this;
        }
    
        function void main() {
            return;
        }
    }""")
    t = Tokenizer(fake_file)
    c = CompilationEngine(t)
    node = c.compile()

    sd_1 = node.value[3].value
    sd_2 = node.value[4].value
    sd_3 = node.value[5].value

    # Test constructor
    assert(sd_1.value[0].name == 'keyword')
    assert(sd_1.value[0].value == 'constructor')

    assert(sd_1.value[1].name == 'identifier')
    assert(sd_1.value[1].value == 'X')

    assert(sd_1.value[2].name == 'identifier')
    assert(sd_1.value[2].value == 'new')

    assert(sd_1.value[3].name == 'symbol')
    assert(sd_1.value[3].value == '(')

    assert(sd_1.value[4].name == 'parameterList')
    assert(sd_1.value[4].value == [])

    assert(sd_1.value[5].name == 'symbol')
    assert(sd_1.value[5].value == ')')

    assert(sd_1.value[6].name == 'subroutineBody')
    assert(sd_1.value[6].value[0].value == '{')

    # Test method
    assert(sd_2.value[0].name == 'keyword')
    assert(sd_2.value[0].value == 'method')

    # Test function
    assert(sd_3.value[0].name == 'keyword')
    assert(sd_3.value[0].value == 'function')

    # Test that class closes properly
    assert(node.value[6].name == 'symbol')
    assert(node.value[6].value == '}')
