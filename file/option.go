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

func WithIndent(value string) OptionJSON {
	return func(o *ConfigJSON) {
		o.Indent = value
	}
}

func WithEscape(value bool) OptionJSON {
	return func(o *ConfigJSON) {
		o.Escape = value
	}
}

func WithPretty() OptionJSON {
	return WithIndent("\t")
}

func jsonEncoderWithOptions(w io.Writer, opts []OptionJSON) *json.Encoder {
	opt := &ConfigJSON{}
	for _, o := range opts {
		o(opt)
	}
	if len(opts) == 0 {
		opt.Indent = ""
		opt.Escape = true
	}

	indent := opt.Indent
	escape := opt.Escape

	enc := json.NewEncoder(w)
	enc.SetIndent("", indent)
	enc.SetEscapeHTML(escape)

	return enc
}
