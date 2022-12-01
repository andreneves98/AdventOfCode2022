package utils

import (
	"bufio"
	"log"
	"os"

	"github.com/pkg/errors"
)

// ReadFile attempts to open a file at the given path
// for reading its contents line by line. If successfull,
// the returned value is a string slice with each line and
// a nil error. Otherwise, returns an empty slice and the error.
func ReadFile(path string) ([]string, error) {
	// Try to open the file for reading
	readFile, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file at '%s'", path)
		return []string{}, errors.Wrap(err, "failed to open file")
	}

	// Scan the file and split it into lines
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Loop over the lines and append them to a slice
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	// Close the file resources
	readFile.Close()

	return lines, nil
}
