package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/alexbarksdale/bubl/util"
	"github.com/derekparker/trie"
)

// invalidArgs is a helper function that sends an invalid amount of arguments message to the user.
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

// loadBubbles reads the BublSavePath and unmarshals the bubble.json file.
// It also generates a trie tree of all the aliases for improved search speed.
func loadBubbles() ([]Bubble, *trie.Trie) {
	file, err := ioutil.ReadFile(util.BublSavePath)
	if err != nil {
		log.Fatal("ERROR: Unable to load bubbles! ", err)
	}

	bubbles := []Bubble{}
	json.Unmarshal(file, &bubbles)

	t := trie.New()
	for _, bubl := range bubbles {
		alias := strings.ToLower(bubl.Alias)
		t.Add(alias, bubl.Path)
	}
	return bubbles, t
}

// bubbleExist validates a bubble by searching through a trie with a given alias.
func bubbleExist(t *trie.Trie, alias string) bool {
	node, found := t.Find(strings.ToLower(alias))
	if found {
		meta := node.Meta()
		fmt.Printf("Bubble '%v' already exists, please try another alias.\n\n", alias)
		fmt.Println("Existing Bubble")
		fmt.Println("────────────────")
		fmt.Printf("Alias: %v\nPath: %v\n\n", alias, meta)
		return true
	}
	return false
}

// findBubbleSrc locates a file/directory of a bubble by searching through a trie with a given alias.
func findBubbleSrc(t *trie.Trie, alias string) (string, bool) {
	var src string

	searchTerm := strings.ToLower(alias)

	node, found := t.Find(searchTerm)
	if found {
		meta := node.Meta()
		src = fmt.Sprintf("%v", meta)
		return src, found
	}
	return src, false
}

// removeBubble iterates over a []Bubble to find a matching bubble alias
// and removes it by swapping it with the last item and returning n-1 items.
func removeBubble(b []Bubble, alias string) ([]Bubble, bool) {
	for i, v := range b {
		if strings.ToLower(v.Alias) == strings.ToLower(alias) {
			b[len(b)-1], b[i] = b[i], b[len(b)-1]
			return b[:len(b)-1], true
		}
	}
	return b, false
}
