package files

import (
	"bufio"
	"os"
)

// ReadFile reads a file and returns the contents as a string
func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	var contents string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contents += scanner.Text()
	}

	return contents, nil
}
