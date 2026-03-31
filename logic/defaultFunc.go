package logic

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"git/ssengerb/my-ls-1/models"
)

func DefaultFunc(list []models.File, flag models.FlagOptions, path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln("Error reading directory:", err)  //if it fails reads an err
	}

	list = make([]models.File, len(files)+2) //slices for all files 

	cwd, err := os.Getwd()  //gets the current working dir
	if err != nil {
		fmt.Println(err)   //print err and stop the function
		return
	}
	fullPath := filepath.Join(cwd, path) //combine the current dir and given path
	parentDir := filepath.Dir(fullPath) //parent dir path
	FillCurrentDir(path, &list[len(files)+1])
	list[len(files)+1].Name = "."
	FillCurrentDir(parentDir, &list[len(files)])
	list[len(files)].Name = ".."

	for i, file := range files {
		list[i].Name = file.Name()
		FillFile(path, &list[i])
	}

	if flag.Flag_t {
		SortTime(list)
	} else {
		SortDefault(list)
	}

	if flag.Flag_r {
		SortReverse(list)
	}

	if flag.Flag_l {
		PrintFull(flag, list, 1)
	} else {
		PrintDefault(flag, list)
	}
}
