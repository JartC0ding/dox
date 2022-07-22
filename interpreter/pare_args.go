package interpreter

import (
	fileopener "dox/file_opener"
	mode_manager "dox/modes"
	"fmt"
	"os"
)

var modes []string = []string{
	"c", // Create
	"d", // Delete
	"e", // Edit
}

// Parses all given commandline arguments (argv)
func ParseArgv(argv []string, argc int) {
	/*
		Parsing:
			- Parse mode (new, del, edit, or nothing)
			- Parse File (filename, or nothing)

			if both are nothing. Open the document describing the whole project
	*/

	if argc > 3 {
		// error. Too many arguments
		fmt.Println("Error. Too many arguments")
		os.Exit(0)
	}

	// Parse mode and file
	var mode string = ""
	var file string = ""

	if argc > 1 {
		mode = argv[1]

		// validate mode
		var is_valid bool = false
		for _, val := range modes {
			if mode == val {
				is_valid = true
			}
		}

		if argc > 2 && is_valid {
			file = argv[2]
		} else if !is_valid && argc > 2 {
			// error. Unexpected argument
			fmt.Println("Error. Unexpected argument")
			os.Exit(0)
		} else if is_valid && argc <= 2 {
			// error. Argument expected
			fmt.Println("Error. Argument expected")
			os.Exit(0)
		} else {
			file = mode
			mode = ""
		}
	} else {
		// error. Too few arguments
		fmt.Println("Error. Too few arguments")
		os.Exit(0)
	}

	// Call function based on inputs
	if mode != "" && file != "" {
		// normal mode
		mode_manager.SwitchMode(mode, file)
	} else if mode == "" && file != "" {
		// open file
		fileopener.OpenFile(file)
	}

}
