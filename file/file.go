package file

import (
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Alias string `json:"alias"`
	Path  string `json:"path"`
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteFile(filename string, b []byte) {
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Fatal("Failed to create bubl config file!")
	}
}
