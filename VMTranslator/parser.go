package main

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"unicode"
)

type ParserState int
const (
	LookingForContext ParserState = iota
	Taking
	LookingForCommentEnd
	Exiting
)

type Parser struct {
	state ParserState
	scanner *bufio.Scanner
}

func NewParser(scanner *bufio.Scanner) *Parser {
	parser := &Parser{
		state: LookingForContext,
		scanner: scanner,
	}

	return parser
}

func (p *Parser) NextInstruction() Instruction{
	for p.scanner.Scan() {
		tokens, state, err := generateTokens(p.scanner.Text(), p.state)
		if err != nil {
			panic(err)
		}

		p.state = state

		if len(tokens) == 0 {
			continue
		}

		return tokenize(tokens)
	}

	return Instruction{command:CommandTerminate}
}

/*
 * Takes an instruction and tokenize it
 * Eg. "push token 3 // whatever" -> ["push", "token", "3"]
 *
 * Since comments might be multiline, a ParserState need to
 * be supplied. You should pass the result from last run and
 * "LookingForContext" for the first run
 */
func generateTokens(instruction string, initialState ParserState) (tokens []string, state ParserState, err error) {
	buffer := new(bytes.Buffer)
	state = initialState

	instructionLength := len(instruction)
	for i, char := range instruction {
		if char == '/' {
			if state == Taking {
				token := string(buffer.Bytes())
				if len(token) > 0 {
					tokens = append(tokens, token)
				}
				buffer = new(bytes.Buffer)
			}

			if state == LookingForCommentEnd {
				prevChar := instruction[i-1]
				if prevChar == '*' {
					state = LookingForContext
				}

				continue
			}

			if i == instructionLength-1 {
				return tokens, state, errors.New("syntax error, stray ending slash")
			} else {
				nextChar := instruction[i+1]
				if nextChar == '/' {
					state = Exiting
				} else if nextChar == '*' {
					state = LookingForCommentEnd
				}
			}

		} else if unicode.IsSpace(char) {
			if state == Taking {
				tokens = append(tokens, string(buffer.Bytes()))
				buffer = new(bytes.Buffer)

				state = LookingForContext
			}
		} else {
			if state == LookingForContext {
				state = Taking
			}

			if state == Taking {
				buffer.WriteRune(char)
			}

			if i == instructionLength - 1 {
				tokens = append(tokens, string(buffer.Bytes()))
			}
		}

		if state == Exiting {
			break
		}
	}

	if state == Exiting {
		state = LookingForContext
	}

	return tokens, state, nil
}

func tokenize(tokens []string) Instruction{
	ac := stringToArithmeticCommand(tokens[0])
	if ac != ACUnknownCommand {
		return NewInstruction(CommandArithmetic, tokens[0], -1)
	}

	var command Command
	var arg1 string
	var arg2 int

	switch command = stringToCommand(tokens[0]); {
	default:
		if len(tokens) > 1 {
			arg1 = tokens[1]
		}

		if len(tokens) > 2 {
			arg2, _ = strconv.Atoi(tokens[2])
		}
	}

	return NewInstruction(command, arg1, arg2)
}
