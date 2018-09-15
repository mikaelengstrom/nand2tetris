package main

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
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
	var builder strings.Builder

	state = initialState

	instructionLength := len(instruction)
	for i, char := range instruction {
		switch char {
		case '/':
			if state == Taking {
				tokens = append(tokens, builder.String())
				builder.Reset()
			}

			if state == LookingForCommentEnd {
				prevChar := instruction[i-1]
				if prevChar == '*' {
					state = LookingForContext
				}

				continue
			}

			if i == instructionLength - 1 {
				return tokens, state, errors.New("syntax error, stray ending slash")
			} else {
				nextChar := instruction[i+1]
				if nextChar == '/' {
					state = Exiting
				} else if nextChar == '*' {
					state = LookingForCommentEnd
				}
			}
			break


		case ' ':
			if state == Taking {
				tokens = append(tokens, builder.String())
				builder.Reset()

				state = LookingForContext
			}
			break




		default:
			if state == LookingForContext {
				state = Taking
			}

			if state == Taking {
				builder.WriteRune(char)
			}

			if i == instructionLength - 1 {
				tokens = append(tokens, builder.String())
			}
			break
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
	command := CommandDoesNotExist
	arg1 := ""
	arg2 := -1

	ac := stringToArithmeticCommand(tokens[0])
	if ac != ACUnknownCommand {
		command = CommandArithmetic
		arg1 = tokens[0]
	}

	ma := stringToMemoryAccessCommand(tokens[0])
	if ma != MAUnknownCommand {
		i, e := strconv.Atoi(tokens[2])
		if e != nil {
			panic(e)
		}

		arg1 = tokens[1]
		arg2 = i

		if ma == MAPop {
			command = CommandPop
		} else {
			command = CommandPush
		}
	}


	return NewInstruction(command, arg1, arg2)
}
