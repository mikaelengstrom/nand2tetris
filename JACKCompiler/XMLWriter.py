from functools import reduce
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
        #self._indent_increase()

    def close_tag(self, name):
        #self._indent_decrease()
        self.out_file.write('{}</{}>\n'.format(self.indent, name))

    def write_token(self, token):
        self.out_file.write('{}<{}> {} </{}>\n'.format(
            self.indent, token.type, html.escape(token.token), token.type))

    def _indent_decrease(self):
        self.indent = self.indent[:-4]

    def _indent_increase(self):
        self.indent += "    "

