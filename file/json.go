package file

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

// JSONEncode encodes a value as JSON and writes it into filename.
// By default indent (tab) and escaping are enabled.
func JSONEncode(filename string, v any, opts ...OptionJSON) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %q: %w", filename, err)
	}
	defer f.Close()

	enc := jsonEncoderWithOptions(f, opts)

	err = enc.Encode(v)
	if err != nil {
		return fmt.Errorf("could not json encode contents to file %q: %w", filename, err)
	}

	return nil
}

// JSONDecode opens filename and tries to decode its contents as JSON.
func JSONDecode[T any](filename string) (T, error) {
	var zero T

	f, err := os.Open(filename)
	if err != nil {
		return zero, fmt.Errorf("could not read file %q: %w", filename, err)
	}
	defer f.Close()

	return JSONDecodeFS[T](f)
}

// JSONDecodeFS reads a file contents and tries to decode them as JSON.
func JSONDecodeFS[T any](f fs.File) (T, error) {
	var zero T
	var v T
	err := json.NewDecoder(f).Decode(&v)
	if err != nil {
		return zero, fmt.Errorf("could not json decode file contents: %w", err)
	}

	return v, nil
}
