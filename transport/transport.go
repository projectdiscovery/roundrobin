package transport

import (
	"github.com/projectdiscovery/roundrobin"
)

type RoundTransport struct {
	rb *roundrobin.RoundRobin
}

func New(proxies ...string) (*RoundTransport, error) {
	rb, err := roundrobin.New(proxies...)
	if err != nil {
		return nil, err
	}
	return &RoundTransport{rb: rb}, nil
}

func NewWithOptions(upstreamRequestNumber int, proxies ...string) (*RoundTransport, error) {
	rbOptions := roundrobin.DefaultOptions
	rbOptions.RotateAmount = int32(upstreamRequestNumber)
	rb, err := roundrobin.NewWithOptions(rbOptions, proxies...)
	if err != nil {
		return nil, err
	}
	return &RoundTransport{rb: rb}, nil
}

func (rt *RoundTransport) Next() string {
	next := rt.rb.Next()
	return next.String()
}
