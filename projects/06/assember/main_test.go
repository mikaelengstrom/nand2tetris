package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestStripComments(t *testing.T) {
	if stripComments("   // hej") != "   " {
		t.Error("Did not remove all whitespace or comments")
	}

	res := stripComments("/* aasdfasdfasdf */ A = E; whatevs  ")
	if res != " A = E; whatevs  " {
		t.Error("Did not remove /* */ comments")
	}

	if res == "" {
		t.Error("stripCommentsAndWhitespace removed everything :)")
	}
}

func TestParseAInstruction(t *testing.T) {
	st := NewSymbolTable()
	assertEqual(t, parseAInstruction("@1", st), "0000000000000001", "")
	assertEqual(t, parseAInstruction("@9", st), "0000000000001001", "")
	assertEqual(t, parseAInstruction("@32767", st), "0111111111111111", "")
}

func TestParseCInstruction(t *testing.T) {
	// Destinations
	assertEqual(t, parseCInstruction("null=D-A"),"1110010011000000", "")
	assertEqual(t, parseCInstruction(" M= D-A"), "1110010011001000", "")
	assertEqual(t, parseCInstruction("D = D-A"), "1110010011010000", "")
	assertEqual(t, parseCInstruction("MD= D-A"), "1110010011011000", "")
	assertEqual(t, parseCInstruction("A = D-A"), "1110010011100000", "")
	assertEqual(t, parseCInstruction("AM= D-A"), "1110010011101000", "")
	assertEqual(t, parseCInstruction("AD= D-A"), "1110010011110000", "")
	assertEqual(t, parseCInstruction("AMD=D-A"), "1110010011111000", "")


	// Jump bits
	assertEqual(t, parseCInstruction("1; null"), "1110111111000000", "")
	assertEqual(t, parseCInstruction("1; JGT"), "1110111111000001", "")
	assertEqual(t, parseCInstruction("1; JEQ"), "1110111111000010", "")
	assertEqual(t, parseCInstruction("1; JGE"), "1110111111000011", "")
	assertEqual(t, parseCInstruction("1; JLT"), "1110111111000100", "")
	assertEqual(t, parseCInstruction("1; JNE"), "1110111111000101", "")
	assertEqual(t, parseCInstruction("1; JLE"), "1110111111000110", "")
	assertEqual(t, parseCInstruction("1; JMP"), "1110111111000111", "")

	// Control bits (a=0)
	assertEqual(t, parseCInstruction("M=0"), "1110101010001000", "")
	assertEqual(t, parseCInstruction("M=1"), "1110111111001000", "")
	assertEqual(t, parseCInstruction("M=-1"), "1110111010001000", "")
	assertEqual(t, parseCInstruction("M=D"), "1110001100001000", "")
	assertEqual(t, parseCInstruction("M=A"), "1110110000001000", "")
	assertEqual(t, parseCInstruction("M=!D"), "1110001101001000", "")
	assertEqual(t, parseCInstruction("M=!A"), "1110110001001000", "")
	assertEqual(t, parseCInstruction("M=-D"), "1110001111001000", "")
	assertEqual(t, parseCInstruction("M=-A"), "1110110011001000", "")
	assertEqual(t, parseCInstruction("M=D+1"), "1110011111001000", "")
	assertEqual(t, parseCInstruction("M=A+1"), "1110110111001000", "")
	assertEqual(t, parseCInstruction("M=D-1"), "1110001110001000", "")
	assertEqual(t, parseCInstruction("M=A-1"), "1110110010001000", "")
	assertEqual(t, parseCInstruction("M=D+A"), "1110000010001000", "")
	assertEqual(t, parseCInstruction("M=D-A"), "1110010011001000", "")
	assertEqual(t, parseCInstruction("M=A-D"), "1110000111001000", "")
	assertEqual(t, parseCInstruction("M=D&A"), "1110000000001000", "")
	assertEqual(t, parseCInstruction("M=D|A"), "1110010101001000", "")

	// Control bits (a=1)
	assertEqual(t, parseCInstruction("M=M"), "1111110000001000", "")
	assertEqual(t, parseCInstruction("M=!M"), "1111110001001000", "")
	assertEqual(t, parseCInstruction("M=-M"), "1111110011001000", "")
	assertEqual(t, parseCInstruction("M=M+1"), "1111110111001000", "")
	assertEqual(t, parseCInstruction("M=M-1"), "1111110010001000", "")
	assertEqual(t, parseCInstruction("M=D+M"), "1111000010001000", "")
	assertEqual(t, parseCInstruction("M=D-M"), "1111010011001000", "")
	assertEqual(t, parseCInstruction("M=M-D"), "1111000111001000", "")
	assertEqual(t, parseCInstruction("M=D&M"), "1111000000001000", "")
	assertEqual(t, parseCInstruction("M=D|M"), "1111010101001000", "")
}

func TestSymbolTableGetOrCreate(t *testing.T) {
	st := NewSymbolTable()

	assertEqual(t, st.GetOrCreate("R0"), 0, "")
	assertEqual(t, st.GetOrCreate("hej"), 16, "")
	assertEqual(t, st.GetOrCreate("hej2"), 17, "")
}

func TestSymbolTableAddLabel(t *testing.T) {
	st := NewSymbolTable()

	// Simulate appearnce before label declaration
	st.GetOrCreate("WHATEVERLABEL")
	st.AddLabel("WHATEVERLABEL", 1)

	assertEqual(t, st.GetOrCreate("WHATEVERLABEL"), 1, "")
}

func TestSymbolTableMaybeReplaceSymbol(t *testing.T) {
	st := NewSymbolTable()

	// Simulate appearnce before label declaration
	st.GetOrCreate("somevariable")
	st.GetOrCreate("WHATEVERLABEL")
	st.AddLabel("WHATEVERLABEL", 1)
	st.MaybeReplaceSymbol("@somevariable")

	v, _ := st.MaybeReplaceSymbol("@somevariable")
	assertEqual(t, v, "@16", "")

	v, _ = st.MaybeReplaceSymbol("@WHATEVERLABEL")
	assertEqual(t, v, "@1", "")

	v, _ = st.MaybeReplaceSymbol("@R0")
	assertEqual(t, v, "@0", "")
}
