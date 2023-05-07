package files

import (
	"fmt"
	"os"
)

const (
	NoRootDir = iota
	NoSubDir
	NoFile
	AllExists
)

const (
	Suffix  = ".toml"
	RootDir = "./Config"
	APIDir  = "./Config/API"
)

func CheckDir(name string) int {
	if _, err := os.Stat(RootDir); os.IsNotExist(err) {
		return NoRootDir
	}
	if _, err := os.Stat(APIDir); os.IsNotExist(err) {
		return NoSubDir
	}
	if _, err := os.Stat(APIDir + "/" + name + Suffix); os.IsNotExist(err) {
		return NoFile
	}
	return AllExists
}

func CreateKey(name string, key string) {
	err := os.MkdirAll(APIDir, 0755)
	if err != nil {
		println(err)
	}
	file, err := os.Create(APIDir + "/" + name + Suffix)
	if err != nil {
		println(err)
	}
	defer file.Close()
	file.WriteString(fmt.Sprintf("KEY = \"%s\"", key))
}
