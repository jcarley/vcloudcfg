package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

var (
	Info *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
}

func main() {
	os.Exit(realMain())
}

func realMain() int {

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	args := os.Args[1:]
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], args)
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("vcloudcfg"),
	}

	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitCode
}
