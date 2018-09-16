// command: push, arg1: constant, arg2: 111
@111
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 333
@333
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: constant, arg2: 888
@888
D=A
@SP
A=M
M=D
@SP
M=M+1
// command: pop, arg1: static, arg2: 8
@SP
AM=M-1
D=M
@StaticTest.vm.8
M=D
// command: pop, arg1: static, arg2: 3
@SP
AM=M-1
D=M
@StaticTest.vm.3
M=D
// command: pop, arg1: static, arg2: 1
@SP
AM=M-1
D=M
@StaticTest.vm.1
M=D
// command: push, arg1: static, arg2: 3
@StaticTest.vm.3
D=M
@SP
A=M
M=D
@SP
M=M+1
// command: push, arg1: static, arg2: 1
@StaticTest.vm.1
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
// command: push, arg1: static, arg2: 8
@StaticTest.vm.8
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
