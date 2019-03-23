package Controller

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var (
	IsQuiet = false
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func PrintError(args string) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Printf("The '%s' option does not exist.\n", red(args))
}

func PrintDevelop(args string) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Printf("The '%s' option is in develop.\n", blue(args))
}
