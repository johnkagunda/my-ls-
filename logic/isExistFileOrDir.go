package logic

import (
	"fmt"
	"log"
	"os"
)
/*Take a list of file or directory names (input)
Check if each one exists in the current directory
Print an error and exit if any file or directory is missing*/
func IsExistFileOrDir(input []string) {
	dir, err := os.Open(".")
	if err != nil {
		log.Fatalln("Error opening directory:", err)
	}
	defer dir.Close()
 
	entries, err := dir.ReadDir(0)
	if err != nil {
		log.Fatalln("Error reading directory:", err)
	}

	for _, v := range input {
		find := false
		for _, entry := range entries {
			if v == entry.Name() {
				find = true
			}
		}

		if !find {
			fmt.Println("my-ls: cannot access '" + v + "': No such file or directory")
			os.Exit(1)
		}
	}
}
