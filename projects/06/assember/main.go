package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open in/out-files and start scanning...
	var fileName = os.Args[1]

	inFile, err := os.Open(fileName)
	assertNil(err)
	defer inFile.Close()

	var outFileName = os.Args[1] + ".hack"
	outFile, err := os.Create(outFileName)
	assertNil(err)
	defer outFile.Close()

	scanner := bufio.NewScanner(bufio.NewReader(inFile))

	// Iterate through the file and push instructions to "rows" according to
	// the following rules:
	//
	//  - Remove comments and trim leading and trailing whitespace
	//  - If line is blank - skip it
	//  - If a (LABEL) hits, store it in the symboltable, then skip the line
	i := 0
	var rows []string
	st := NewSymbolTable()

	for scanner.Scan() {
		line := stripCommentsAndWhitespace(scanner.Text())

		if len(line) == 0 {
			continue
		}

		// Capture labels. (LABEL). Remove the ending parenthesis
		if line[0] == '(' {
			st.AddLabel(line[1:len(line) - 1], i)
			continue
		}

		rows = append(rows, line)
		i++
	}

	// Iterate the file again, parse the instruction and write to outfile
	// We do this process in two iterations because (LABELS) might be declared
	// after their usage according to the HACK-language specifications
	for _, instruction := range rows {
		outFile.WriteString(parseInstruction(instruction, st) + "\n")
	}

}

func assertNil(e error) {
	if e != nil {
		panic(e)
	}
}

func stripCommentsAndWhitespace(str string) string {
	stripped := stripComments(str)
	return strings.TrimSpace(stripped)
}

func stripComments(str string) string {
	type ParserState int
	const (
		StartTake ParserState = iota
		StopTake
		DropRest
	)

	commentSymbols := map[string] ParserState {
		"//": DropRest,
		"/*": StopTake,
		"*/": StartTake,
	}

	takeFromPos := 0

	var sb strings.Builder

	strLength := len(str)
	for pos, char := range str {

		token := ""
		if pos + 1 < strLength {
			token = string(char) + string(str[pos + 1])
		}

		match, ok := commentSymbols[token]; if ok {
			switch match {
			case StartTake:
				takeFromPos = pos + 2
				break
			case StopTake:
				takeFromPos = 99999
				break
			case DropRest:
				return sb.String()
			default:
				panic("Should never happen")
			}
		}

		if takeFromPos <= pos {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

func parseInstruction(instruction string, st *SymbolTable) string {
	// In HACK-language, instructions starting with @ is A-instructions, rest is C
	if instruction[0] == '@' {
		return parseAInstruction(instruction, st)
	} else {
		return parseCInstruction(instruction)
	}

	return instruction
}

func parseAInstruction(instruction string, st *SymbolTable) string {
	instruction, err := st.MaybeReplaceSymbol(instruction)
	assertNil(err)

	binaryString := intStringToBinaryString(instruction[1:])

	zeroFill := strings.Repeat("0", 16 - len(binaryString))
	return zeroFill + binaryString
}

func intStringToBinaryString(integerAsString string) string {
	i, err := strconv.ParseInt(integerAsString, 10, 64)
	assertNil(err)

	return strconv.FormatInt(i, 2)
}

func parseCInstruction(str string) string {
	destBits := map[string] string{
		"": "000",
		"null": "000",
		"M":    "001",
		"D":    "010",
		"MD":   "011",
		"A":    "100",
		"AM":   "101",
		"AD":   "110",
		"AMD":  "111",
	}

	jumpBits := map[string] string{
		"": "000",
		"null": "000",
		"JGT":  "001",
		"JEQ":  "010",
		"JGE":  "011",
		"JLT":  "100",
		"JNE":  "101",
		"JLE":  "110",
		"JMP":  "111",
	}

	compBits := map[string] string{
		"0":   "0101010",
		"1":   "0111111",
		"-1":  "0111010",
		"D":   "0001100",
		"A":   "0110000",
		"!D":  "0001101",
		"!A":  "0110001",
		"-D":  "0001111",
		"-A":  "0110011",
		"D+1": "0011111",
		"A+1": "0110111",
		"D-1": "0001110",
		"A-1": "0110010",
		"D+A": "0000010",
		"D-A": "0010011",
		"A-D": "0000111",
		"D&A": "0000000",
		"D|A": "0010101",

		"M":   "1110000",
		"!M":  "1110001",
		"-M":  "1110011",
		"M+1": "1110111",
		"M-1": "1110010",
		"D+M": "1000010",
		"D-M": "1010011",
		"M-D": "1000111",
		"D&M": "1000000",
		"D|M": "1010101",
	}

	dest := ""
	comp := ""
	jump := ""

	isJumpInstruction := false

	var sb strings.Builder

	for _, char := range str {
		switch char {
		case ' ':
			continue
		case '=':
			dest = strings.TrimSpace(sb.String())
			sb.Reset()
			continue
			break
		case ';':
			comp = strings.TrimSpace(sb.String())
			isJumpInstruction = true
			sb.Reset()
			continue
			break
		}

		sb.WriteRune(char)
	}

	if isJumpInstruction {
		jump = strings.TrimSpace(sb.String())
	} else {
		comp = strings.TrimSpace(sb.String())
	}

	header := "111"
	return header + compBits[comp] + destBits[dest] + jumpBits[jump]
}



type SymbolTable struct  {
	table map[string]int
	nextSlot int
}

func (st *SymbolTable) GetOrCreate(key string) int {
	_, ok := st.table[key]; if ok {
		return st.table[key]
	}
	st.table[key] = st.nextSlot
	st.nextSlot++

	return st.table[key]
}

func (st *SymbolTable) AddLabel(key string, row int) {
	st.table[key] = row
}

func (st *SymbolTable) MaybeReplaceSymbol(symbol string) (string, error){
	if symbol[0] != '@' {
		return "", errors.New(symbol + " is not an A instruction")
	}

	var address int

	// If value already exists in SymbolTable (eg. @LABEL or builtin, like @SCREEN) - use that one
	_, ok := st.table[symbol[1:]]; if ok {
		address = st.table[symbol[1:]]
	} else {
		// Check if value is a number, use that
		i, err := strconv.Atoi(symbol[1:len(symbol)])
		if err == nil {
			address = i
		} else {
			// For strings we fetch from symboltable
			address = st.GetOrCreate(symbol[1:])
		}
	}

	return "@" + strconv.Itoa(address), nil
}

func NewSymbolTable () *SymbolTable {
	st := &SymbolTable{
		table: map[string]int{
			// Builtin registers
			"R0": 0,
			"R1": 1,
			"R2": 2,
			"R3": 3,
			"R4": 4,
			"R5": 5,
			"R6": 6,
			"R7": 7,
			"R8": 8,
			"R9": 9,
			"R10": 10,
			"R11": 11,
			"R12": 12,
			"R13": 13,
			"R14": 14,
			"R15": 15,

			// Peripherals
			"SCREEN": 16384,
			"KBD": 24576,

			// Other builtins
			"SP": 0,
			"LCL": 1,
			"ARG": 2,
			"THIS": 3,
			"THAT": 4,
			"WRITE": 18,
			"END": 22,
		},
		nextSlot: 16,
	}

	return st
}
