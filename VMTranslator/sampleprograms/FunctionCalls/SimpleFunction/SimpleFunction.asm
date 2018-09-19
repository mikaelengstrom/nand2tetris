// command: function, arg1: SimpleFunction.test, arg2: 2
(SimpleFunction.test)
@2
D=A
(JMP1-loop)
@JMP1-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP1-loop
0; JMP
(JMP1-initialized)
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
// command: push, arg1: local, arg2: 1
@1
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
// command: arithmetic, arg1: not, arg2: -1
@SP
A=M-1
M=!M
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
// command: arithmetic, arg1: sub, arg2: -1
@SP
AM=M-1
D=M
A=A-1
M=M-D
// command: return, arg1: , arg2: 0
@5
D=A
@LCL
A=M
A=A-D
D=M
@R14
M=D
@SP
A=M-1
D=M
@ARG
A=M
M=D
D=A+1
@SP
M=D
@LCL
D=M
@R13
AM=D-1
D=M
@THAT
M=D
@R13
AM=M-1
D=M
@THIS
M=D
@R13
AM=M-1
D=M
@ARG
M=D
@R13
AM=M-1
D=M
@LCL
M=D
@R13
A=M-1
@R14
A=M
0; JMP
