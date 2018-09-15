/**
 * This module translates Instructions to machine code.
 */

package main

import "strconv"

func translatePush(instruction Instruction) []string {
	return []string {
		commentHeader(instruction),
		"@" + strconv.Itoa(instruction.arg2),
		"D=A",

		"@SP",
		"A=M",
		"M=D",

		"D=A",
		"@SP",
		"M=D+1",
	}
}

func translateArithmetic(instruction Instruction) []string {
	switch stringToArithmeticCommand(instruction.arg1) {
	case ACAdd:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M",

			"A=A-1",
			"D=M",
			"A=A-1",
			"M=D+M",

			"D=A",
			"@SP",
			"M=D+1",
		}

	case ACEq:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M",

			"A=A-1",
			"D=M",
			"A=A-1",
			"D=D-M;",

			"@JMP-eq",
			"D; JEQ",
			"D=0",
			"@JMP-end",
			"0; JMP",
			"(JMP-eq)",
			"D=-1",

			"(JMP-end)",
			"@SP",
			"A=M-1",
			"A=A-1",
			"M=D",

			"D=A+1",
			"@SP",
			"M=D",
		}

	case ACLt:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M",

			"A=A-1",
			"D=M",
			"A=A-1",
			"D=D-M;",

			"@JMP-gt",
			"D; JGT",
			"D=0",
			"@JMP-end",
			"0; JMP",
			"(JMP-gt)",
			"D=-1",

			"(JMP-end)",
			"@SP",
			"A=M-1",
			"A=A-1",
			"M=D",

			"D=A+1",
			"@SP",
			"M=D",
		}

	case ACGt:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M",

			"A=A-1",
			"D=M",
			"A=A-1",
			"D=D-M;",

			"@JMP-gt",
			"D; JLT",
			"D=0",
			"@JMP-end",
			"0; JMP",
			"(JMP-gt)",
			"D=-1",

			"(JMP-end)",
			"@SP",
			"A=M-1",
			"A=A-1",
			"M=D",

			"D=A+1",
			"@SP",
			"M=D",
		}

	case ACSub:
		return []string {
			commentHeader(instruction),
			"@SP",
			"AM=M-1",

			"D=M",
			"A=A-1",
			"M=M-D",
		}

	case ACAnd:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M",

			"A=A-1",
			"D=M",
			"A=A-1",
			"M=D&M",

			"D=A",
			"@SP",
			"M=D+1",
		}

	case ACOr:
		return []string {
			commentHeader(instruction),
			"@SP",
			"AM=M-1",

			"A=A-1",
			"D=M",
			"A=A+1",
			"D=D|M",
			"A=A-1",
			"M=D",
		}

	case ACNeg:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M-1",
			"M=!M",
			"M=M+1",
		}

	case ACNot:
		return []string {
			commentHeader(instruction),
			"@SP",
			"A=M-1",
			"M=!M",
		}
	}

	return []string{}
}


func translate (instruction Instruction) []string {
	switch instruction.command {
	case CommandPush:
		return translatePush(instruction)
	case CommandArithmetic:
		return translateArithmetic(instruction)
		break
	}

	return []string{}
}

func commentHeader(instruction Instruction) string {
	return "// " + instruction.String()
}


func insertJumpIndexes(row string, i int) string {
	if row[0] == '@' || row[0] == '(' {
		if len(row) > 3 && row[1:4] == "JMP" {
			return row[0:4] + strconv.Itoa(i) + row[4:]
		}
	}

	return row
}
