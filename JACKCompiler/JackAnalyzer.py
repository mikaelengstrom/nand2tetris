import glob
import os
import sys

from CompilationEngine import CompilationEngine
from Tokenizer import Tokenizer, TokenizerReachedEndOfFileException
from XMLWriter import XMLWriter

in_file = sys.argv[1]


sources = glob.glob("./{}/*.jack".format(in_file)) if os.path.isdir(in_file) else [in_file]
for source in sources:
    base_name = source[:-len(".jack")]
    in_file = source
    tokenizer_outfile = "{}T.xml".format(base_name)
    compilation_engine_outfile = "{}.xml".format(base_name)

    with open(tokenizer_outfile, 'w') as tokenizer_file_out:
        tokenizer_xml_writer = XMLWriter(tokenizer_file_out)

        tokenizer_xml_writer.open_tag('tokens')

        with open(in_file, 'rb') as f_in:
            tokenizer = Tokenizer(f_in)

            while True:
                try:
                    tokenizer_xml_writer.write_token(tokenizer.advance())
                except TokenizerReachedEndOfFileException:
                    print('Reached end')
                    break

        tokenizer_xml_writer.close_tag('tokens')

    with open(compilation_engine_outfile, 'w') as ce_file_out:
        ce_xml_writer = XMLWriter(ce_file_out)

        with open(in_file, 'rb') as f_in:
            tokenizer = Tokenizer(f_in)
            ce = CompilationEngine(tokenizer)

            ce_xml_writer.write_node(ce.compile())
