package main 

import (
	"path/filepath"
	"io/ioutil"
	"regexp"
	"os"
	"fmt"
	"sort"
	"strings"
)

// Generate a cache tag from a file
func GenerateCacheTags(relName, dir, cache, mark string) error {
	file := filepath.Join(dir, relName)
	absFile, err := filepath.Abs(file)
	if err != nil { return err }
	cacheFile := filepath.Join(cache, relName)
	tags := ""
	quotedMark := regexp.QuoteMeta(mark)

	data, err := ioutil.ReadFile(file)
	if err != nil { return err }

	re, err := regexp.Compile(fmt.Sprintf(`%s([^\n ]+)%s`, quotedMark, quotedMark))
	if err != nil { return err }
	matches := re.FindAllStringSubmatch(string(data), -1)
	for _, m := range matches {
		id := m[1]
		tagPattern := fmt.Sprintf("/%s%s%s", mark, id, mark) 
		tags += fmt.Sprintf("%s\t%s\t%s\n", id, absFile, tagPattern)
	}

	err = os.MkdirAll(filepath.Dir(cacheFile), 0755)
	if err != nil { return err }
	err = ioutil.WriteFile(cacheFile, []byte(tags), 0644)
	if err != nil { return err }
	return nil
}

// Concat all per-file tags to the final tags file.
func ConcatCacheTags(cache, usertags, output string) error {
	data, _ := ioutil.ReadFile(usertags)
	tag := string(data)

	err := filepath.Walk(cache, func(p string, i os.FileInfo, e error) error {
		if e != nil { return e }
		if i.IsDir() { return nil }

		data, e := ioutil.ReadFile(p)
		if e != nil { return e }
		tag += string(data)

		return nil
	})
	if err != nil { return err }

	tags := strings.Split(tag, "\n")
	sort.Strings(tags)
	tag = strings.TrimSpace(strings.Join(tags, "\n"))

	tag = fmt.Sprintf(`!_TAG_FILE_SORTED	1

%s`, tag)  
	err = ioutil.WriteFile(output, []byte(tag), 0644)
	if err != nil { return err }
	return nil
}
