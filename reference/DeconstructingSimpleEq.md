Detailed deconstruction of SimpleEq

SimpleEq.vm:

```
push constant 17 // x=17
push constant 17 // y=17
eq               // x=y?
```

Process (CPUEmulator):
* Step through execution cycles

Expected end result:
* RAM[0] = 257 (i.e. SP++, SP++, SP--)
* RAM[256] = -1 (i.e. x=y evaluated to true)

ASM:

```
       // initialise (RAM[0] i.e. SP=256)
@256   // A=256
D=A    // D=256
@SP    // A=0
M=D    // RAM[0]=256

       // push constant 17: *SP=17, SP++ (i.e. RAM[256]=17, RAM[0]=257)
@17    // A=17
D=A    // D=17
@SP    // A=0
A=M    // A=RAM[0] (i.e A=256)
M=D    // RAM[256]=17
@SP    // A=0
M=M+1  // RAM[0]=RAM[0]+1 (i.e. RAM[0]=257, or SP++)

       // push constant 17: *SP=17, SP++ (i.e. RAM[257]=17, RAM[0]=258)
@17    // A=17
D=A    // D=17
@SP    // A=0
A=M    // A=RAM[0] (i.e A=257)
M=D    // RAM[257]=17
@SP    // A=0
M=M+1  // RAM[0]=RAM[0]+1 (i.e. RAM[0]=258, or SP++)

       // if (RAM[SP-1]=RAM[SP-2]) then RAM[SP-2]=-1 else RAM[SP-2]=0, SP--
@SP    // A=0
A=M-1  // A=RAM[0]-1 (i.e. A=257)
D=M    // D=RAM[257] (i.e. D=17)
A=A-1  // A=256
D=M-D  // D=RAM[256]-D=17-17=0
@TRUE0 // set A to address of (TRUE0) label
D;JEQ  // if D=0 jump to (TRUE0) i.e. skips code that sets result to 0
@SP    // A=0
A=M-1  // A=RAM[256]-1 (i.e. A=257)
A=A-1  // A=256
M=0    // RAM[256]=0
@CONT0 // set A to address of (CONT0) label
0;JMP  // unconditional jump to (CONT0) i.e. skips code that sets result to -1
(TRUE0)// label
@SP    // A=0
A=M-1  // A=RAM[256]-1 (i.e. A=257)
A=A-1  // A=256
M=-1   // RAM[256]=-1
(CONT0)// label
@SP    // A=0
M=M-1  // RAM[0]=RAM[0]-1 (i.e. RAM[0]=257, or SP--)
```

Arrived at the desired RAM[0]=257 and RAM[256]=-1 values.

Only change needed to test for "less than" or "greater than" is to replace D;JEQ with either D;JGT or D;JLT.
