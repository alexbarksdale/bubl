package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
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

	OPTIONS:
	-bundle
		Bundle together an arbitrary amount of bubbles to generate.

{{.Pop}}
	Remove a bubble template.

{{.List}}
	List out created bubbles.
`
	CreateUsage = `bubl create <template-path> <bubl-alias>`
	GenUsage    = `bubl gen (options) <bubl-alias>`
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

	genBundlePtr := genCommand.Bool("bundle", false, "Buble up bubbles")

	if len(os.Args) < 2 {
		displayUsage()
		return
	}

	input := os.Args[2:]
	inputLen := len(input)

	// Check if the argument given is a valid command.
	switch os.Args[1] {
	case "create":
		createCommand.Parse(input)
	case "gen":
		genCommand.Parse(input)
	case "pop":
		popCommand.Parse(input)
	case "list":
		listCommand.Parse(input)
	default:
		fmt.Printf("Command '%v' does not exist!\n\n", os.Args[1])
		displayUsage()
		return
	}

	if createCommand.Parsed() {
		if inputLen != 2 {
			invalidArgs("create", CreateUsage, 2, inputLen)
		}
		CreateBubble(os.Args[2], os.Args[3])
	}

	if genCommand.Parsed() {
		switch {
		case *genBundlePtr && inputLen < 2:
			fmt.Println("You must provide bubbles to bundle.")
			fmt.Println("")
		case *genBundlePtr:
			var wg sync.WaitGroup

			wg.Add(len(os.Args[3:]))

			for _, bubl := range os.Args[3:] {
				go func(b string) {
					GenBubble(b)
					wg.Done()
				}(bubl)
			}
			wg.Wait()
		case inputLen == 1:
			GenBubble(os.Args[2])
		default:
			invalidArgs("gen", GenUsage, 1, inputLen)
		}
	}

	if popCommand.Parsed() {
		if inputLen != 1 {
			invalidArgs("pop", PopUsage, 1, inputLen)
		}
		PopBubble(os.Args[2])
	}

	if listCommand.Parsed() {
		ListBubbles()
	}
}
