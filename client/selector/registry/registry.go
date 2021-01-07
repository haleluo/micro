// Package registry uses the go-micro registry for selection
package registry

import (
	"fmgo.io/microv2/go-micro/v2/client/selector"
)

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
