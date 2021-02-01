Detailed deconstruction of SimpleAdd

SimpleAdd.vm:

```
push constant 7 // x=7
push constant 8 // y=8
add             // y+x
```

Init test setup (CPUEmulator):
* RAM[0] = 256
* Step through 22 execution cycles

Expected end result:
* RAM[0] = 257 (i.e. SP++, SP++, SP--)
* RAM[256] = 15 (i.e. y+x = 8+7)

ASM:

```
      // *SP=7, SP++ (i.e. RAM[256]=7, RAM[0]=257):
@7    // PC= 1, A=7
D=A   // PC= 2, D=7
@SP   // PC= 3, A=0
A=M   // PC= 4, A=RAM[0] (i.e A=256)
M=D   // PC= 5, RAM[256]=7                                : RAM[256]=7 (i.e. x=7)
@SP   // PC= 6, A=0
M=M+1 // PC= 7, RAM[0]=RAM[0]+1 (i.e. RAM[0]=257)         : SP++

      // *SP=8, SP++ (i.e. RAM[257]=8, RAM[0]=258):
@8    // PC= 8, A=8
D=A   // PC= 9, D=8
@SP   // PC=10, A=0
A=M   // PC=11, A=RAM[0] (i.e. A=257)
M=D   // PC=12, RAM[257]=8                                : RAM[257]=8 (i.e. y=8)
@SP   // PC=13, A=0
M=M+1 // PC=14, RAM[0]=RAM[0]+1 (i.e. RAM[0]=258)         : SP++

      // RAM[SP-2]=RAM[SP-1]+RAM[SP-2], SP--:
@SP   // pc=15, A=0
A=M-1 // pc=16, A=RAM[0]-1 (i.e. A=257)
D=M   // pc=17, D=RAM[257] (i.e. D=8)                     : D=8 (or D=y)
A=A-1 // pc=18, A=256
M=D+M // pc=19, RAM[256]=D+RAM[256] (i.e RAM[256]=8+7=15) : RAM[256]=8+7 (or RAM[256]=y+x)
@SP   // pc=20, A=0
M=M-1 // pc=21, RAM[0]=RAM[0]-1 (i.e. RAM[0]=257)         : SP--
```

Arrived at the desired RAM[0]=257 and RAM[256]=15 values.
