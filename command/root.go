package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/alexbarksdale/bubl/util"
)

// Description of all the command usages.
const (
	bublUsage = `Usage: bubl <command>

{{.Create}}
	Create a bubble by providing a path to your template 
	and an alias to identify your bubble.

{{.Gen}}
	Generate a template from a bubble to your current directory.

{{.Pop}}
	Remove a bubble template.

{{.List}}
	List out created bubbles.

`
	CreateUsage = `bubl create <template-path> <bubl-alias>`
	GenUsage    = `bubl gen <bubl-alias>`
	PopUsage    = `bubl pop <bubl-alias>`
	ListUsage   = `bubl list`
)

// displayUsage populates the bublUsage template.
func displayUsage() {
	type Usage struct {
		Create, Gen, Pop, List string
	}

	var u = []Usage{{CreateUsage, GenUsage, PopUsage, ListUsage}}

	t := template.Must(template.New("bublUsage").Parse(bublUsage))

	for _, usage := range u {
		if err := t.Execute(os.Stdout, usage); err != nil {
			log.Fatalln("ERROR: Failed creating template!", err)
		}
	}
}

// invalidArgs is a helper function that sends an invalid amount of arguments message to the user.
func invalidArgs(cmd, cmdUsage string, validArg, argsGiven int) {
	if argsGiven == 1 {
		fmt.Printf("ERROR: '%v' takes %v argument, but 1 was given.\n\n", cmd, validArg)
	} else {
		fmt.Printf("ERROR: '%v' takes %v arguments, but %v were given.\n\n", cmd, validArg, argsGiven)
	}
	fmt.Println(cmdUsage)
	fmt.Println("")
	os.Exit(1)
}

// Execute parses a command line argument and sends off the corresponding command.
func Execute() {
	// A config will only be generated if it doesn't exist.
	if err := util.CreateSave(); err != nil {
		log.Fatal("ERROR: Failed to create bubble save file!", err)
	}

	// Commands
	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	popCommand := flag.NewFlagSet("remove", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list", flag.ExitOnError)

	if len(os.Args) < 2 {
		displayUsage()
		return
	}

	input := os.Args[2:]
	inputLen := len(input)

	// Check if the argument given is a valid command.
	switch os.Args[1] {
	case "create":
		if inputLen != 2 {
			invalidArgs("create", CreateUsage, 2, inputLen)
		}
		createCommand.Parse(input)
	case "gen":
		if inputLen != 1 {
			invalidArgs("gen", GenUsage, 1, inputLen)
		}
		genCommand.Parse(input)
	case "pop":
		if inputLen != 1 {
			invalidArgs("pop", PopUsage, 1, inputLen)
		}
		popCommand.Parse(input)
	case "list":
		listCommand.Parse(input)
	default:
		fmt.Printf("Command '%v' does not exist!\n\n", os.Args[1])
		displayUsage()
		return
	}

	if createCommand.Parsed() {
		CreateBubble(os.Args[2], os.Args[3])
	}

	if genCommand.Parsed() {
		GenBubble(os.Args[2])
	}

	if popCommand.Parsed() {
		PopBubble(os.Args[2])
	}

	if listCommand.Parsed() {
		ListBubbles()
	}
}
