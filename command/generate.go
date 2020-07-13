package command

import (
	"fmt"
	"log"
	"os"

	"github.com/alexbarksdale/bubl/util"
)

func findBubbleSrc(b []Bubble, alias string) (string, bool) {
	var src string

	for _, v := range b {
		if v.Alias == alias {
			src = v.Path
			return src, true
		}
	}
	return src, false
}

func GenBubble(alias string) {
	bubbles := LoadBubbles()

	src, found := findBubbleSrc(bubbles, alias)
	if !found {
		fmt.Printf("Bubble '%v' does not exist, please try another alias.\n\n", alias)
		// TODO: Show available bubbles
		return
	}

	info, err := os.Stat(src)
	if err != nil {
		log.Fatal("ERROR: Unable to read source information!\n", err)
	}

	fmt.Printf("Generating bubble: '%v'\n", alias)
	fmt.Println("Large files may take a few seconds.")
	fmt.Println("")

	if info.IsDir() {
		util.CopyDir(src, info.Name())
	} else {
		util.CopyFile(src, info.Name())
	}

	fmt.Printf("Successfully generated bubble: '%v'\n", alias)
	fmt.Println("")
}
