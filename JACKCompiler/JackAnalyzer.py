import glob
import os
import sys

from Tokenizer import Tokenizer, TokenizerReachedEndOfFileException
from XMLWriter import XMLWriter

in_file = sys.argv[1]


sources = glob.glob("./{}/*.jack".format(in_file)) if os.path.isdir(in_file) else [in_file]
for source in sources:
    base_name = source[:-len(".jack")]
    in_file = source
    tokenizer_outfile = "{}T.xml".format(base_name)

    with open(tokenizer_outfile, 'w') as f_out:
        xml_writer = XMLWriter(f_out)

        xml_writer.open_tag('tokens')

        with open(in_file, 'rb') as f_in:
            tokenizer = Tokenizer(f_in)

            while True:
                try:
                    xml_writer.write_token(tokenizer.advance())
                except TokenizerReachedEndOfFileException:
                    print('Reached end')
                    break

        xml_writer.close_tag('tokens')
