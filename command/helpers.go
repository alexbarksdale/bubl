package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/alexbarksdale/bubl/util"
	"github.com/derekparker/trie"
)

func LoadBubbles() ([]Bubble, *trie.Trie) {
	file, err := ioutil.ReadFile(util.BublSavePath)
	if err != nil {
		log.Fatal("ERROR: Unable to load bubbles! ", err)
	}

	bubbles := []Bubble{}
	json.Unmarshal(file, &bubbles)

	t := trie.New()
	for _, bubl := range bubbles {
		t.Add(bubl.Alias, bubl.Path)
	}
	return bubbles, t
}

func BubbleExist(t *trie.Trie, alias string) bool {
	node, found := t.Find(alias)
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

func FindBubbleSrc(t *trie.Trie, alias string) (string, bool) {
	var src string

	node, found := t.Find(alias)
	if found {
		meta := node.Meta()
		src = fmt.Sprintf("%v", meta)
		return src, found
	}
	return src, false
}

func RemoveBubble(b []Bubble, alias string) ([]Bubble, bool) {
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
