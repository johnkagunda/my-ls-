package logic

import "git/ssengerb/my-ls-1/models"


/*sorts by time
Newest files come first
Oldest files go last
 Uses Bubble Sort (O(n²))*/


//It sorts files based on their Time field (usually modification time).
func SortTime(files []models.File) {
	for i := 0; i < len(files); i++ {  //if file j is older than file j+1 swap them 
		for j := 0; j < len(files)-i-1; j++ {
			if files[j].Time.Before(files[j+1].Time) {
				files[j], files[j+1] = files[j+1], files[j]
			}
		}
	}
}
