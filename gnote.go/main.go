package main

import (
	. "github.com/gutenye/gnote/globals/rc"
	. "github.com/gutenye/gnote/globals/ui"
	"github.com/ogier/pflag"
	"io/ioutil"
  "gopkg.in/yaml.v1"
	"log"
	"os"
	"path/filepath"
)

const VERSION = "1.0.0"

var homeRc = filepath.Join(os.Getenv("HOME"), ".gnoterc")
var homeConfig = filepath.Join(os.Getenv("HOME"), ".gnote")
var USAGE = `$ gnote <cmd> [options]

COMMAND:
  tags                     # generate tags file
  watch                    # watch note directory

OPTIONS
  -v, --version
  -h, --help
      --dir                # note directory
      --output             # output 'tags' file path
      --mark               # mark character
      --cache              # per-file tags cache directory
`

func main() {
	Ui = log.New(os.Stdout, "", 0)
	pflag.Usage = func() {
		Ui.Print(USAGE)
	}
	var version = pflag.BoolP("version", "v", false, "print version number")
	var dir = pflag.StringP("dir", "", "", "note directory")
	var output = pflag.StringP("output", "", "", "output file")
	var mark = pflag.StringP("mark", "", "", "mark character")
	var cache = pflag.StringP("cache", "", "", "cahce directory")
	pflag.Parse()
	if *version {
		Ui.Printf("gnote %s", VERSION)
		os.Exit(0)
	}

	if IsExist(homeRc) {
		d, e := ioutil.ReadFile(homeRc)
		if e != nil {
			Ui.Fatal(e)
		}
		yaml.Unmarshal(d, &Rc)
	}
	Rc.Cache = "~/.cache/gnote"
	if *dir != "" {
		Rc.Dir = *dir
	}
	if *output != "" {
		Rc.Output = *output
	}
	if *cache != "" {
		Rc.Cache = *cache
	}
	if *mark != "" {
		Rc.Mark = *mark
	}

	var err error
	Rc.Dir, err = AbsWithExtend(Rc.Dir)
	if err != nil {
		Ui.Panic(Rc.Dir)
	}
	if IsNotExist(Rc.Dir) {
		Ui.Printf("--dir `%s` does not exists", Rc.Dir)
		os.Exit(1)
	}
	Rc.Dir, _ = filepath.EvalSymlinks(Rc.Dir)
	if err != nil {
		Ui.Panic(err)
	}

	Rc.Output, err = AbsWithExtend(Rc.Output)
	if err != nil {
		Ui.Panic(err)
	}

	Rc.Cache, _ = AbsWithExtend(Rc.Cache)
	if err != nil {
		Ui.Panic(err)
	}
	if IsNotExist(Rc.Cache) {
		err := os.MkdirAll(Rc.Cache, 0755)
		if err != nil {
			Ui.Panic(err)
		}
	}

	Rc.Usertags = filepath.Join(homeConfig, "tags")

	switch pflag.Arg(0) {
	case "tags":
		Tags()
	case "watch":
		Tags()
		Watch()
	default:
		pflag.Usage()
	}
}
