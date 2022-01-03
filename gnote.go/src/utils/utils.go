package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func EmptyDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		err := os.RemoveAll(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func PrintfAndExit(format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(1)
}

func WriteFileWithMkdirAll(path string, data []byte, perm fs.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, perm)
}
