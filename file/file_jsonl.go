package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type FileJSONLEncoder struct {
	f   io.Closer
	enc *json.Encoder
}

func NewFileJSONLEncoder(filename string) (*FileJSONLEncoder, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("could not create file %q: %v", filename, err)
	}

	return &FileJSONLEncoder{f: f, enc: json.NewEncoder(f)}, nil
}

func (e *FileJSONLEncoder) Close() error {
	return e.f.Close()
}

func (e *FileJSONLEncoder) Encode(v any) error {
	return e.enc.Encode(v)
}

type FileJSONLDecoder[T any] struct {
	f   io.Closer
	dec *json.Decoder
}

func NewFileJSONLDecoderFS[T any](f fs.File) (*FileJSONLDecoder[T], error) {
	return &FileJSONLDecoder[T]{f: f, dec: json.NewDecoder(f)}, nil
}

func NewFileJSONLDecoder[T any](filename string) (*FileJSONLDecoder[T], error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %q: %v", filename, err)
	}

	return NewFileJSONLDecoderFS[T](f)
}

func (e *FileJSONLDecoder[T]) Close() error {
	return e.f.Close()
}

func (e *FileJSONLDecoder[T]) Decode() (T, error) {
	var v T
	err := e.dec.Decode(&v)
	return v, err
}
