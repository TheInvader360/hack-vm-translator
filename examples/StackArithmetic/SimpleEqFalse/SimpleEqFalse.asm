// push constant 123
@123
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 321
@321
D=A
@SP
A=M
M=D
@SP
M=M+1

// eq
@SP
A=M-1
D=M
A=A-1
D=M-D
@TRUE0
D;JEQ
@SP
A=M-1
A=A-1
M=0
@CONT0
0;JMP
(TRUE0)
@SP
A=M-1
A=A-1
M=-1
(CONT0)
@SP
M=M-1

