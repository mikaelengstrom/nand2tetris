package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputFilename := os.Args[1]

	infile := openFile(inputFilename)
	defer infile.Close()
	scanner := bufio.NewScanner(bufio.NewReader(infile))

	parser := NewParser(scanner)

	for {
		instruction := parser.NextInstruction()
		if instruction.command == CommandTerminate {
			break
		}

		fmt.Println(instruction)
	}
}


func openFile(filename string) *os.File{
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return file
}
