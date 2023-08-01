package file

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type JSONLEncoder struct {
	f   io.Closer
	enc *json.Encoder
}

func NewJSONLEncoder(filename string) (*JSONLEncoder, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("could not create file %q: %v", filename, err)
	}

	return &JSONLEncoder{f: f, enc: json.NewEncoder(f)}, nil
}

func (e *JSONLEncoder) Close() error {
	return e.f.Close()
}

func (e *JSONLEncoder) Encode(v any) error {
	return e.enc.Encode(v)
}

type JSONLDecoder[T any] struct {
	f   io.Closer
	dec *json.Decoder
}

func NewJSONLDecoderFS[T any](f fs.File) (*JSONLDecoder[T], error) {
	return &JSONLDecoder[T]{f: f, dec: json.NewDecoder(f)}, nil
}

func NewJSONLDecoder[T any](filename string) (*JSONLDecoder[T], error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %q: %v", filename, err)
	}

	return NewJSONLDecoderFS[T](f)
}

func (e *JSONLDecoder[T]) Close() error {
	return e.f.Close()
}

func (e *JSONLDecoder[T]) Decode() (T, error) {
	var v T
	err := e.dec.Decode(&v)
	return v, err
}
