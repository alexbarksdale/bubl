package command

import (
	"log"
	"os"
	"testing"

	"github.com/alexbarksdale/bubl/util"
	"github.com/stretchr/testify/assert"
)

func TestCreateBubble(t *testing.T) {
	// Ensure we have a bubble save file
	util.CreateSave()

	// Used as an example path
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory", err)
	}

	tests := []struct {
		alias string
		path  string
		want  Bubble
	}{
		{
			alias: "Bubble_Test_Red",
			path:  dir,
			want: Bubble{
				Alias: "Bubble_Test_Red",
				Path:  dir,
			},
		},
		{
			alias: "Bubble_Test_Blue",
			path:  dir,
			want: Bubble{
				Alias: "Bubble_Test_Blue",
				Path:  dir,
			},
		},
	}

	for _, test := range tests {
		bubl := CreateBubble(test.path, test.alias)
		assert.IsType(t, test.want, bubl)

		assert.Equal(t, test.want.Alias, bubl.Alias)
		assert.Equal(t, test.want.Path, bubl.Path)
	}

	_, trie := LoadBubbles()
	for _, test := range tests {
		assert.True(t, BubbleExist(trie, test.want.Alias))
	}

	// Clean up created test bubbles
	for _, test := range tests {
		PopBubble(test.alias)
	}
}
