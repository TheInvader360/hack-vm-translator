package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/TheInvader360/hack-vm-translator/generator"
	"github.com/TheInvader360/hack-vm-translator/handler"
	"github.com/TheInvader360/hack-vm-translator/parser"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) != 2 {
		handler.FatalError(errors.New("Missing path parameter"))
	}
	inputFilepath := os.Args[1]
	outputFilepath := ""

	info, err := os.Stat(inputFilepath)
	handler.FatalError(err)

	if info.IsDir() {
		outputFilename := fmt.Sprintf("%s.asm", filepath.Base(inputFilepath))
		if strings.HasSuffix(inputFilepath, "/") {
			inputFilepath = inputFilepath + "*"
		} else {
			inputFilepath = inputFilepath + "/*"
		}
		outputFilepath = strings.Replace(inputFilepath, "*", outputFilename, 1)
	} else {
		if !strings.HasSuffix(inputFilepath, ".vm") {
			handler.FatalError(errors.New("Expected a vm file (*.vm)"))
		}
		outputFilepath = strings.Replace(inputFilepath, ".vm", ".asm", 1)
	}

	files, err := filepath.Glob(inputFilepath)
	handler.FatalError(err)

	generator := generator.NewGenerator()
	asm := ""
	for _, file := range files {
		if strings.HasSuffix(file, "Sys.vm") {
			asm = generator.Bootstrap()
		}
	}

	fmt.Println("------------------------------")
	for _, file := range files {
		if strings.HasSuffix(file, ".vm") {
			data, err := ioutil.ReadFile(file)
			handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't read file: %s", file)))

			parser := parser.NewParser()
			parser.Sanitize(data)
			commands, err := parser.ParseSource()
			handler.FatalError(err)

			asm = asm + generator.GenerateAssembly(strings.Replace(filepath.Base(file), ".vm", "", 1), commands)
		}
	}
	fmt.Print(asm)

	output := []byte(asm)
	if len(output) == 0 {
		handler.FatalError(errors.New("Zero length output - path must contain valid vm files (*.vm)"))
	}

	err = ioutil.WriteFile(outputFilepath, output, 0777)
	handler.FatalError(errors.Wrap(err, fmt.Sprintf("Can't write file: %s", outputFilepath)))
}
