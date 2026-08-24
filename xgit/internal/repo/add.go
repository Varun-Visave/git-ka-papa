package repo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func AddRepo(path string) error {
	fileInfo, err := os.Stat(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("File does not exists")
		}
		fmt.Printf("An error occurred: %v\n", err)
		return err
	}
	if fileInfo.IsDir() {
		return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				if info.Name() == ".xgit" {
					return filepath.SkipDir
				}
				return nil
			}
			fmt.Println(p)
			fmt.Println(info.Name())
			fmt.Println(info.ModTime())
			fmt.Println(info.Size())
			fmt.Println(info.IsDir())
			return nil
		})
	}

	fmt.Println(path)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	return nil
}
