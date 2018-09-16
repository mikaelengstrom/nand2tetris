package main

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	fileNameOrDirectory := os.Args[1]

	file, err := os.Stat(fileNameOrDirectory)
	if err != nil {
		panic(err)
	}

	inFiles := []string{fileNameOrDirectory}

	re := regexp.MustCompile("\\..+$")
	outputFilename := re.ReplaceAllString(fileNameOrDirectory, "") + ".asm"
	if file.Mode().IsDir() {
		inFiles, err = filepath.Glob(filepath.Join(fileNameOrDirectory, "*.vm"))
		if err != nil {
			panic(err)
		}

		dirName := filepath.Base(fileNameOrDirectory)
		outputFilename = filepath.Join(fileNameOrDirectory, dirName + ".asm")
	}


	outfile := createFile(outputFilename)
	defer outfile.Close()

	var instructionIndex int
	for _, infilePath := range inFiles {
		infile := openFile(infilePath)

		_, fileName := filepath.Split(infilePath)
		codeWriterPrefix := re.ReplaceAllString(fileName, "")

		scanner := bufio.NewScanner(bufio.NewReader(infile))
		parser := NewParser(scanner)

		for {
			instruction := parser.NextInstruction()
			if instruction.command == CommandTerminate {
				break
			}

			instructionIndex++

			asmInstructions := translate(instruction, codeWriterPrefix)
			for _, row := range asmInstructions {
				row = insertJumpIndexes(row, instructionIndex)
				outfile.WriteString(row + "\n")
			}

		}

		infile.Close()
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
