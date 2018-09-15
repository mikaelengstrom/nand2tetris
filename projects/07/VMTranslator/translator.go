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

