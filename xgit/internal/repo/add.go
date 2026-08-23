package repo

import (
	"fmt"
	"os"
	"path/filepath"
)

func AddRepo(path string) error {
	return filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		fmt.Println(p)
		return nil
	})
}
