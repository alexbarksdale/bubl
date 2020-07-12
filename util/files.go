package util

import (
	"log"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateConfig() {
	if !FileExists("bubbles.json") {
		file, err := os.Create("bubbles.json")
		if err != nil {
			log.Fatal("ERROR: Unable to create file!", err)
		}
		file.Close()
	}
}
