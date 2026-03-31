package logic // Defines the package name

import (
	"fmt" // Used for printing messages to the console
	"os"  // Used for exiting the program

	"git/ssengerb/my-ls-1/models" // Imports custom models (contains FlagOptions struct)
)

// CheckFlagsAndInput parses command-line arguments
// It separates flags (like -l, -a) from normal inputs (like file names)
func CheckFlagsAndInput(s []string) (models.FlagOptions, []string) {

	input := []string{} // Slice to store non-flag arguments (e.g., file names)

	var flag models.FlagOptions // Struct to store which flags are enabled (all default to false)

	// Loop through all arguments
	for i := 0; i < len(s); i++ {

		// Check if argument starts with '-' and is longer than 1 character
		// This means it's a flag (e.g., "-l", "-la")
		if s[i][0] == '-' && len(s[i]) > 1 {

			// Loop through each character after '-'
			// This allows combined flags like "-lar"
			for j := 1; j < len(s[i]); j++ {

				// Check which flag character is present
				switch s[i][j] {

				case 'l':
					flag.Flag_l = true // Enable -l flag

				case 'R':
					flag.Flag_R = true // Enable -R flag

				case 'a':
					flag.Flag_a = true // Enable -a flag

				case 't':
					flag.Flag_t = true // Enable -t flag

				case 'r':
					flag.Flag_r = true // Enable -r flag

				default:
					// If an unknown flag is found:
					// Print allowed flags
					fmt.Print("-l\n-R\n-a\n-r\n-t\nUse only this flags!\n")

					// Exit the program with error code 1
					os.Exit(1)
				}
			}

		} else {
			// If it's not a flag, treat it as input (e.g., file or directory name)
			input = append(input, s[i])
		}
	}

	// Return:
	// 1. The flags struct showing which flags were set
	// 2. The slice of input arguments
	return flag, input
}