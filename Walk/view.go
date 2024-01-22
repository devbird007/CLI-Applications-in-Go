package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	root := flag.String("root", ".", "Root directory to start")
	flag.Parse()

	relativeViewer("testdata", *root)
}

func relativeViewer(root, path string) {
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(relDir)
}
