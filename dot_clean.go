package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func exit(errString string) {
	os.Stderr.WriteString(errString + "\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 || os.Args[1] != "-m" {
		exit("usage: dot_clean [-m] [dir]")
	}

	abs, err := filepath.Abs(os.Args[2])
	if err != nil {
		exit(err.Error())
	}

	stat, err := os.Stat(abs)
	if err != nil {
		exit(err.Error())
	}

	if !stat.IsDir() {
		exit("Failed trying to change dir to " + abs + "\nBad Pathname: Not a directory")
	}

	err = filepath.Walk(abs, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasPrefix(path[strings.LastIndex(path, string(os.PathSeparator))+1:], "._") {
			return nil
		}

		return os.Remove(path)
	})
	if err != nil {
		exit(err.Error())
	}
}
