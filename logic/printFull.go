package logic

import (
	"fmt"
	"strconv"

	"git/ssengerb/my-ls-1/models"
)

/*PrintFull prints a list of files in long format with details like:

Permissions (-rw-r--r--)
Number of links
Owner and group
File size
Last modified date and time
File name (with colors for directories and symlinks)*/

func PrintFull(flag models.FlagOptions, files []models.File, num int) {
	var total int64

	// Calculate total blocks
	for i := 0; i < len(files); i++ {
		if flag.Flag_a || !files[i].IsHidden {
			total += files[i].Total
		}
	}

	if num == 1 {
		fmt.Println("total", total)
	}

	// Print each file
	for i := 0; i < len(files); i++ {
		if !flag.Flag_a && files[i].IsHidden {
			continue
		}

		// Permissions
		fmt.Print(files[i].Permissions + " ")
		if len(files[i].Permissions) <= 10 {
			fmt.Print(" ")
		}

		// Links
		linksStr := strconv.FormatUint(files[i].Links, 10)
		if len(linksStr) == 1 {
			fmt.Print(" ")
		}
		fmt.Print(linksStr + " ")

		// Owner
		fmt.Print(files[i].Owner)
		for j := len(files[i].Owner); j < Format.MaxOwnerName; j++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")

		// Group
		fmt.Print(files[i].Group)
		for j := len(files[i].Group); j < Format.MaxGroupName; j++ {
			fmt.Print(" ")
		}
		fmt.Print(" ")

		// Size
		sizeStr := strconv.FormatInt(files[i].Size, 10)
		for j := len(sizeStr); j < Format.MaxSize; j++ {
			fmt.Print(" ")
		}
		fmt.Print(sizeStr + " ")

		// Date (Month + Day + Time)
		month := files[i].Time.Format("Jan") // English month abbreviation
		day := strconv.Itoa(files[i].Time.Day())
		if len(day) == 1 {
			fmt.Print(" ")
		}
		fmt.Print(month + " " + day + " ")
		fmt.Print(files[i].Time.Format("15:04") + " ")

		// Name (with colors for directories and symlinks)
		if files[i].IsDir {
			fmt.Print("\033[34m" + files[i].Name + "\033[0m")
		} else if files[i].IsSymlink {
			fmt.Print("\033[96m" + files[i].Name + "\033[0m" + " -> " + files[i].SymlinkTarget)
		} else {
			fmt.Print(files[i].Name)
		}

		fmt.Println()
	}
}