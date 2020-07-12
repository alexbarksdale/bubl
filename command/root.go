package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/alexbarksdale/bubl/util"
)

const (
	bublUsage = `Usage: bubl <command>

{{.Create}}
	Create a new bubl.

{{.Gen}}
	Generate a boilerplate project from a bubl.
`
	CreateUsage = `bubl create <template-path> <bubl-alias>`
	GenUsage    = `bubl gen <bubl-alias>`
)

func invalidArgs(cmd, cmdUsage string, validArg int) {
	fmt.Printf("ERROR: '%v' takes %v arguments, but %v were given.\n\n", cmd, validArg, len(os.Args[2:]))
	fmt.Println(cmdUsage)
	os.Exit(1)
}

func printUsage() {
	type Usage struct {
		Create, Gen string
	}

	var u = []Usage{{CreateUsage, GenUsage}}

	t := template.Must(template.New("bublUsage").Parse(bublUsage))

	for _, usage := range u {
		if err := t.Execute(os.Stdin, usage); err != nil {
			log.Fatalln("ERROR: Failed creating template!", err)
		}
	}
}

func Execute() {
	// A config will only be generated if it doesn't exist.
	util.CreateConfig()

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)

	argLen := len(os.Args)

	if argLen < 2 {
		printUsage()
		os.Exit(1)
	}

	input := os.Args[2:]

	switch os.Args[1] {
	case "create":
		if argLen != 4 {
			invalidArgs("create", CreateUsage, 2)
		}
		createCommand.Parse(input)
	case "gen":
		if argLen != 3 {
			invalidArgs("gen", GenUsage, 1)
		}
		genCommand.Parse(input)
	default:
		fmt.Printf("ERROR: command '%v' does not exist!\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}

	if createCommand.Parsed() {
		CreateBubl(os.Args[2], os.Args[3])
	}

	if genCommand.Parsed() {
		fmt.Println(os.Args[2:])
	}
}
