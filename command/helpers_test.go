package command

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func BenchmarkFindBubbleSrc(b *testing.B) {
	// Used as an example path
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory", err)
	}

	alias := "Bubble_Find_Bubble_Src"
	CreateBubble(dir, alias)

	_, trie := LoadBubbles()

	for i := 0; i < b.N; i++ {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			b.ReportAllocs()
			c, _ := FindBubbleSrc(trie, alias)
			fmt.Println(c)
		})
	}

	fmt.Println("Cleaning up benchmark...")
	PopBubble(alias)
}
