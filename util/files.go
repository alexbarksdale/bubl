package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

// FileExists validates a given filename is a file.
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// BublSavePath stores the location of bubbles.json.
var BublSavePath string

// CreateSave creates a bubbles.json file in the user's config directory to save bubbles.
func CreateSave() error {
	var saveErr error

	cfg, err := os.UserConfigDir()
	if err != nil {
		saveErr = err
	}

	bublDir := filepath.Join(cfg, "bubl")
	BublSavePath = filepath.Join(bublDir, "bubbles.json")

	if !FileExists(BublSavePath) {
		if err := os.Mkdir(bublDir, 0777); err != nil {
			fmt.Println("It looks like you're missing a 'bubbles.json' file, creating one now...")
			fmt.Printf("Any bubbles you create will be stored here.\n\n")
		}

		file, err := os.Create(BublSavePath)
		if err != nil {
			saveErr = err
		}
		file.Close()
	}
	return saveErr
}

// CopyFile copies a file in a given src to a dst.
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	fileDst, err := os.Create(dst)
	if err != nil {
		fmt.Println("File may already exist in this directory.")
		fmt.Println("Please remove any matching file names and try again.")
		fmt.Println("")
		return err
	}
	defer fileDst.Close()

	if _, err := io.Copy(fileDst, srcFile); err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, srcInfo.Mode())
}

// CopyDir copies a directory in a given src to a dst.
func CopyDir(src, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dst, info.Mode()); err != nil {
		fmt.Println("Directory may already exist.")
		fmt.Println("Please remove any matching directory names and try again.")
		fmt.Println("")
		return err
	}

	dir, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, file := range dir {
		srcPath := path.Join(src, file.Name())
		dstPath := path.Join(dst, file.Name())

		if file.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				log.Fatal("ERROR: Failed to copy directory\n", err)
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				log.Fatal("ERROR: Failed to copy file\n", err)
			}
		}
	}
	return err
}
