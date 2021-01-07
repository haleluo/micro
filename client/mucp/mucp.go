// Package mucp provides an mucp client
package mucp

import (
	"fmgo.io/microv2/go-micro/v2/client"
)

// NewClient returns a new micro client interface
func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
