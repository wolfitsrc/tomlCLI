package main

import (
	"fmt"
	"tomlCLI/pkg/files"
	"tomlCLI/pkg/flags"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
)

func main() {
	Flags := flags.ParseFlags()

	if Flags.AddAPI.Name != "" && Flags.AddAPI.Key != "" {
		// Gives the user the option to create the directory and file if they don't exist
		// The -c flag is made to prevent accidental creation of directories and files from accidently being in the wrong directory
		if Flags.CreateDir {
			fmt.Println(Green + "Creating directory and file" + Reset)
			files.CreateKey(Flags.AddAPI.Name, Flags.AddAPI.Key)
		} else {
			fmt.Println(Blue + "Checking if directory and file exist" + Reset)
			out := files.CheckDir(Flags.AddAPI.Name)
			switch out {
			case files.NoRootDir:
				println(Red + "Root directory does not exist" + Reset)
			case files.NoSubDir:
				println(Red + "Subdirectory does not exist" + Reset)
			case files.NoFile:
				files.CreateKey(Flags.AddAPI.Name, Flags.AddAPI.Key)
			case files.AllExists:
				println(Blue + "File already exists" + Reset)
			}
		}
	}
}
