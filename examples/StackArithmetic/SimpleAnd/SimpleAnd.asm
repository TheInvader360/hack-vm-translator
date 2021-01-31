// push constant 14
@14
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 7
@7
D=A
@SP
A=M
M=D
@SP
M=M+1

// and
@SP
A=M-1
D=M
A=A-1
M=D&M
@SP
M=M-1

