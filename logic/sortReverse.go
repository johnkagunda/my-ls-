package logic

import "git/ssengerb/my-ls-1/models"

/*This function:

Swaps first with last, second with second-last, etc.
Reverses the slice in-place
Is used to implement reverse sorting (-r)*/

func SortReverse(files []models.File) {  //It reverses the order of the files slice.
	n := len(files)
	for i := 0; i < n/2; i++ {
		files[i], files[n-i-1] = files[n-i-1], files[i]
	}
}
