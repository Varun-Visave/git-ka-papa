package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func StatusRepo() error {
	indexPath := filepath.Join(".xgit", "index")

	entries, err := readIndex(indexPath)
	if err != nil {
		return err
	}

	var untracked []string
	var modified []string

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == ".xgit" {
				return filepath.SkipDir
			}
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		hashString, _, err := hashBlob(content)
		if err != nil {
			return err
		}

		relativePath := normalizePath(path)

		storedHash, exists := entries[relativePath]
		if !exists {
			untracked = append(untracked, relativePath)
		} else if storedHash != hashString {
			modified = append(modified, relativePath)
		}
		return nil
	})

	if err != nil {
		return err
	}

	if len(untracked) == 0 && len(modified) == 0 {
		fmt.Println("nothing to commit, working tree clean")
		return nil
	}

	if len(modified) > 0 {
		fmt.Println("Changes not staged for commit: ")
		for _, path := range modified {
			fmt.Printf(" modified: %s\n", path)
		}
	}
	if len(untracked) > 0 {
		fmt.Println("Untracked file: ")
		for _, path := range untracked {
			fmt.Printf(" %s\n", path)
		}
	}

	return nil
}
