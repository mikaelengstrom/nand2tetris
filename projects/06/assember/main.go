package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var fileName = os.Args[1]

	inFile, err := os.Open(fileName)
	assertNotNil(err)
	defer inFile.Close()

	var outFileName = os.Args[1] + ".hack"
	outFile, err := os.Create(outFileName)
	assertNotNil(err)
	defer outFile.Close()


	scanner := bufio.NewScanner(bufio.NewReader(inFile))

	var rows []string
	i := 0
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if len(line) == 0 {
			continue
		}

		rows = append(rows, parseInstruction(line))
		i++
	}

	for _, val := range rows {
		outFile.WriteString(val + "\n")
	}

}

func assertNotNil(e error) {
	if e != nil {
		panic(e)
	}
}

func parseLine(str string) string {
	stripped := stripComments(str)
	return strings.TrimSpace(stripped)
}

func stripComments(str string) string {
	const (
		START_TAKE = iota
		STOP_TAKE
		DROP_REST
	)

	commentSymbols := map[string] int {
		"//": DROP_REST,
		"/*": STOP_TAKE,
		"*/": START_TAKE,
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
			case START_TAKE:
				takeFromPos = pos + 2
				break
			case STOP_TAKE:
				takeFromPos = 99999
				break
			case DROP_REST:
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

func parseInstruction(str string) string {
	if str[0] == '@' {
		return parseAInstruction(str)
	} else {
		return parseCInstruction(str)
	}
	return str
}

func parseAInstruction(str string) string {
	i, err := strconv.ParseInt(str[1:], 10, 64)
	assertNotNil(err)

	binaryString := strconv.FormatInt(i, 2)
	zeroFill := strings.Repeat("0", 16 - len(binaryString))
	return zeroFill + binaryString
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


