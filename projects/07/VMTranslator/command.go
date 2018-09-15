package main

type Command int
const (
	CommandArithmetic Command = iota
	CommandPush
	CommandPop
	CommandLabel
	CommandGoTo
	CommandIf
	CommandFunction
	CommandReturn
	CommandCall
	CommandTerminate
	CommandDoesNotExist
)

func (c Command) String() string{
	switch c {
	case CommandArithmetic:
		return "arithmetic"
	case CommandPush:
		return "push"
	case CommandPop:
		return "pop"
	case CommandTerminate:
		return "terminate"
	default:
		return "unknown"
	}
}

type ArithmeticCommand string
const (
	ACAdd ArithmeticCommand = "add"
	ACSub ArithmeticCommand = "sub"
	ACNeg ArithmeticCommand = "neg"
	ACEq ArithmeticCommand = "eq"
	ACGt ArithmeticCommand = "gt"
	ACLt ArithmeticCommand = "lt"
	ACAnd ArithmeticCommand = "and"
	ACOr ArithmeticCommand = "or"
	ACNot ArithmeticCommand = "not"

	ACUnknownCommand ArithmeticCommand = "unknown"
)

func stringToArithmeticCommand(str string) ArithmeticCommand {
	acMap := map[string]ArithmeticCommand {
		"add": ACAdd,
		"sub": ACSub,
		"neg": ACNeg,
		"eq": ACEq,
		"gt": ACGt,
		"lt": ACLt,
		"and": ACAnd,
		"or": ACOr,
		"not": ACNot,
	}

	_, ok := acMap[str]; if ok {
		return acMap[str]
	}

	return ACUnknownCommand
}

type MemoryAccessCommand string
const (
	MAPop MemoryAccessCommand = "pop"
	MAPush MemoryAccessCommand = "push"

	MAUnknownCommand MemoryAccessCommand = "unknown"
)

func stringToMemoryAccessCommand(str string) MemoryAccessCommand {
	maMap := map[string]MemoryAccessCommand {
		"pop": MAPop,
		"push": MAPush,
	}

	_, ok := maMap[str]; if ok {
		return maMap[str]
	}

	return MAUnknownCommand
}