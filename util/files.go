package util

import (
	"io"
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

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal("ERROR: Unable to open file!", err)
	}
	defer srcFile.Close()

	fileDst, err := os.Create(dst)
	if err != nil {
		log.Fatal("ERROR: Unable to create file!", err)
	}
	defer fileDst.Close()

	if _, err := io.Copy(fileDst, srcFile); err != nil {
		log.Fatal("ERROR: Unable to copy file!", err)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		log.Fatal("ERROR: Unable to read file information!", err)
	}

	return os.Chmod(dst, srcInfo.Mode())

}
