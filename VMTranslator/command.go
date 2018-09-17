package main

type Command int
const (
	CommandArithmetic Command = iota
	CommandPush
	CommandPop
	CommandLabel
	CommandGoTo
	CommandIfGoTo
	CommandFunction
	CommandReturn
	CommandCall
	CommandTerminate
	CommandDoesNotExist
)

func stringToCommand(str string) Command {
	cMap := map[string]Command {
		"arithmetic": CommandArithmetic,
		"push": CommandPush,
		"pop": CommandPop,
		"label": CommandLabel,
		"goto": CommandGoTo,
		"if-goto": CommandIfGoTo,
		"function": CommandFunction,
		"return": CommandReturn,
		"call": CommandCall,
		"terminate": CommandTerminate,
	}

	_, ok := cMap[str]; if ok {
		return cMap[str]
	}

	return CommandDoesNotExist
}

func (c Command) String() string{
	switch c {
	case CommandArithmetic:
		return "arithmetic"
	case CommandPush:
		return "push"
	case CommandPop:
		return "pop"
	case CommandLabel:
		return "label"
	case CommandGoTo:
		return "goto"
	case CommandIfGoTo:
		return "if-goto"
	case CommandFunction:
		return "function"
	case CommandCall:
		return "call"
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

type MemoryRegister string
const (
	MRConstant MemoryRegister = "constant"
	MRLocal MemoryRegister = "local"
	MRArgument MemoryRegister = "argument"
	MRThis MemoryRegister = "this"
	MRThat MemoryRegister = "that"
	MRTemp MemoryRegister = "temp"
	MRStatic MemoryRegister = "static"
	MRPointer MemoryRegister = "pointer"

	MRUnknownRegister MemoryRegister = "unknown"
)

func stringToMemoryRegister(str string) MemoryRegister {
	mrMap := map[string]MemoryRegister {
		"constant": MRConstant,
		"local": MRLocal,
		"argument": MRArgument,
		"this": MRThis,
		"that": MRThat,
		"temp": MRTemp,
		"static": MRStatic,
		"pointer": MRPointer,
	}

	_, ok := mrMap[str]; if ok {
		return mrMap[str]
	}

	return MRUnknownRegister
}
