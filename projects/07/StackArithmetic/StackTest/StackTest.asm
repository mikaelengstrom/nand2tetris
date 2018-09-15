// command: push, arg1: constant, arg2: 17
@17
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 17
@17
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: eq, arg2: -1
@SP
A=M
A=A-1
D=M
A=A-1
D=D-M;
@JMP3-eq
D; JEQ
D=1
@JMP3-end
(JMP3-eq)
D=D-1
(JMP3-end)
@SP
A=M-1
A=A-1
M=D
D=A+1
@SP
M=D
// command: push, arg1: constant, arg2: 17
@17
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 16
@16
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: eq, arg2: -1
@SP
A=M
A=A-1
D=M
A=A-1
D=D-M;
@JMP6-eq
D; JEQ
D=1
@JMP6-end
(JMP6-eq)
D=D-1
(JMP6-end)
@SP
A=M-1
A=A-1
M=D
D=A+1
@SP
M=D
// command: push, arg1: constant, arg2: 16
@16
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 17
@17
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: eq, arg2: -1
@SP
A=M
A=A-1
D=M
A=A-1
D=D-M;
@JMP9-eq
D; JEQ
D=1
@JMP9-end
(JMP9-eq)
D=D-1
(JMP9-end)
@SP
A=M-1
A=A-1
M=D
D=A+1
@SP
M=D
// command: push, arg1: constant, arg2: 892
@892
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 891
@891
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: lt, arg2: -1
@SP
A=M
MD=D-M;
@JMP12-lt
D; JEQ
D=1
(JMP12-lt)
D=D-1
@SP
AM=M-1
A=A-1
M=D
// command: push, arg1: constant, arg2: 891
@891
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 892
@892
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: lt, arg2: -1
@SP
A=M
MD=D-M;
@JMP15-lt
D; JEQ
D=1
(JMP15-lt)
D=D-1
@SP
AM=M-1
A=A-1
M=D
// command: push, arg1: constant, arg2: 891
@891
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 891
@891
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: arithmetic, arg1: lt, arg2: -1
@SP
A=M
MD=D-M;
@JMP18-lt
D; JEQ
D=1
(JMP18-lt)
D=D-1
@SP
AM=M-1
A=A-1
M=D
// command: push, arg1: constant, arg2: 32767
@32767
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 32766
@32766
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 32766
@32766
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 32767
@32767
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 32766
@32766
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 32766
@32766
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 57
@57
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 31
@31
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 53
@53
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
// command: push, arg1: constant, arg2: 112
@112
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
// command: push, arg1: constant, arg2: 82
@82
D=A
@SP
A=M
M=D
D=A
@SP
M=D+1
