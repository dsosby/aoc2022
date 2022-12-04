package util

import (
	"bufio"
	"os"
)

// Reads all lines into memory - hopefully that doesn't bite me
func ReadLines(inputFile string) []string {
	file, fileErr := os.Open(inputFile)
	if fileErr != nil {
		panic(fileErr)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
