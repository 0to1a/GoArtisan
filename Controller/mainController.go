package Controller

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	IsQuiet = false
)

// CheckErr for identified error some procedure.
//
// err value when not nill is mean procedure was error.
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PrintError for show to command line about error on options.
func PrintError(args string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("The '%s' option does not exist.\n", red(args))
}

// PrintDevelop for show the option is not already right now.
func PrintDevelop(args string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("The '%s' option is in develop.\n", blue(args))
}
