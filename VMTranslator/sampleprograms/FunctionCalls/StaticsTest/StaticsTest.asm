@256
D=A
@SP
M=D
// command: call, arg1: Sys.init, arg2: 0
@SP
D=M
@0
D=D-A
@R13
M=D
@JMP0-return-address
D=A
@SP
A=M
M=D
@LCL
D=M
@SP
AM=M+1
M=D
@ARG
D=M
@SP
AM=M+1
M=D
@THIS
D=M
@SP
AM=M+1
M=D
@THAT
D=M
@SP
AM=M+1
M=D
D=A+1
@LCL
M=D
@SP
M=M+1
@R13
D=M
@ARG
M=D
@Sys.init
0; JMP
(JMP0-return-address)
// command: function, arg1: Class1.set, arg2: 0
(Class1.set)
@0
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
// command: pop, arg1: static, arg2: 0
@SP
AM=M-1
D=M
@Class1.0
M=D
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
// command: pop, arg1: static, arg2: 1
@SP
AM=M-1
D=M
@Class1.1
M=D
// command: push, arg1: constant, arg2: 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
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
// command: function, arg1: Class1.get, arg2: 0
(Class1.get)
@0
D=A
(JMP8-loop)
@JMP8-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP8-loop
0; JMP
(JMP8-initialized)
// command: push, arg1: static, arg2: 0
@Class1.0
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: static, arg2: 1
@Class1.1
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
// command: function, arg1: Class2.set, arg2: 0
(Class2.set)
@0
D=A
(JMP13-loop)
@JMP13-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP13-loop
0; JMP
(JMP13-initialized)
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
// command: pop, arg1: static, arg2: 0
@SP
AM=M-1
D=M
@Class2.0
M=D
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
// command: pop, arg1: static, arg2: 1
@SP
AM=M-1
D=M
@Class2.1
M=D
// command: push, arg1: constant, arg2: 0
@0
D=A
@SP
A=M
M=D
@SP
M=M+1
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
// command: function, arg1: Class2.get, arg2: 0
(Class2.get)
@0
D=A
(JMP20-loop)
@JMP20-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP20-loop
0; JMP
(JMP20-initialized)
// command: push, arg1: static, arg2: 0
@Class2.0
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: static, arg2: 1
@Class2.1
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
// command: function, arg1: Sys.init, arg2: 0
(Sys.init)
@0
D=A
(JMP25-loop)
@JMP25-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP25-loop
0; JMP
(JMP25-initialized)
// command: push, arg1: constant, arg2: 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 8
@8
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: call, arg1: Class1.set, arg2: 2
@SP
D=M
@2
D=D-A
@R13
M=D
@JMP28-return-address
D=A
@SP
A=M
M=D
@LCL
D=M
@SP
AM=M+1
M=D
@ARG
D=M
@SP
AM=M+1
M=D
@THIS
D=M
@SP
AM=M+1
M=D
@THAT
D=M
@SP
AM=M+1
M=D
D=A+1
@LCL
M=D
@SP
M=M+1
@R13
D=M
@ARG
M=D
@Class1.set
0; JMP
(JMP28-return-address)
// command: pop, arg1: temp, arg2: 0
@SP
AM=M-1
D=M
@R5
M=D
// command: push, arg1: constant, arg2: 23
@23
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 15
@15
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: call, arg1: Class2.set, arg2: 2
@SP
D=M
@2
D=D-A
@R13
M=D
@JMP32-return-address
D=A
@SP
A=M
M=D
@LCL
D=M
@SP
AM=M+1
M=D
@ARG
D=M
@SP
AM=M+1
M=D
@THIS
D=M
@SP
AM=M+1
M=D
@THAT
D=M
@SP
AM=M+1
M=D
D=A+1
@LCL
M=D
@SP
M=M+1
@R13
D=M
@ARG
M=D
@Class2.set
0; JMP
(JMP32-return-address)
// command: pop, arg1: temp, arg2: 0
@SP
AM=M-1
D=M
@R5
M=D
// command: call, arg1: Class1.get, arg2: 0
@SP
D=M
@0
D=D-A
@R13
M=D
@JMP34-return-address
D=A
@SP
A=M
M=D
@LCL
D=M
@SP
AM=M+1
M=D
@ARG
D=M
@SP
AM=M+1
M=D
@THIS
D=M
@SP
AM=M+1
M=D
@THAT
D=M
@SP
AM=M+1
M=D
D=A+1
@LCL
M=D
@SP
M=M+1
@R13
D=M
@ARG
M=D
@Class1.get
0; JMP
(JMP34-return-address)
// command: call, arg1: Class2.get, arg2: 0
@SP
D=M
@0
D=D-A
@R13
M=D
@JMP35-return-address
D=A
@SP
A=M
M=D
@LCL
D=M
@SP
AM=M+1
M=D
@ARG
D=M
@SP
AM=M+1
M=D
@THIS
D=M
@SP
AM=M+1
M=D
@THAT
D=M
@SP
AM=M+1
M=D
D=A+1
@LCL
M=D
@SP
M=M+1
@R13
D=M
@ARG
M=D
@Class2.get
0; JMP
(JMP35-return-address)
// command: label, arg1: WHILE, arg2: 0
(Sys.WHILE)
// command: goto, arg1: WHILE, arg2: 0
@Sys.WHILE
0; JMP