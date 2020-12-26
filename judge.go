package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jessevdk/go-flags"
)

var debug bool = false

func main() {

	var opts struct {

		// Custom return code on success; Default 0
		RCS *int `short:"s" long:"rc_on_success" description:"custom return code on success; Default 0"`

		// Custom return code on failure; Default 1
		RCF *int `short:"f" long:"rc_on_failure" description:"custom return code on failure; Default 1"`

		// Slice of bool will append 'true' each time the option
		// is encountered (can be set multiple times, like -vvv)
		Verbose []bool `short:"v" long:"verbose" description:"show verbose debug information"`

		// Example of a pointer
		Lines *int `short:"l" long:"lines" description:"line count"`

		// Example of a value name
		//File string `short:"f" long:"file" description:"Set file to validate | Required" value-name:"PATH"`
	}

	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	_, err := flags.Parse(&opts)
	if err != nil {
		//panic(err)
		fmt.Println("wtf dood")
	}

	// Iterate over passed arguments to validate if mandatory ones exist
	// Exclude -v verbose and first 2 elements (go binary execution path and provided file PATH for testing)
	var arg_slice []string
	for i := 2; i < len(os.Args); i++ {
		if os.Args[i] != "-v" {
			arg_slice = append(arg_slice, os.Args[i])
		}
	}
	// Same argument list, but without numerical strings
	var arg_slice_no_numeric []string
	for i := 0; i < len(arg_slice); i++ {
		if isNumeric(arg_slice[i]) == false {
			arg_slice_no_numeric = append(arg_slice_no_numeric, arg_slice[i])
		}
	}

	// Check if custom return codes were defined
	//for i := 1; i < len(arg_slice); i++ {
	//	if i == "-s"
	//}

	// Check if debug output is enabled
	if len(opts.Verbose) > 0 {
		debug = true
		fmt.Println("Verbose output enabled.\n")
	}

	// This block checks for errors
	// Check if there are arguments passed into application
	if len(os.Args[1:]) > 0 {
		filepath := os.Args[1]

		// Check if file or directory path provided is valid
		if _, err := os.Stat(filepath); err == nil {
			if debug == true {
				msg_out(1, "file or directory path is valid")
			}

		} else if _, err := os.Stat(filepath); err != nil {
			msg_err(1, "provided path is not valid")
		}

		if len(arg_slice_no_numeric) < 1 {
			msg_err(1, "no arguments provided")
		} else if len(opts.Verbose) > 0 {
			msg_out(1, "valid arguments provided:")
			fmt.Println(arg_slice_no_numeric)
		}

		//fmt.Println(arg_slice_no_numeric)

	} else {
		msg_err(1, "no valid file or directory path provided")
	}
	// ---- end of argument validation block

}

func msg_out(verbosity int, message string) {
	if verbosity > 0 {
		fmt.Println(":)", message)
	}
}

func msg_err(verbosity int, message string) {
	if verbosity > 0 {
		fmt.Println(":(", message)
	}
	fmt.Println("\nUsage: ./judge PATH <arg1> <argX> ..\nOr use --help to show more info")
	os.Exit(1)
}

// Check if a string is numeric (integers & floats)
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
