package handler

import (
	"fmt"
	"os"
)

// FatalError - handles an error by printing a message then exiting
func FatalError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
