package main

import "fmt"

type Instruction struct {
	command Command
	arg1 string
	arg2 int
}

func (i Instruction) String() string {
	return fmt.Sprintf("command: %v, arg1: %s, arg2: %v", i.command, i.arg1, i.arg2)
}

func NewInstruction(command Command, arg1 string, arg2 int) Instruction{
	return Instruction{command: command, arg1:arg1, arg2:arg2}
}
