// command: push, arg1: constant, arg2: 7
@7
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 8
@8
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
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
