// command: push, arg1: constant, arg2: 10
@10
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
// command: push, arg1: constant, arg2: 21
@21
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 22
@22
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: argument, arg2: 2
@2
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
// command: pop, arg1: argument, arg2: 1
@1
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
// command: push, arg1: constant, arg2: 36
@36
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: this, arg2: 6
@6
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
// command: push, arg1: constant, arg2: 42
@42
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 45
@45
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: that, arg2: 5
@5
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
// command: push, arg1: constant, arg2: 510
@510
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: temp, arg2: 6
@SP
AM=M-1
D=M
@R11
M=D
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
// command: push, arg1: that, arg2: 5
@5
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
// command: push, arg1: this, arg2: 6
@6
D=A
@THIS
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: this, arg2: 6
@6
D=A
@THIS
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
// command: arithmetic, arg1: sub, arg2: -1
@SP
AM=M-1
D=M
A=A-1
M=M-D
// command: push, arg1: temp, arg2: 6
@Temp6
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
