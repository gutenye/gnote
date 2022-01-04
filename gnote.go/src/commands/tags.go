package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/gutenye/gnote/gnote.go/src/utils"
	"github.com/yargevad/filepathx"
)

type Tags struct {
	NoteDir string `default:"${noteDir}" help:"Note directory"`

	NoteExtension string `default:"${noteExtension}" help:"Note extension"`

	NoteMarker string `default:"${noteMarker}" help:"Note marker"`

	Output string `default:"${output}" help:"Output tags file"`

	CacheDir string `default:"${cacheDir}" help:"Cache directory"`

	Watch bool `help:"Watch mode"`

	pattern *regexp.Regexp `hidden:""`
}

func (t *Tags) Run() error {
	t.pattern = createPattern(t.NoteMarker)

	t.checkDirs()
	t.createTags()

	if t.Watch {
		t.startWatch()
	}

	return nil
}

// Make sure note directory exists
func (t *Tags) checkDirs() {
	if _, err := os.Stat(t.NoteDir); os.IsNotExist(err) {
		log.Fatalln("Note directory not found.", err)
	}

	if err := os.MkdirAll(t.CacheDir, 0755); err != nil {
		log.Fatalln("Failed to create cache dir.", err)
	}
}

// Create tags file
func (t *Tags) createTags() {
	t.emptyCacheDir()

	notePaths := t.listNotes(t.NoteDir)
	for _, notePath := range notePaths {
		t.createTagsInCache(notePath)
	}

	t.createAllTagsFromCache()

	fmt.Printf("Created %s\n", t.Output)
}

// Empty cache directory
func (t *Tags) emptyCacheDir() {
	err := utils.EmptyDir(t.CacheDir)
	if err != nil {
		log.Fatalln("Failed to empty cache dir", err)
	}
}

func (t *Tags) listNotes(dir string) []string {
	pattern := fmt.Sprintf("%s/**/*%s", dir, t.NoteExtension)
	files, err := filepathx.Glob(pattern)
	if err != nil {
		log.Fatalln("Failed to list notes", err)
	}
	notes := []string{}
	for _, v := range files {
		notePath := utils.RelPath(v, t.NoteDir)
		notes = append(notes, notePath)
	}
	return notes
}

// Create <cache>/a.gnote
func (t *Tags) createTagsInCache(notePath string) {
	tagsContent := t.extractTagsFromFile(notePath)
	t.writeTagsToCache(notePath, tagsContent)
}

// Read <note>/a.gnote and returns tags content
func (t *Tags) extractTagsFromFile(notePath string) string {
	fullNotePath, err := filepath.Abs(filepath.Join(t.NoteDir, notePath))
	if err != nil {
		log.Fatalln("Failed to get absolute path.", err)
	}
	content, err := os.ReadFile(fullNotePath)
	if err != nil {
		log.Fatalln("Failed to note file.", err)
	}
	return extractTagsFromText(string(content), fullNotePath, t.NoteMarker, t.pattern)
}

func (t *Tags) writeTagsToCache(notePath string, tagsContent string) {
	if tagsContent == "" {
		return
	}
	fullNoteCachePath := filepath.Join(t.CacheDir, notePath)
	err := utils.WriteFileWithMkdirAll(fullNoteCachePath, []byte(tagsContent), 0644)
	if err != nil {
		log.Fatalln("Failed to write tags to cache.", err)
	}
}

func (t *Tags) createAllTagsFromCache() {
	paths, err := filepathx.Glob(fmt.Sprintf("%s/**/*%s", t.CacheDir, t.NoteExtension))
	if err != nil {
		log.Fatalln("Failed to list notes.", err)
	}
	allTagsContent := ""
	for _, path := range paths {
		tagsContent, err := os.ReadFile(path)
		if err != nil {
			log.Fatalln("Failed to read tags from cache.", err)
		}
		allTagsContent += "\n" + string(tagsContent)
	}
	allTagsContent = sortTags(allTagsContent)
	result := fmt.Sprintf("!_TAG_FILE_SORTED\t1\n%s", allTagsContent)
	err = utils.WriteFileWithMkdirAll(t.Output, []byte(result), 0644)
	if err != nil {
		log.Fatalln("Failed to write tags.", err)
	}
}

func (t *Tags) startWatch() {
	fmt.Printf("Watching %s\n", t.NoteDir)
	utils.WatchDir(t.NoteDir, func(event fsnotify.Event) {
		switch event.Op {
		case fsnotify.Create, fsnotify.Write:
			t.watchChanged(event.Name)
		case fsnotify.Remove, fsnotify.Remove | fsnotify.Write, fsnotify.Rename, fsnotify.Remove | fsnotify.Rename:
			t.watchRemoved(event.Name)
		}
	})
}

func (t *Tags) watchChanged(fullNotePath string) {
	fmt.Printf("Changed: %s\n", fullNotePath)
	var notePaths []string
	if stat, _ := os.Stat(fullNotePath); stat.IsDir() {
		notePaths = t.listNotes(fullNotePath)
	} else if t.isNoteFile(fullNotePath) {
		notePath := utils.RelPath(fullNotePath, t.NoteDir)
		notePaths = append(notePaths, notePath)
	}
	for _, notePath := range notePaths {
		t.createTagsInCache(notePath)
	}
	t.createAllTagsFromCache()
}

func (t *Tags) watchRemoved(fullNotePath string) {
	notePath := utils.RelPath(fullNotePath, t.NoteDir)
	fmt.Printf("Removed: %s\n", notePath)
	t.removeCache(notePath)
	t.createAllTagsFromCache()
}

func (t *Tags) isNoteFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		log.Fatalln("isNoteFile failed.", err)
	}
	if stat.Mode().IsRegular() && filepath.Ext(path) == t.NoteExtension {
		return true
	}
	return false
}

func (t *Tags) removeCache(notePath string) {
	fullNoteCachePath := filepath.Join(t.CacheDir, notePath)
	os.RemoveAll(fullNoteCachePath)
}

func extractTagsFromText(content string, fullNodePath string, noteMarker string, pattern *regexp.Regexp) string {
	ids := []string{}
	matches := pattern.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		ids = append(ids, match[1:]...)
	}
	tagLines := []string{}
	for _, id := range ids {
		jump := fmt.Sprintf("/%s%s%s", noteMarker, id, noteMarker)
		tagLine := fmt.Sprintf("%s\t%s\t%s", id, fullNodePath, jump)
		tagLines = append(tagLines, tagLine)
	}
	return strings.Join(tagLines, "\n")
}

func createPattern(noteMarker string) *regexp.Regexp {
	marker := regexp.QuoteMeta(noteMarker)
	return regexp.MustCompile(fmt.Sprintf(`%s([^\s]+)%s`, marker, marker))
}

func sortTags(tagsContent string) string {
	lines := strings.Split(tagsContent, "\n")
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}
