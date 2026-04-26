package crawler

import (
	"os"
	"path/filepath"
)

func FindFiles(rootDir string) (filePaths []string, err error) {
	err = filepath.WalkDir(rootDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return filePaths, err
	}
	return filePaths, nil
}
