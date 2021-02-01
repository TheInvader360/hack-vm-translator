load SimpleEqTrue.asm,
output-file SimpleEqTrue.out,
compare-to SimpleEqTrue.cmp,
output-list RAM[0]%D2.6.2 RAM[256]%D2.6.2;

set RAM[0] 256,  // initializes the stack pointer 

repeat 50 {      // enough cycles to complete the execution
  ticktock;
}

output;          // the stack pointer and the stack base
