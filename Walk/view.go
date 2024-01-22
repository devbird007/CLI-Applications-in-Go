package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := flag.String("root", ".", "Root directory to start")
	flag.Parse()

	if err := runs(*root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func runs(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		relativeViewer(root, "dir2")
		fmt.Println(filepath.Dir("go"))
		return err
	})
}

func relativeViewer(root, path string) {
	relDir, err := filepath.Rel(root, filepath.Dir(path))
	if err != nil {
		fmt.Errorf("%s", err)
	}
	fmt.Println(relDir)
}
