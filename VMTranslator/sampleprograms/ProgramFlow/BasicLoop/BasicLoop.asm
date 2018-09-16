// command: push, arg1: constant, arg2: 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: local, arg2: 0
@0
D=A
@LCL
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
// command: push, arg1: local, arg2: 0
@0
D=A
@LCL
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
// command: pop, arg1: local, arg2: 0
@0
D=A
@LCL
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
// command: push, arg1: local, arg2: 0
@0
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
