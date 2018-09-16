// command: push, arg1: constant, arg2: 3030
@3030
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: pointer, arg2: 0
@SP
AM=M-1
D=M
@THIS
M=D
// command: push, arg1: constant, arg2: 3040
@3040
D=A
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
// command: push, arg1: constant, arg2: 32
@32
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: this, arg2: 2
@2
D=A
@THIS
D=D+M
@R14
M=D
@SP
AM=M-1
D=M
@R14
A=M
M=D
// command: push, arg1: constant, arg2: 46
@46
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: that, arg2: 6
@6
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
// command: push, arg1: pointer, arg2: 0
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: pointer, arg2: 1
@THAT
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
// command: push, arg1: this, arg2: 2
@2
D=A
@THIS
A=D+M
D=M
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
// command: push, arg1: that, arg2: 6
@6
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
