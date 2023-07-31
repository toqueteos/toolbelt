package toolbelt

import (
	"bufio"
	"fmt"
	"os"
)

// FileEachLines reads `filename` and calls a function for each line
func FileEachLines(filename string, mapFn func(string)) error {
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

// FileLines reads `filename` and returns a string slice containing its lines
func FileLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file %q: %w", filename, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("could not scan file %q: %w", filename, err)
	}

	return lines, nil
}
