// command: function, arg1: Sys.init, arg2: 0
(Sys.Sys.init)
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
// command: push, arg1: constant, arg2: 4000
@4000
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
// command: push, arg1: constant, arg2: 5000
@5000
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
// command: call, arg1: Sys.main, arg2: 0
@SP
D=M
@0
D=D-A
@R13
M=D
@JMP6-return-address
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
@Sys.Sys.main
0; JMP
(JMP6-return-address)
// command: pop, arg1: temp, arg2: 1
@SP
AM=M-1
D=M
@R6
M=D
// command: label, arg1: LOOP, arg2: 0
(Sys.LOOP)
// command: goto, arg1: LOOP, arg2: 0
@Sys.LOOP
0; JMP
// command: function, arg1: Sys.main, arg2: 5
(Sys.Sys.main)
@5
D=A
(JMP10-loop)
@JMP10-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP10-loop
0; JMP
(JMP10-initialized)
// command: push, arg1: constant, arg2: 4001
@4001
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
// command: push, arg1: constant, arg2: 5001
@5001
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
// command: push, arg1: constant, arg2: 200
@200
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: local, arg2: 1
@1
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
// command: push, arg1: constant, arg2: 40
@40
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: local, arg2: 2
@2
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
// command: push, arg1: constant, arg2: 6
@6
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: local, arg2: 3
@3
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
// command: push, arg1: constant, arg2: 123
@123
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: call, arg1: Sys.add12, arg2: 1
@SP
D=M
@1
D=D-A
@R13
M=D
@JMP22-return-address
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
@Sys.Sys.add12
0; JMP
(JMP22-return-address)
// command: pop, arg1: temp, arg2: 0
@SP
AM=M-1
D=M
@R5
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
// command: push, arg1: local, arg2: 2
@2
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: local, arg2: 3
@3
D=A
@LCL
A=D+M
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: local, arg2: 4
@4
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
// command: function, arg1: Sys.add12, arg2: 0
(Sys.Sys.add12)
@0
D=A
(JMP34-loop)
@JMP34-initialized
D=D-1; JLT
@SP
A=M
M=0
@SP
M=M+1
@JMP34-loop
0; JMP
(JMP34-initialized)
// command: push, arg1: constant, arg2: 4002
@4002
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
// command: push, arg1: constant, arg2: 5002
@5002
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
// command: push, arg1: constant, arg2: 12
@12
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