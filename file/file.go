package file

import (
	"bufio"
	"fmt"
	"os"
)

// Lines reads `filename` and returns a string slice containing its lines
func Lines(filename string) ([]string, error) {
	var lines []string

	err := LinesFunc(filename, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		return nil, err
	}

	return lines, nil
}

// LinesFunc reads `filename` and calls a function for each line
func LinesFunc(filename string, mapFn func(string)) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("could not read file %q: %w", filename, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		mapFn(line)
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("could not scan file %q: %w", filename, err)
	}

	return scanner.Err()
}
