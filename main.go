package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/0to1a/GoArtisan/Controller"
	"github.com/fatih/color"
)

const (
	version = "1.0.0"
)

// PrintStatus to print information or command list for GoArtisan.
func PrintStatus() {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("Golang Artisan (GoArt) %s\n", green(version))
	fmt.Println()

	fmt.Printf("%s\n", yellow("Usage:"))
	fmt.Printf("  %s\n", "command [options] [arguments]")
	fmt.Println()

	fmt.Printf("%s\n", yellow("Options:"))
	fmt.Printf("  %s %s\n", green("-h, --help         "), "Display this help message")
	fmt.Printf("  %s %s\n", green("-q, --quiet        "), "Do not output any message")
	fmt.Printf("  %s %s\n", green("-V, -v, --version  "), "Display this application version")
	fmt.Println()

	fmt.Printf("%s\n", yellow("Available commands:    "))
	fmt.Printf("  %s %s\n", green("create             "), "Create a new package project")
	fmt.Printf("  %s %s\n", green("env                "), "Display the current framework environment")
	fmt.Printf("  %s %s\n", green("help               "), "Help lists commands")
	fmt.Printf("  %s %s\n", green("migrate            "), "Run the database migrations")

	fmt.Printf(" %s\n", yellow("auth"))
	fmt.Printf("  %s %s\n", green("auth:clear-resets  "), "Flush expired password reset tokens")
	fmt.Printf("  %s %s\n", green("auth:generate      "), "Set the application key")
	fmt.Printf("  %s %s\n", green("auth:reset         "), "Flush expired password reset tokens")
	fmt.Printf(" %s\n", yellow("docs"))
	fmt.Printf("  %s %s\n", green("docs:generate      "), "Create documentation for all function")
	fmt.Printf("  %s %s\n", green("docs:repair        "), "Repair docs to clean & beautiful to read")
	fmt.Printf(" %s\n", yellow("key"))
	fmt.Printf("  %s %s\n", green("key:generate       "), "Set the application key")
	fmt.Printf(" %s\n", yellow("make"))
	fmt.Printf("  %s %s\n", green("make:controller    "), "Create a new Controller service")
	fmt.Printf("  %s %s\n", green("make:model         "), "Create a new Model and Interface")
	fmt.Printf("  %s %s\n", green("make:view          "), "Create a new View")
	fmt.Printf("  %s %s\n", green("make:migration     "), "Build migration from model")
	fmt.Printf(" %s\n", yellow("route"))
	fmt.Printf("  %s %s\n", green("route:list         "), "List all registered routes")
}

func main() {
	var (
		arg     string
		argNext []string
	)
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			arg = os.Args[i]
			if arg[:1] == "-" {
				// Options List
				switch arg {
				case "-h", "--help":
					PrintStatus()
					break
				case "-q", "--quiet":
					Controller.IsQuiet = true
					break
				case "-v", "-V", "--version":
					fmt.Printf("Golang Artisan (GoArtisan) %s\n", green(version))
					return
				default:
					Controller.PrintError(arg)
				}
			} else {
				// Command List
				argNext = os.Args[i:]
				if strings.Contains(arg, ":") {
					arg = arg[:strings.Index(arg, ":")]
				}
				switch arg {
				case "create":
					if len(argNext) > 1 {
						Controller.CreateFolder(argNext)
					} else {
						fmt.Printf("Error:	%s\n", red("Missing 'name' arguments"))
					}
					return
				case "env":
					Controller.PrintDevelop(arg)
					return
				case "help":
					if len(argNext) > 1 {
						switch argNext[1] {
						case "create":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "env":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "migrate":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "auth":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "docs":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "key":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "make":
							Controller.PrintDevelop("help" + argNext[1])
							return
						case "route":
							Controller.PrintDevelop("help" + argNext[1])
							return
						default:
							Controller.PrintError(arg)
						}
					} else {
						PrintStatus()
					}
					return
				case "migrate":
					Controller.PrintDevelop(arg)
					return
				case "auth":
					Controller.PrintDevelop(arg)
					return
				case "docs":
					Controller.PrintDevelop(arg)
					return
				case "key":
					Controller.PrintDevelop(arg)
					return
				case "make":
					Controller.PrintDevelop(arg)
					return
				case "route":
					Controller.PrintDevelop(arg)
					return
				default:
					Controller.PrintError(arg)
					return
				}
			}
		}
	} else {
		PrintStatus()
	}
}
