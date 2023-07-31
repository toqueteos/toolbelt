package toolbelt

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
)

type OptionJSON func(*ConfigJSON)

type ConfigJSON struct {
	Indent string
	Escape bool
}

func WithPretty() OptionJSON {
	return func(o *ConfigJSON) {
		o.Indent = "\t"
	}
}

func WithEscape(value bool) OptionJSON {
	return func(o *ConfigJSON) {
		o.Escape = value
	}
}

// FileJSONEncode encodes a value as JSON and writes it into filename.
// By default indent (tab) and escaping are enabled.
func FileJSONEncode(filename string, v any, opts ...OptionJSON) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %q: %w", filename, err)
	}
	defer f.Close()

	opt := &ConfigJSON{}
	for _, o := range opts {
		o(opt)
	}
	if len(opts) == 0 {
		opt.Indent = "\t"
		opt.Escape = true
	}

	indent := opt.Indent
	escape := opt.Escape

	enc := json.NewEncoder(f)
	enc.SetIndent("", indent)
	enc.SetEscapeHTML(escape)

	err = enc.Encode(v)
	if err != nil {
		return fmt.Errorf("could not json encode contents to file %q: %w", filename, err)
	}

	return nil
}

// FileJSONDecode opens filename and tries to decode its contents as JSON.
func FileJSONDecode[T any](filename string) (T, error) {
	var zero T

	f, err := os.Open(filename)
	if err != nil {
		return zero, fmt.Errorf("could not read file %q: %w", filename, err)
	}
	defer f.Close()

	return FileJSONDecodeFS[T](f)
}

// FileJSONDecodeFS reads a file contents and tries to decode them as JSON.
func FileJSONDecodeFS[T any](f fs.File) (T, error) {
	var zero T
	var v T
	err := json.NewDecoder(f).Decode(&v)
	if err != nil {
		return zero, fmt.Errorf("could not json decode file contents: %w", err)
	}

	return v, nil
}
