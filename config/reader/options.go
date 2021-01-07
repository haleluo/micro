package reader

import (
	"fmgo.io/microv2/go-micro/v2/config/encoder"
	"fmgo.io/microv2/go-micro/v2/config/encoder/hcl"
	"fmgo.io/microv2/go-micro/v2/config/encoder/json"
	"fmgo.io/microv2/go-micro/v2/config/encoder/toml"
	"fmgo.io/microv2/go-micro/v2/config/encoder/xml"
	"fmgo.io/microv2/go-micro/v2/config/encoder/yaml"
)

type Options struct {
	Encoding map[string]encoder.Encoder
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoding: map[string]encoder.Encoder{
			"json": json.NewEncoder(),
			"yaml": yaml.NewEncoder(),
			"toml": toml.NewEncoder(),
			"xml":  xml.NewEncoder(),
			"hcl":  hcl.NewEncoder(),
			"yml":  yaml.NewEncoder(),
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return options
}

func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		if o.Encoding == nil {
			o.Encoding = make(map[string]encoder.Encoder)
		}
		o.Encoding[e.String()] = e
	}
}
