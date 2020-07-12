package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type bubble struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

func bubbleExist(b []bubble, alias string) bool {
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
	file, err := ioutil.ReadFile("bubbles.json")
	if err != nil {
		log.Fatal("ERROR: Unable to read Bubble!", err)
	}

	bubbles := []bubble{}
	json.Unmarshal(file, &bubbles)

	if bubbleExist(bubbles, alias) {
		os.Exit(1)
	}

	bubl := bubble{
		Alias: alias,
		Path:  path,
	}

	bubbles = append(bubbles, bubl)

	b, err := json.Marshal(bubbles)
	if err != nil {
		log.Fatal("ERROR: Unable to marshal bubbles.json!", err)
	}

	if err := ioutil.WriteFile("bubbles.json", b, 0644); err != nil {
		log.Fatal("ERROR: Failed to save bubble to file!")
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
