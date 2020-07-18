package command

import "fmt"

// TODO: Show another message for having no bubbles

// ListBubbles iterates over bubbles.json and prints out each bubble.
func ListBubbles() {
	bubbles, _ := loadBubbles()
	if len(bubbles) > 1 {
		fmt.Println("Your Bubbles")
		fmt.Println("───────────── ")
		for _, bubl := range bubbles {
			fmt.Printf("Alias: %v\n", bubl.Alias)
			fmt.Printf("Path: %v\n\n", bubl.Path)
		}
	} else {
		fmt.Println("You don't have any bubbles!")
		fmt.Println("")
		fmt.Println("Create one with:")
		fmt.Println(CreateUsage)
		fmt.Println("")
	}
}
