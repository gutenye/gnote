package main

import (
	"testing"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
	"strings"
	. "./globals/ui"
	"log"
)

func init() {
	Ui = log.New(os.Stdout, "", 0)
}

var file, _ = filepath.Abs("testtmp/note/finance/hello.gnote")
var note=strings.TrimLeft(`
*guten*
hello world
*tag*
`, "\n")
var tags = strings.TrimLeft(fmt.Sprintf(`
guten	%s	/*guten*
tag	%s	/*tag*
`, file, file), "\n")

func initTest() {
	os.RemoveAll("testtmp")
	os.MkdirAll("testtmp/note/finance", 0755)
	os.MkdirAll("testtmp/cache", 0755)
	ioutil.WriteFile("testtmp/note/finance/hello.gnote", []byte(note), 0644)
}

func TestGenerateCacheTags(t *testing.T) {
	initTest()
	err := GenerateCacheTags("finance/hello.gnote", "testtmp/note", "testtmp/cache", "*")
	if err != nil { t.Fatal(err) }

	data, err := ioutil.ReadFile("testtmp/cache/finance/hello.gnote")
	if err != nil { t.Fatal(err) }
	text, tags := strings.TrimSpace(string(data)), strings.TrimSpace(tags)
	if text != tags {
		t.Fatalf("generated cache tags don't match. \n%s\n---\n%s", text, tags)
	}
}

func TestConcatCacheTags(t *testing.T) {
	e := ConcatCacheTags("testdata/cache", "", "testtmp/tags")
	if e != nil { t.Fatal(e) }
	data, e := ioutil.ReadFile("testtmp/tags")
	if e != nil { t.Fatal(e) }

	data2, e := ioutil.ReadFile("testdata/tags")
	if e != nil { t.Fatal(e) }
	tags, expect := strings.TrimSpace(string(data)), strings.TrimSpace(string(data2))
	if tags != expect {
		t.Fatalf("concat cache tags don't match. \n%s\n---\n%s", tags, expect)
	}
}
