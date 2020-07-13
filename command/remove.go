package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alexbarksdale/bubl/util"
)

func removeBubble(b []Bubble, alias string) ([]Bubble, bool) {
	for i, v := range b {
		if v.Alias == alias {
			// Swap the last bubble with current bubble
			b[len(b)-1], b[i] = b[i], b[len(b)-1]
			// Remove n-1 bubbles
			return b[:len(b)-1], true
		}
	}
	return b, false
}

func PopBubble(alias string) {
	bubbles := LoadBubbles()

	removedBubble, success := removeBubble(bubbles, alias)
	if !success {
		fmt.Printf("Unable to find bubble: '%v'\n", alias)
		// TODO: Print list bubble info
		return
	}

	b, err := json.Marshal(removedBubble)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles!\n", err)
	}

	if err := ioutil.WriteFile(util.BublConfig, b, 0644); err != nil {
		log.Fatal("ERROR: Failed to save bubble to file!\n")
	}

	fmt.Printf("Successfully removed bubble: '%v'\n", alias)
}
