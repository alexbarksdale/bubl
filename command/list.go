package command

import "fmt"

func ListBubbles() {
	bubbles, _ := LoadBubbles()

	fmt.Println("Your Bubbles")
	fmt.Println("───────────── ")
	for _, bubl := range bubbles {
		fmt.Printf("Alias: %v\n", bubl.Alias)
		fmt.Printf("Path: %v\n", bubl.Path)
		fmt.Println("")
	}
}
