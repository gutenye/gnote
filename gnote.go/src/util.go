package main

import (
	"os"
	"path/filepath"
  "fmt"
  "reflect"
  "strings"
)

// Support ~
func AbsWithExtend(path string) (string, error) {
	if path[0] == '~' {
		path = os.Getenv("HOME")+path[1:]
	}

  ret, err := filepath.Abs(path)
  return ret, err
}

// path
func IsExist(path string) bool {
	f, e := os.Open(path)
	if f != nil {
		f.Close()
		return true
	} else {
		return os.IsExist(e)
	}
	return false
}

func IsNotExist(path string) bool {
  return ! IsExist(path)
}

// Empty a directory
func EmptyAll(dir string) error {
	f, err := os.Open(dir)
	if err != nil { return err }
	names, err := f.Readdirnames(-1)
	f.Close()
	for _, name := range names {
		err := os.RemoveAll(filepath.Join(dir, name))
		if err != nil { return err }
	}
	return nil
}

// Usage
//
//    import . "github.com/GutenYe/tagen.go/pd"
//
//    Pd(1, 2, 3)         -> "1 2 3"
//    Pd("%d %d", 1, 2)   -> "1 2"
//
// import (
//   "fmt"
//   "reflect"
//   "strings"
// )

func Pd(values ...interface{}) (n int, err error) {
	if len(values) <= 0 { return }

	v := reflect.ValueOf(values[0])
	if v.Kind() == reflect.String && strings.Contains(v.String(), "%") {
		n, err = fmt.Printf(v.String()+"\n", values[1:]...)
	} else {
		n, err = fmt.Println(values...)
	}

	return n, err
}
