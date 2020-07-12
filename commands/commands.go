package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"
)

const (
	bublUsage = `Usage: bubl <command>

{{.Create}}
	Create a new bubl.

{{.Gen}}
	Generate a boilerplate project from a bubl.
`
	createUsage = `bubl create <template-path> <bubl-alias>`
	genUsage    = `bubl gen <bubl-alias>`
)

func printUsage() {
	type Usage struct {
		Create, Gen string
	}

	var u = []Usage{{createUsage, genUsage}}

	t := template.Must(template.New("bublUsage").Parse(bublUsage))

	for _, r := range u {
		if err := t.Execute(os.Stdin, r); err != nil {
			log.Fatalln("Failed creating template", err)
		}
	}
}

func invalidArgs(cmd, cmdUsage string, validArg int) {
	fmt.Printf("ERROR: '%v' takes %v arguments, but %v were given.\n\n", cmd, validArg, len(os.Args[2:]))
	fmt.Println(cmdUsage)
	os.Exit(1)
}

func Execute() {
	// bubl commands
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)

	argLen := len(os.Args)

	if argLen < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		if argLen != 4 {
			invalidArgs("create", createUsage, 2)
		}
		createCommand.Parse(os.Args[2:])
	case "gen":
		if argLen != 3 {
			invalidArgs("gen", genUsage, 1)
		}
		genCommand.Parse(os.Args[2:])
	default:
		fmt.Printf("ERROR: command '%v' does not exist!\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	if createCommand.Parsed() {
		fmt.Println(os.Args[2:])
	}

	if genCommand.Parsed() {
		fmt.Println(os.Args[2:])
	}
}
