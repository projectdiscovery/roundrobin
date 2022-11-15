package transport

import (
	"github.com/projectdiscovery/roundrobin"
)

// RoundTransport iterates proxies using RoundRobin
type RoundTransport struct {
	rb *roundrobin.RoundRobin
}

// New returns a New RoundTransport Structure
func New(proxies ...string) (*RoundTransport, error) {
	rb, err := roundrobin.New(proxies...)
	if err != nil {
		return nil, err
	}
	return &RoundTransport{rb: rb}, nil
}

// NewWithOptions returns a New RoundTransport Struct with custom options
func NewWithOptions(upstreamRequestNumber int, proxies ...string) (*RoundTransport, error) {
	rbOptions := roundrobin.DefaultOptions
	rbOptions.RotateAmount = int32(upstreamRequestNumber)
	rb, err := roundrobin.NewWithOptions(rbOptions, proxies...)
	if err != nil {
		return nil, err
	}
	return &RoundTransport{rb: rb}, nil
}

// Next returns next proxy/item
func (rt *RoundTransport) Next() string {
	next := rt.rb.Next()
	return next.String()
}
