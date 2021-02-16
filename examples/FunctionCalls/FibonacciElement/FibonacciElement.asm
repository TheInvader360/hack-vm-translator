
// Bootstrap...
@256
D=A
@SP
M=D
@SP
D=M
@R13
M=D
@ret0
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@R13
D=M
@0
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Sys.init
0;JMP
(ret0)

// function Main.fibonacci 0
(Main.fibonacci)
@SP
A=M
D=A
@SP
M=D

// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1

// lt
@SP
A=M-1
D=M
A=A-1
D=M-D
@TRUE1
D;JLT
@SP
A=M-1
A=A-1
M=0
@CONT1
0;JMP
(TRUE1)
@SP
A=M-1
A=A-1
M=-1
(CONT1)
@SP
M=M-1

// if-goto IF_TRUE
@SP
AM=M-1
D=M
@Main.fibonacci$IF_TRUE
D;JNE

// goto IF_FALSE
@Main.fibonacci$IF_FALSE
0;JMP

// label IF_TRUE
(Main.fibonacci$IF_TRUE)

// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// return
@LCL
D=M
@5
A=D-A
D=M
@R13
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
AM=M-1
D=M
@THAT
M=D
@LCL
AM=M-1
D=M
@THIS
M=D
@LCL
AM=M-1
D=M
@ARG
M=D
@LCL
A=M-1
D=M
@LCL
M=D
@R13
A=M
0;JMP

// label IF_FALSE
(Main.fibonacci$IF_FALSE)

// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// push constant 2
@2
D=A
@SP
A=M
M=D
@SP
M=M+1

// sub
@SP
A=M-1
D=M
A=A-1
M=M-D
@SP
M=M-1

// call Main.fibonacci 1
@SP
D=M
@R13
M=D
@ret2
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@R13
D=M
@1
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Main.fibonacci
0;JMP
(ret2)

// push argument 0
@ARG
D=M
@0
A=D+A
D=M
@SP
A=M
M=D
@SP
M=M+1

// push constant 1
@1
D=A
@SP
A=M
M=D
@SP
M=M+1

// sub
@SP
A=M-1
D=M
A=A-1
M=M-D
@SP
M=M-1

// call Main.fibonacci 1
@SP
D=M
@R13
M=D
@ret3
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@R13
D=M
@1
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Main.fibonacci
0;JMP
(ret3)

// add
@SP
A=M-1
D=M
A=A-1
M=D+M
@SP
M=M-1

// return
@LCL
D=M
@5
A=D-A
D=M
@R13
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
AM=M-1
D=M
@THAT
M=D
@LCL
AM=M-1
D=M
@THIS
M=D
@LCL
AM=M-1
D=M
@ARG
M=D
@LCL
A=M-1
D=M
@LCL
M=D
@R13
A=M
0;JMP

// function Sys.init 0
(Sys.init)
@SP
A=M
D=A
@SP
M=D

// push constant 4
@4
D=A
@SP
A=M
M=D
@SP
M=M+1

// call Main.fibonacci 1
@SP
D=M
@R13
M=D
@ret4
D=A
@SP
A=M
M=D
@SP
M=M+1
@LCL
D=M
@SP
A=M
M=D
@SP
M=M+1
@ARG
D=M
@SP
A=M
M=D
@SP
M=M+1
@THIS
D=M
@SP
A=M
M=D
@SP
M=M+1
@THAT
D=M
@SP
A=M
M=D
@SP
M=M+1
@R13
D=M
@1
D=D-A
@ARG
M=D
@SP
D=M
@LCL
M=D
@Main.fibonacci
0;JMP
(ret4)

// label WHILE
(Sys.init$WHILE)

// goto WHILE
@Sys.init$WHILE
0;JMP

