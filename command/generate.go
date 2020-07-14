package command

import (
	"fmt"
	"log"
	"os"

	"github.com/alexbarksdale/bubl/util"
)

func GenBubble(alias string) {
	_, trie := LoadBubbles()

	src, found := FindBubbleSrc(trie, alias)
	if !found {
		fmt.Printf("Bubble '%v' does not exist, please try another alias.\n\n", alias)
		// TODO: Show available bubbles
		return
	}

	info, err := os.Stat(src)
	if err != nil {
		fmt.Printf("The path to '%v' is not valid.\n\n%v", alias, err)
		return
	}

	fmt.Printf("Generating bubble: '%v'\n", alias)
	fmt.Printf("Large files may take a few seconds.\n\n")

	if info.IsDir() {
		if err := util.CopyDir(src, info.Name()); err != nil {
			log.Fatal("ERROR: Failed to copy directory\n", err)
		}
	} else {
		if err := util.CopyFile(src, info.Name()); err != nil {
			log.Fatal("ERROR: Failed to copy file\n", err)
		}
	}

	fmt.Printf("Successfully generated bubble '%v' to your current directory.\n\n", alias)
}
