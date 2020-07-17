package command

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPopBubble(t *testing.T) {
	// Used as an example path
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get working directory", err)
	}

	tests := []struct {
		alias string
		want  bool
	}{
		{
			alias: "Bubble_Test_True",
			want:  true,
		},
		{
			alias: "Bubble_Test_False",
			want:  false,
		},
	}

	// Only create the first test bubble
	CreateBubble(dir, tests[0].alias)

	for _, test := range tests {
		ok := PopBubble(test.alias)
		if ok {
			assert.True(t, test.want, ok)
		} else {
			// If bubble doesn't exist
			assert.False(t, test.want, ok)
		}

	}
}
