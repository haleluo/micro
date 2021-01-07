// Package mdns provides a multicast dns registry
package mdns

import (
	"context"

	"fmgo.io/microv2/go-micro/v2/registry"
)

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}

// Domain sets the mdnsDomain
func Domain(d string) registry.Option {
	return func(o *registry.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, "mdns.domain", d)
	}
}
