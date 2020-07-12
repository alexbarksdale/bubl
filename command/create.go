package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type bubble struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

// May not be the most efficient when 'bubbles.json' gets really large
// but this will do perfectly fine for now.
func CreateBubl(path, alias string) {
	file, err := ioutil.ReadFile("bubbles.json")
	if err != nil {
		log.Fatal("ERROR: Unable to read Bubble!", err)
	}

	bubl := []bubble{}
	json.Unmarshal(file, &bubl)

	newBubl := bubble{
		Alias: alias,
		Path:  path,
	}

	bubl = append(bubl, newBubl)

	b, err := json.Marshal(bubl)
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
