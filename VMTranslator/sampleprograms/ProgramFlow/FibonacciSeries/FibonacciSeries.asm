// command: push, arg1: argument, arg2: 1
@1
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: pointer, arg2: 1
@SP
AM=M-1
D=M
@THAT
M=D
// command: push, arg1: constant, arg2: 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: that, arg2: 0
@0
D=A
@THAT
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: push, arg1: constant, arg2: 1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: that, arg2: 1
@1
D=A
@THAT
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: push, arg1: argument, arg2: 0
@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: arithmetic, arg1: sub, arg2: -1
@SP
AM=M-1
D=M
A=A-1
M=M-D
// command: pop, arg1: argument, arg2: 0
@0
D=A
@ARG
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: label, arg1: MAIN_LOOP_START, arg2: 0
(FibonacciSeries.MAIN_LOOP_START)
// command: push, arg1: argument, arg2: 0
@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: if-goto, arg1: COMPUTE_ELEMENT, arg2: 0
@SP
AM=M-1
D=M
@FibonacciSeries.COMPUTE_ELEMENT
D; JGT
// command: goto, arg1: END_PROGRAM, arg2: 0
@FibonacciSeries.END_PROGRAM
0; JMP
// command: label, arg1: COMPUTE_ELEMENT, arg2: 0
(FibonacciSeries.COMPUTE_ELEMENT)
// command: push, arg1: that, arg2: 0
@0
D=A
@THAT
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: that, arg2: 1
@1
D=A
@THAT
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: arithmetic, arg1: add, arg2: -1
@SP
A=M
A=A-1
D=M
A=A-1
M=D+M
D=A
@SP
M=D+1
// command: pop, arg1: that, arg2: 2
@2
D=A
@THAT
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: push, arg1: pointer, arg2: 1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: arithmetic, arg1: add, arg2: -1
@SP
A=M
A=A-1
D=M
A=A-1
M=D+M
D=A
@SP
M=D+1
// command: pop, arg1: pointer, arg2: 1
@SP
AM=M-1
D=M
@THAT
M=D
// command: push, arg1: argument, arg2: 0
@0
D=A
@ARG
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: arithmetic, arg1: sub, arg2: -1
@SP
AM=M-1
D=M
A=A-1
M=M-D
// command: pop, arg1: argument, arg2: 0
@0
D=A
@ARG
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: goto, arg1: MAIN_LOOP_START, arg2: 0
@FibonacciSeries.MAIN_LOOP_START
0; JMP
// command: label, arg1: END_PROGRAM, arg2: 0
(FibonacciSeries.END_PROGRAM)
