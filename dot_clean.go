package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-m" {
		os.Stderr.WriteString("usage: dot_clean [-m] [dir]\n")
		os.Exit(1)
	}

	abs, err := filepath.Abs(os.Args[2])
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	stat, err := os.Stat(abs)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}

	if !stat.IsDir() {
		os.Stderr.WriteString("Failed trying to change dir to " + abs + "\n")
		os.Stderr.WriteString("Bad Pathname: Not a directory\n")
		os.Exit(0)
	}

	err = filepath.Walk(abs, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasPrefix(path[strings.LastIndex(path, string(os.PathSeparator))+1:], "._") {
			return nil
		}

		return os.Remove(path)
	})
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
