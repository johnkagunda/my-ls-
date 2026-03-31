package logic

import (
	"strings"
	"unicode"

	"git/ssengerb/my-ls-1/models"
)



/*1. Visible files come first
2. Hidden files come last
3. Names are sorted alphabetically (case-insensitive)
 4. Letters/numbers come before symbols
  5. Shorter names come first if they share a prefix*/


func SortDefault(files []models.File) { //sort slices of model.FIle 
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(files)-i-1; j++ {
			if files[j].IsHidden && !files[j+1].IsHidden { //if the current file is hidden and the next is not hidden do nothing 
				continue
			}
			if !files[j].IsHidden && files[j+1].IsHidden {
				files[j], files[j+1] = files[j+1], files[j] //if the current is hidden and the next one is visible swap them
			} else if !lsSort(strings.ToLower(files[j].Name), strings.ToLower(files[j+1].Name)) {//lsort for ordering
				files[j], files[j+1] = files[j+1], files[j] //converts both names to lowercase , casE insensitive sorting
			}//if order is wrong swap them 
		}
	}
} //visible files move befor hidden ones

func lsSort(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	for i := 0; i < len(r1) && i < len(r2); i++ {
		isRune1Letter := unicode.IsLetter(r1[i]) || unicode.IsNumber(r1[i]) //check if  the letter is a-z , 0-9
		isRune2Letter := unicode.IsLetter(r2[i]) || unicode.IsNumber(r2[i])

		if isRune1Letter != isRune2Letter {
			return isRune2Letter
		}  

		if r1[i] != r2[i] {
			return r1[i] < r2[i]
		}
	}
	return len(r1) < len(r2)
}
