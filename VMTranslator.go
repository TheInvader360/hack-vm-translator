package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TheInvader360/hack-vm-translator/generator"
	"github.com/TheInvader360/hack-vm-translator/handler"
	"github.com/TheInvader360/hack-vm-translator/parser"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing file parameter")
		return
	}
	inputFilename := os.Args[1]

	if !strings.HasSuffix(inputFilename, ".vm") {
		fmt.Println("Expected a vm file (*.vm)")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inputFilename)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't read file: %s", inputFilename)))
	fmt.Println("SOURCE:\n" + string(data) + "\n----------")

	parser := parser.NewParser()
	parser.Sanitize(data)
	fmt.Println("SANITIZED:\n" + strings.Join(parser.SourceLines, "\n") + "\n----------")
	commands, err := parser.ParseSource()
	handler.FatalError(err)
	fmt.Println("PARSED:")
	for _, command := range commands {
		fmt.Println(command)
	}
	fmt.Println("----------")

	generator := generator.NewGenerator()
	assemblyLines, err := generator.GenerateAssembly(commands)
	fmt.Println("ASSEMBLY:")
	for _, line := range assemblyLines {
		fmt.Println(line)
	}
	fmt.Println("----------")

	outputFilename := strings.Replace(inputFilename, ".vm", ".asm", 1)
	output := []byte("lines\nof\ncode\n")
	err = ioutil.WriteFile(outputFilename, output, 0777)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't write file: %s", outputFilename)))
}
