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
	Generate a template from a bubble.

{{.Pop}}
	Remove a bubble template.

`
	CreateUsage = `bubl create <template-path> <bubl-alias>`
	GenUsage    = `bubl gen <bubl-alias>`
	PopUsage    = `bubl pop <bubl-alias>`
)

func printUsage() {
	type Usage struct {
		Create, Gen, Pop string
	}

	var u = []Usage{{CreateUsage, GenUsage, PopUsage}}

	t := template.Must(template.New("bublUsage").Parse(bublUsage))

	for _, usage := range u {
		if err := t.Execute(os.Stdout, usage); err != nil {
			log.Fatalln("ERROR: Failed creating template!", err)
		}
	}
}

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

func LoadBubbles() []Bubble {
	file, err := ioutil.ReadFile(util.BublSavePath)
	if err != nil {
		log.Fatal("ERROR: Unable to read bubbles!", err)
	}

	bubbles := []Bubble{}
	json.Unmarshal(file, &bubbles)

	return bubbles
}

func Execute() {
	// A config will only be generated if it doesn't exist.
	if err := util.CreateSave(); err != nil {
		log.Fatal("ERROR: Failed to create bubl save file!", err)
	}

	createCommand := flag.NewFlagSet("create", flag.ExitOnError)
	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	popCommand := flag.NewFlagSet("remove", flag.ExitOnError)

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	input := os.Args[2:]
	inputLen := len(input)

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
	default:
		fmt.Printf("Command '%v' does not exist!\n\n", os.Args[1])
		printUsage()
		return
	}

	if createCommand.Parsed() {
		CreateBubl(os.Args[2], os.Args[3])
	}

	if genCommand.Parsed() {
		GenBubble(os.Args[2])
	}

	if popCommand.Parsed() {
		PopBubble(os.Args[2])
	}
}
