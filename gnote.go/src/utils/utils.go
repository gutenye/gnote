package utils

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
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

func WriteFileWithMkdirAll(path string, data []byte, perm fs.FileMode) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, perm)
}

func WatchDir(dir string, handleEvent func(fsnotify.Event)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln("Failed to start watcher.", err)
	}
	defer watcher.Close()

	err = filepath.WalkDir(dir, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if dirEntry.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})

	if err != nil {
		log.Fatalln("Failed to walk dir.", err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					if stat, _ := os.Stat(event.Name); stat.IsDir() {
						watcher.Add(event.Name)
					}
				}
				handleEvent(event)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("watch error.", err)
			}
		}
	}()

	<-done
}

func RelPath(path, base string) string {
	result, err := filepath.Rel(base, path)
	if err != nil {
		log.Fatalln("relpath failed.", err)
	}
	return result
}
