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

	// Options
	genBundlePtr := genCommand.Bool("bundle", false, "Bundle together an arbitrary amount of bubbles to generate.")

	if len(os.Args) < 2 {
		displayUsage()
		return
	}

	// Check if the input given is a valid command.
	input := os.Args[2:]

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

	// TODO: Consider moving parsed logic separately too clean up this file.

	inputLen := len(input)

	if createCommand.Parsed() {
		path := os.Args[2]
		alias := os.Args[3]
		if inputLen != 2 {
			invalidArgs("create", CreateUsage, 2, inputLen)
		}
		CreateBubble(path, alias)
	}

	if genCommand.Parsed() {
		switch {
		case *genBundlePtr && inputLen < 2:
			fmt.Printf("You must provide bubbles to bundle.\n\n")
		case *genBundlePtr:
			var wg sync.WaitGroup

			bubbles := os.Args[3:]

			wg.Add(len(bubbles))

			for _, bubl := range bubbles {
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
