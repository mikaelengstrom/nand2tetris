package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFilename := os.Args[1]
	outputFilename := inputFilename[:len(inputFilename) - 3] + ".asm"

	infile := openFile(inputFilename)
	defer infile.Close()

	outfile := createFile(outputFilename)
	defer outfile.Close()

	scanner := bufio.NewScanner(bufio.NewReader(infile))

	parser := NewParser(scanner)

	for {
		instruction := parser.NextInstruction()
		if instruction.command == CommandTerminate {
			break
		}

		asmInstructions := translate(instruction)
		for _, row := range asmInstructions {
			fmt.Println(row)
			outfile.WriteString(row + "\n")
		}

	}
}


func openFile(filename string) *os.File{
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}

func createFile(filename string) *os.File{
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	return file
}
