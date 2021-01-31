// push constant 12
@12
D=A
@SP
A=M
M=D
@SP
M=M+1

// push constant 9
@9
D=A
@SP
A=M
M=D
@SP
M=M+1

// or
@SP
A=M-1
D=M
A=A-1
M=D|M
@SP
M=M-1

