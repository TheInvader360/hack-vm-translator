# Hack VM Translator

A VM Translator that translates VM language programs into Hack assembly code (nand2tetris)

### Usage

```bash
git clone https://github.com/TheInvader360/hack-vm-translator
cd hack-vm-translator
```

Then:

```bash
go run VMTranslator.go examples/StackArithmetic/SimpleAdd/SimpleAdd.vm
```

The translated program is exported to the source directory with the same base filename but a .asm extension (so the given example would generate examples/StackArithmetic/SimpleAdd/SimpleAdd.asm)
