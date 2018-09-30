import html


class XMLWriter:
    out_file = None
    indent_count = 0
    indent = ""

    def __init__(self, out_file):
        self.out_file = out_file
        self.indent = ""

    def open_tag(self, name):
        self.out_file.write('{}<{}>\n'.format(self.indent, name))
        self._indent_increase()

    def close_tag(self, name):
        self._indent_decrease()
        self.out_file.write('{}</{}>\n'.format(self.indent, name))

    def write_token(self, token):
        self.out_file.write('<{}> {} </{}>\n'.format(
            token.type, html.escape(token.token), token.type))

    def write_node(self, node):
        if type(node.value) == list:
            self.open_tag(node.name)
            for n in node.value:
                self.write_node(n)
            self.close_tag(node.name)
        else:
            self.out_file.write('{}<{}> {} </{}>\n'.format(
                self.indent, node.name, html.escape(node.value), node.name))

    def _indent_decrease(self):
        self.indent = self.indent[:-2]

    def _indent_increase(self):
        self.indent += "  "

