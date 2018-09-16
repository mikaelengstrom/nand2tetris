package main

import "testing"

func TestGenerateTokens(t *testing.T) {
	tokens, _, err := generateTokens("// Start with comment", LookingForContext )
	assertEqual(t, len(tokens), 0, "")
	assertEqual(t, err, nil, "")

	tokens, _, err = generateTokens("", LookingForContext )
	assertEqual(t, len(tokens), 0, "")
	assertEqual(t, err, nil, "")

	tokens, _, err = generateTokens("  // hej", LookingForContext)
	assertEqual(t, len(tokens), 0, "")
	assertEqual(t, err, nil, "")


	_, _, err = generateTokens("push constant 2 /", LookingForContext)
	assertEqual(t, err == nil, false, "")

	tokens, _, err = generateTokens("push constant 3 //", LookingForContext)
	assertEqual(t, tokens[0], "push", "")
	assertEqual(t, tokens[1], "constant", "")
	assertEqual(t, tokens[2], "3", "")
	assertEqual(t, err, nil, "")

	tokens, _, err = generateTokens("add", LookingForContext)
	assertEqual(t, tokens[0], "add", "")

	tokens, _, err = generateTokens("push /* some stupid c//ommend */ constant 3 //", LookingForContext)
	assertEqual(t, tokens[0], "push", "")
	assertEqual(t, tokens[1], "constant", "")
	assertEqual(t, tokens[2], "3", "")
	assertEqual(t, err, nil, "")
}

func TestTokenize(t *testing.T) {
	instruction := tokenize([]string {"add"})
	assertEqual(t, instruction.command, CommandArithmetic, "")
	assertEqual(t, instruction.arg1, "add", "")
	assertEqual(t, instruction.arg2, -1, "")

	instruction = tokenize([]string { "push", "constant", "3", })
	assertEqual(t, instruction.command, CommandPush, "")
	assertEqual(t, instruction.arg1, "constant", "")
	assertEqual(t, instruction.arg2, 3, "")

	instruction = tokenize([]string {"pop", "constant", "3"})
	assertEqual(t, instruction.command, CommandPop, "")
	assertEqual(t, instruction.arg1, "constant", "")
	assertEqual(t, instruction.arg2, 3, "")
}
