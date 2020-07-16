package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alexbarksdale/bubl/util"
)

func PopBubble(alias string) bool {
	bubbles, _ := LoadBubbles()

	removedBubble, success := RemoveBubble(bubbles, alias)
	if !success {
		fmt.Printf("Unable to find bubble: '%v'\n\n", alias)

		fmt.Println("View your bubbles with: ")
		fmt.Println(ListUsage)
		fmt.Println("")
		return false
	}

	b, err := json.Marshal(removedBubble)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles!\n", err)
	}

	if err := ioutil.WriteFile(util.BublSavePath, b, 0644); err != nil {
		log.Fatal("ERROR: Failed to save bubble to file!\n", err)
	}

	fmt.Printf("Successfully removed bubble: '%v'\n\n", alias)
	return true
}
