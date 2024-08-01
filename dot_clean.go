package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-m" {
		log.Fatal("Usage: dot_clean [-m] [dir]")
	}

	filepath.Walk(filepath.Clean(os.Args[2]), func(path string, info fs.FileInfo, err error) error {
		filename := path[strings.LastIndex(path, "/")+1:]

		if strings.HasPrefix(filename, "._") {
			fmt.Println(path)
		}

		return nil
	})
}
