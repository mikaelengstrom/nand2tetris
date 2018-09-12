// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// for n R0:
//   R2 = R2 + R1
// exit

// Clean up from previos run
@R2
M=0

// Setup variables
@i
M=0

// Assert not negative, will be overriden if thats the case
@isneg
M=0

// Calculate number of iterations and set to N
@R0
D=M

@POS
D; JGE

@isneg
M=-1
D=!D
D=D+1
(POS)
  @n
  M=D

(LOOP)
  // if n - i is 0: break;
  @n
  D=M
  @i
  D=D-M

  @BREAK
  D;JEQ

  @R1
  D=M
  @R2
  M=D+M

  // i++
  @i
  M=M+1

  @LOOP
  0;JMP

(BREAK)
  // if R0 is possitive, goto END
  @isneg
  D=M
  @END
  D; JEQ

  // Invert R2
  @R2
  D=!M
  M=D+1


(END)
  @END
  0;JMP

