package main

import (
	"os"
	"path/filepath"

	"github.com/alecthomas/kong"
	"github.com/gutenye/gnote/gnote.go/src/commands"
)

var home = os.Getenv("HOME")

func init() {
	if home == "" {
		panic("$HOME is not set")
	}
}

var cli struct {
	Tags commands.Tags `cmd:"" help:"Generate tags file"`
}

func main() {
	ctx := kong.Parse(&cli, kong.Vars{
		"noteDir":       filepath.Join(home, "env/note"),
		"noteExtension": ".gnote",
		"noteMarker":    "∗",
		"output":        filepath.Join(home, "tags"),
		"cacheDir":      filepath.Join(home, ".cache/gnote"),
	})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
