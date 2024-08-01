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

	abs, err := filepath.Abs(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	// volumeName := filepath.VolumeName(abs)
	// isWindows := volumeName != ""
	// isWindows := runtime.GOOS == "windows"
	// fmt.Println(isWindows)

	stat, err := os.Stat(abs)
	if err != nil {
		log.Fatal(err)
	}

	if !stat.IsDir() {
		log.Fatal(abs + " is not a directory")
	}

	filepath.Walk(abs, func(path string, info fs.FileInfo, err error) error {
		if !strings.HasPrefix(path[strings.LastIndex(path, string(os.PathSeparator))+1:], "._"){
			return nil
		}

		fmt.Println("Would remove "+path)

		// if err := os.Remove(path); err != nil {
		// 	return err
		// } else {
		// 	fmt.Println("Removed "+path)
		// }

		return nil
	})
}
