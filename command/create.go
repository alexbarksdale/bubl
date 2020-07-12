package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Bubble struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

func bubbleExist(b []Bubble, alias string) bool {
	for _, v := range b {
		if v.Alias == alias {
			fmt.Printf("Bubble '%v' already exists, please use another alias.\n\n", alias)
			fmt.Println("Existing Bubble")
			fmt.Println("────────────────")
			fmt.Printf("Alias: %v\nPath: %v\n\n", v.Alias, v.Path)
			return true
		}
	}
	return false
}

// May not be the most efficient when 'bubbles.json' gets really large
// but this will do perfectly fine for now.
func CreateBubl(path, alias string) {
	bubbles := LoadBubbles()

	if bubbleExist(bubbles, alias) {
		return
	}

	bubl := Bubble{
		Alias: alias,
		Path:  path,
	}

	bubbles = append(bubbles, bubl)

	b, err := json.Marshal(bubbles)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles!\n", err)
	}

	if err := ioutil.WriteFile("bubbles.json", b, 0644); err != nil {
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
}
