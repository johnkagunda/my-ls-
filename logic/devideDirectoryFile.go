package logic

import (
	"log"
	"os"
)
//separates a list of paths into directories and files.
/*:

Splits a list of paths into directories and files/symlinks
Useful for commands that need to process directories separately from files
Stops the program if any path is invalid*/
func DevideDirectoryFile(s []string) ([]string, []string) {
	var ds []string
	var fs []string

	for _, path := range s {
		fileInfo, err := os.Lstat(path)
		if err != nil {
			log.Fatalln("Error:", err)
		}

		if fileInfo.Mode().IsRegular() {
			fs = append(fs, path)
		} else if fileInfo.Mode().IsDir() {
			ds = append(ds, path)
		} else if fileInfo.Mode()&os.ModeSymlink != 0 {
			fs = append(fs, path)
		}
	}

	return ds, fs
}
