package file

import (
	"encoding/json"
	"io"
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

func jsonEncoderWithOptions(w io.Writer, opts []OptionJSON) *json.Encoder {
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

	enc := json.NewEncoder(w)
	enc.SetIndent("", indent)
	enc.SetEscapeHTML(escape)

	return enc
}
