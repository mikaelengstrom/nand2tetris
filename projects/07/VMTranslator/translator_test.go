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
		"D=A",
		"@SP",
		"M=D+1",
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
