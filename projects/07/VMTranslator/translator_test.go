/**
 * This module translates Instructions to machine code.
 */

package main

import (
	"testing"
)

func TestTranslateCommandPushConstant(t *testing.T) {
	instruction := NewInstruction(CommandPush, "constant", 2)

	result := translate(instruction)

	expected := []string {
		"// " + instruction.String(),
		"@2",
		"D=A",
		"@SP",
		"A=M",
		"M=D",
		"@SP",
		"M=M+1",
	}

	for i, val := range result {
		assertEqual(t, val, expected[i], "")
	}

}

func TestTranslateCommandACAdd(t *testing.T) {
	instruction := NewInstruction(CommandArithmetic, string(ACAdd), -1)

	result := translate(instruction)

	expected := []string {
		"// " + instruction.String(),
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

	for i, val := range result {
		assertEqual(t, val, expected[i], "")
	}

}

func TestTranslateCommandACEq(t *testing.T) {
	instruction := NewInstruction(CommandArithmetic, string(ACEq), -1)

	result := translate(instruction)

	expected := []string {
		"// " + instruction.String(),
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

	for i, val := range result {
		assertEqual(t, val, expected[i], "")
	}

}

func TestTranslateCommandACLt(t *testing.T) {
	instruction := NewInstruction(CommandArithmetic, string(ACLt), -1)

	result := translate(instruction)

	expected := []string {
		"// " + instruction.String(),
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

	for i, val := range result {
		assertEqual(t, val, expected[i], "")
	}

}
