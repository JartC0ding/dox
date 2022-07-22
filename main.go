package main

import (
	"dox/interpreter"
	"os"
)

func main() {
	// parse argv and suspend execution
	interpreter.ParseArgv(os.Args, len(os.Args))
}
