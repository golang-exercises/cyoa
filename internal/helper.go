package internal

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func ParseFlags() ([]byte, bool) {
	var storyPath string
	var cliMode bool
	_, b, _, _ := runtime.Caller(0)

	Root := filepath.Join(filepath.Dir(b), "..")

	flag.StringVar(&storyPath, "story", filepath.Join(Root, "assets/story.json"), "Path to story file")
	flag.BoolVar(&cliMode, "cli", false, "Whether the app should run in the browser or command line.")

	flag.Parse()

	return ReadFile(storyPath), cliMode
}

func ReadFile(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return file
}
