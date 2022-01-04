package commands

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

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
		log.Fatalf("Note directory not found. %s\n", err)
	}

	if err := os.MkdirAll(t.CacheDir, 0755); err != nil {
		log.Fatalf("Failed to create cache dir. %s\n", err)
	}
}

// Create tags file
func (t *Tags) createTags() {
	t.emptyCacheDir()

	notePaths := t.listNotes()
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
		log.Fatal("Failed to empty cache dir. %s\n", err)
	}
}

func (t *Tags) listNotes() []string {
	pattern := fmt.Sprintf("%s/**/*%s", t.NoteDir, t.NoteExtension)
	files, err := filepathx.Glob(pattern)
	if err != nil {
		log.Fatalf("Failed to list notes. %s\n", err)
	}
	notes := []string{}
	for _, v := range files {
		notePath := strings.TrimPrefix(v, t.NoteDir)
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
		log.Fatalf("Failed to get absolute path. %s\n", err)
	}
	content, err := os.ReadFile(fullNotePath)
	if err != nil {
		log.Fatalf("Failed to read file. %s\n", err)
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
		log.Fatalf("Failed to write tags to cache. %s\n", err)
	}
}

func (t *Tags) createAllTagsFromCache() {
	paths, err := filepathx.Glob(fmt.Sprintf("%s/**/*%s", t.CacheDir, t.NoteExtension))
	if err != nil {
		log.Fatalf("Failed to list notes. %s\n", err)
	}
	allTagsContent := ""
	for _, path := range paths {
		tagsContent, err := os.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to read tags from cache. %s\n", err)
		}
		allTagsContent += "\n" + string(tagsContent)
	}
	allTagsContent = sortTags(allTagsContent)
	result := fmt.Sprintf("!_TAG_FILE_SORTED\t1\n%s", allTagsContent)
	err = utils.WriteFileWithMkdirAll(t.Output, []byte(result), 0644)
	if err != nil {
		log.Fatalf("Failed to write tags: %s\n", err)
	}
}

func (t *Tags) startWatch() {
	fmt.Printf("Watching %s\n", t.NoteDir)

	// 	let (tx, rx) = channel();
	// 	let mut watcher = watcher(tx, Duration::from_secs(1)).unwrap();
	// 	watcher
	// 		.watch(&self.note_dir, RecursiveMode::Recursive)
	// 		.unwrap();
	// 	loop {
	// 		match rx.recv() {
	// 			Err(e) => println!("watch error: {:?}", e),
	// 			Ok(event) => match event {
	// 				DebouncedEvent::Write(path) | DebouncedEvent::Create(path) => self.watch_changed(&path),
	// 				DebouncedEvent::NoticeRemove(path) => self.watch_removed(&path),
	// 				DebouncedEvent::Rename(_, new) => {
	// 					self.watch_changed(&new);
	// 				}
	// 				_ => {}
	// 			},
	// 		}
	// 	}
	// }

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
