package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/alexbarksdale/bubl/util"
)

const (
	bublUsage = `Usage: bubl <command>

{{.Create}}
	Create a new bubble template.

{{.Gen}}
	Generate a boilerplate file/project from a bubble.

{{.Pop}}
	Remove a bubble template.
`
	CreateUsage = `bubl create <template-path> <bubl-alias>`
	GenUsage    = `bubl gen <bubl-alias>`
	PopUsage    = `bubl pop <bubl-alias>`
)

func invalidArgs(cmd, cmdUsage string, validArg int) {
	if validArg == 1 {
		fmt.Printf("ERROR: '%v' takes %v argument, but 1 was given.\n\n", cmd, validArg)
	} else {
		fmt.Printf("ERROR: '%v' takes %v arguments, but %v were given.\n\n", cmd, validArg, len(os.Args[2:]))
	}
	fmt.Println(cmdUsage)
	os.Exit(1)
}

func printUsage() {
	type Usage struct {
		Create, Gen, Pop string
	}

	var u = []Usage{{CreateUsage, GenUsage, PopUsage}}

	t := template.Must(template.New("bublUsage").Parse(bublUsage))

	for _, usage := range u {
		if err := t.Execute(os.Stdin, usage); err != nil {
			log.Fatalln("ERROR: Failed creating template!", err)
		}
	}
}

func LoadBubbles() []Bubble {
	file, err := ioutil.ReadFile("bubbles.json")
	if err != nil {
		log.Fatal("ERROR: Unable to read Bubble!", err)
	}

	bubbles := []Bubble{}
	json.Unmarshal(file, &bubbles)

	return bubbles
}

func Execute() {
	// A config will only be generated if it doesn't exist.
	util.CreateConfig()

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	popCommand := flag.NewFlagSet("remove", flag.ExitOnError)

	argLen := len(os.Args)

	if argLen < 2 {
		printUsage()
		return
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
	case "pop":
		if argLen != 3 {
			invalidArgs("pop", PopUsage, 1)
		}
		popCommand.Parse(input)
	default:
		fmt.Printf("ERROR: command '%v' does not exist!\n\n", os.Args[1])
		printUsage()
		return
	}

	if createCommand.Parsed() {
		CreateBubl(os.Args[2], os.Args[3])
	}

	if genCommand.Parsed() {
		fmt.Println(os.Args[2:])
	}

	if popCommand.Parsed() {
		PopBubble(os.Args[2])
	}
}
