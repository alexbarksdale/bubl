package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alexbarksdale/bubl/util"
)

// Bubble stores the alias of bubble and a path to a template.
type Bubble struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

// TODO: Figure out '\' situation for Windows

// CreateBubble takes a path to file/directory they wish to create a bubble for
// and an alias which represents their bubble. A bubble is saved to bubbles.json
// in a user's config dir (os.UserConfigDir()).
func CreateBubble(path, alias string) Bubble {
	var bubl Bubble

	bubbles, trie := LoadBubbles()

	if BubbleExist(trie, alias) {
		return bubl
	}

	bubl = Bubble{
		Alias: alias,
		Path:  path,
	}

	bubbles = append(bubbles, bubl)

	b, err := json.Marshal(bubbles)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles!\n", err)
	}

	if err := ioutil.WriteFile(util.BublSavePath, b, 0644); err != nil {
		log.Fatal("ERROR: Failed to save bubble to file!\n")
	}

	fmt.Println("Successfully created bubble.")
	fmt.Println("")
	fmt.Println("Bubble Info")
	fmt.Println("────────────")
	fmt.Printf("Alias: %v\nPath: %v\n\n", alias, path)

	fmt.Println("Generate your bubble with:")
	fmt.Println(GenUsage)
	fmt.Println("")

	return bubl
}
