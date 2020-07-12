package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		fmt.Printf("Unable to find bubble: '%v'", alias)
		// TODO: Print list bubble info
		return
	}

	b, err := json.Marshal(removedBubble)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles!", err)
	}

	if err := ioutil.WriteFile("bubbles.json", b, 0644); err != nil {
		log.Fatal("ERROR: Failed to save bubble to file!")
	}

	fmt.Printf("Successfully removed bubble: '%v'", alias)
}
