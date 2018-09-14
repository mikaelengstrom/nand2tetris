// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Set variables
   @SCREEN
   D=A
   @address
   M=D

   @8192 // 1024 * 8, whole range of the screen
   D=A
   @length
   M=D

(LOOP)
   @KBD
   D=M
   @lastinput
   M=D

(LISTEN)
   @lastinput
   D=M
   @KBD
   D=D-M
   @LISTEN
   D; JEQ // input did change
// On change, jump to RELEASE on keyup
   @KBD
   D=M
   @DRAWWHITE
   D; JEQ
   @DRAWBLACK
   0; JMP

(DRAWWHITE)
   @color
   M=0
   @DRAW
   0; JMP

(DRAWBLACK)
   @color
   M=-1
   @DRAW
   0; JMP

(DRAW)
   @i
   M=0
   (WRITE)
   // Get address
      @SCREEN
      D=A
      @i
      D=D+M
      @address
      M=D

   // if screen is full: break;
      @length
      D=M
      @i
      D=D-M
      @ENDWRITE
      D; JEQ

   // Set color
      @color
      D=M

   // Write
      @address
      A=M
      M=D

   // Increment iterator
      @i
      M=M+1

      @WRITE
      0; JMP

   (ENDWRITE)
     @LOOP
     0; JMP
