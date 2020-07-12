package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
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
			log.Fatal("ERROR: Unable to create file!\n", err)
		}
		file.Close()
	}
}

func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		log.Fatal("ERROR: Unable to open file!\n", err)
	}
	defer srcFile.Close()

	fileDst, err := os.Create(dst)
	if err != nil {
		log.Fatal("ERROR: Unable to create file!\n", err)
	}
	defer fileDst.Close()

	if _, err := io.Copy(fileDst, srcFile); err != nil {
		log.Fatal("ERROR: Unable to copy file!\n", err)
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		log.Fatal("ERROR: Unable to read file information!\n", err)
	}

	return os.Chmod(dst, srcInfo.Mode())
}

func CopyDir(src, dst string) {
	info, err := os.Stat(src)
	if err != nil {
		log.Fatal("ERROR: Unable to read source information!\n", err)
	}

	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		log.Fatal("ERROR: Unable to make directory!\n", err)
	}

	dir, err := ioutil.ReadDir(src)
	if err != nil {
		log.Fatal("ERROR: Unable to read source directory!\n", err)
	}

	for _, file := range dir {
		srcPath := path.Join(src, file.Name())
		dstPath := path.Join(dst, file.Name())

		if file.IsDir() {
			CopyDir(srcPath, dstPath)
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				log.Fatal("ERROR: Failed to copy file\n", err)
			}
		}
	}
}
