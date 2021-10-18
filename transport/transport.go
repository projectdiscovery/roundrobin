package transport

import (
	"github.com/Mzack9999/roundrobin"
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

func (rt *RoundTransport) Next() string {
	return rt.Next()
}
