// main.go - entry point for the my-ls-1 program
// This file contains the program's main function and shows how flags and
// input are processed before delegating work to the logic package.

package main

import (
"fmt"
"os"

"git/ssengerb/my-ls-1/logic"
"git/ssengerb/my-ls-1/models"
)

// main - program entry point
// The program expects to be invoked as: go run . my-ls [FLAGS] [FILES/DIRS]
// It parses the initial arguments to confirm the subcommand and then uses
// logic.CheckFlagsAndInput to get parsed flags and remaining inputs.
func main() {
// Basic usage check: the first argument after the executable should be
// the literal string "my-ls". If not present, we print usage and continue
// (the original program does not exit here; it just prints usage).
if len(os.Args) < 2 || os.Args[1] != "my-ls" {
fmt.Print("Usage: go run . my-ls [FLAGS]\n\nEX: go run . my-ls -l\n")
}

// Parse flags and input paths. We pass os.Args[2:] because os.Args[0]
// is the binary path and os.Args[1] is expected to be the "my-ls" token.
flag, input := logic.CheckFlagsAndInput(os.Args[2:])

// list will be used as a slice buffer that gets populated by functions
// in the logic package. It is passed around by value but functions that
// fill it receive pointers to its elements.
var list []models.File

// If the user provided no file/directory names, operate on the current
// directory (".") using either ReverseFunc (recursive -R) or DefaultFunc.
if len(input) == 0 {
if flag.Flag_R {
// Recursive listing starting from current directory
logic.ReverseFunc(list, flag, ".")
} else {
// Non-recursive default listing
logic.DefaultFunc(list, flag, ".")
}
} else {
// If inputs are provided, handle them. The behavior differs when
// the recursive (-R) flag is provided.
if flag.Flag_R {
// Separate provided paths into directories (ds) and files (fs)
ds, _ := logic.DevideDirectoryFile(input)
// There is a special-case check used by the original project.
// If the only argument is "-" and no flags are set, it returns.
if len(ds) == 1 && ds[0] == "-" && !flag.Flag_R && !flag.Flag_a && !flag.Flag_l && !flag.Flag_r && !flag.Flag_t {
return
}
// Iterate over directories and run recursive listing for each
for i, dir := range ds {
logic.ReverseFunc(list, flag, dir )
// Print a blank line between multiple directory outputs
if i+1 < len(ds) {
fmt.Println 
}
}
} else {
// Non-recursive but with explicit inputs: delegate to DefaultHard,
// which handles mixed files and directories passed on the command line.
logic.DefaultHard(list, flag, input)
}
}
}
