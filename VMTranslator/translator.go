/**
 * This module translates Instructions to machine code.
 */

package main

import (
	"strconv"
)


func translate (instruction Instruction, prefix string) []string {
	switch instruction.command {
	case CommandPush:
		return translatePush(instruction, prefix)
	case CommandPop:
		return translatePop(instruction, prefix)
	case CommandArithmetic:
		return translateArithmetic(instruction)
	case CommandLabel:
		return translateLabel(instruction, prefix)
	case CommandGoTo:
		return translateGoTo(instruction, prefix)
	case CommandIfGoTo:
		return translateIfGoTo(instruction, prefix)
	}

	return []string{
		commentHeader(instruction),
		"NOTIMPLEMENTED",
	}
}

func translatePush(instruction Instruction, staticPrefix string) []string {
	register := stringToMemoryRegister(instruction.arg1)
	switch register {
	case MRConstant:
		return []string {
			commentHeader(instruction),
			"@" + strconv.Itoa(instruction.arg2),
			"D=A",

			"@SP",
			"A=M",
			"M=D",

			"@SP",
			"M=M+1",
		}
	case MRLocal, MRArgument, MRThat, MRThis:
		baseAddressSymbols := map[MemoryRegister]string{
			MRLocal: "LCL",
			MRArgument : "ARG",
			MRThat: "THAT",
			MRThis: "THIS",
		}

		return []string {
			commentHeader(instruction),
			"@" + strconv.Itoa(instruction.arg2),
			"D=A",

			"@" + baseAddressSymbols[register],
			"A=D+M",
			"D=M",

			"@SP",
			"A=M",
			"M=D",

			"@SP",
			"M=M+1",
		}

	case MRPointer:
		destination := "THIS"
		if instruction.arg2 == 1 {
			destination = "THAT"
		}

		return []string {
			commentHeader(instruction),
			"@" + destination,
			"D=M",

			"@SP",
			"A=M",
			"M=D",

			"@SP",
			"M=M+1",
		}

	case MRTemp:
		return []string {
			commentHeader(instruction),
			"@R" + strconv.Itoa(instruction.arg2 + 5),
			"D=M",

			"@SP",
			"A=M",
			"M=D",

			"@SP",
			"M=M+1",
		}

	case MRStatic:
		return []string {
			commentHeader(instruction),
			"@" + staticPrefix + "." + strconv.Itoa(instruction.arg2),
			"D=M",

			"@SP",
			"A=M",
			"M=D",

			"@SP",
			"M=M+1",
		}
	}

	return []string{}
}

func translatePop(instruction Instruction, staticPrefix string) []string {
	register := stringToMemoryRegister(instruction.arg1)
	switch register {
	case MRConstant:
		// Pushing to constant does not make sense, hence output nada
		return []string {}

	case MRLocal, MRArgument, MRThat, MRThis:
		baseAddressSymbols := map[MemoryRegister]string{
			MRLocal: "LCL",
			MRArgument : "ARG",
			MRThat: "THAT",
			MRThis: "THIS",
		}

		return []string {
			// "pop local 3" – Move the top of the stack (constant) to local index 3
			commentHeader(instruction),

			// Set R14 to @LCL + 3
			"@" + strconv.Itoa(instruction.arg2),
			"D=A",
			"@" + baseAddressSymbols[register],
			"D=D+M",
			"@R14",
			"M=D",

			// decrease stack pointer
			"@SP",
			"AM=M-1",

			// Store M[@SP] -> M[R14]
			"D=M",
			"@R14",
			"A=M",
			"M=D",
		}

	case MRPointer:
		destination := "THIS"
		if instruction.arg2 == 1 {
			destination = "THAT"
		}

		return []string {
			commentHeader(instruction),
			// decrease stack pointer
			"@SP",
			"AM=M-1",

			// Store M[@SP] -> M[@destination]
			"D=M",
			"@" + destination,
			"M=D",
		}

	case MRTemp:
		return []string {
			commentHeader(instruction),
			// decrease stack pointer
			"@SP",
			"AM=M-1",

			// Store M[@SP] -> M[Temp[arg2]]
			"D=M",
			"@R" + strconv.Itoa(instruction.arg2 + 5),
			"M=D",
		}

	case MRStatic:
		return []string{
			commentHeader(instruction),
			// decrease stack pointer
			"@SP",
			"AM=M-1",

			// Store M[@SP] -> @staticvar
			"D=M",
			"@" + staticPrefix + "." + strconv.Itoa(instruction.arg2),
			"M=D",
		}
	}

	return []string{}
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

func translateLabel(instruction Instruction, lblPrefix string) []string {
	label := lblPrefix + "." + instruction.arg1
	return []string {
		commentHeader(instruction),
		"(" + label + ")",
	}
}

func translateGoTo(instruction Instruction, lblPrefix string) []string {
	label := lblPrefix + "." + instruction.arg1

	return []string {
		commentHeader(instruction),
		"@" + label,
		"0; JMP",
	}
}

func translateIfGoTo(instruction Instruction, lblPrefix string) []string {
	label := lblPrefix + "." + instruction.arg1

	return []string {
		commentHeader(instruction),
		"@SP",
		"AM=M-1",
		"D=M",
		"@" + label,
		"D; JGT",
	}
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
