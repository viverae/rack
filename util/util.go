package util

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
)

// Name is the name of the CLI
var Name = "rackcli"

// Version is the current CLI version
var Version = "0.1"

// Contains checks whether a given string is in a provided slice of strings.
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CommonFlags are flags that all commands can use. There exists the possiblity
// of setting app-level (global) flags, but that requires a user to properly
// position them. Including these with the other command-level flags will allow
// users to include them anywhere after the last subcommand (or argument, if applicable).
func commonFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "json",
			Usage: "Return output in JSON format.",
		},
		cli.BoolFlag{
			Name:  "table",
			Usage: "Return output in tabular format. This is the default output format.",
		},
	}
}

// CommandFlags returns the flags for a given command. It takes as a parameter
// a function for returning flags specific to that command, and then appends those
// flags with flags that are valid for all commands.
func CommandFlags(f func() []cli.Flag) []cli.Flag {
	cf := commonFlags()
	return append(cf, f()...)
}

// CheckArgNum checks that the provided number of arguments has the same
// cardinality as the expected number of arguments.
func CheckArgNum(c *cli.Context, expected int) {
	argsLen := len(c.Args())
	if argsLen != expected {
		fmt.Printf("Expected %d args but got %d\nUsage: %s\n", expected, argsLen, c.Command.Usage)
		os.Exit(1)
	}
}

// CompleteFlags returns the possible flags for bash completion.
func CompleteFlags(flags []cli.Flag) {
	for _, flag := range flags {
		flagName := ""
		switch flag.(type) {
		case cli.StringFlag:
			flagName = flag.(cli.StringFlag).Name
		case cli.IntFlag:
			flagName = flag.(cli.IntFlag).Name
		case cli.BoolFlag:
			flagName = flag.(cli.BoolFlag).Name
		default:
			continue
		}
		fmt.Println("--" + flagName)
	}
}
