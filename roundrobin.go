package roundrobin

// code adapted from https://github.com/hlts2/round-robin

import (
	"errors"
	"sync/atomic"
)

// ErrNoItems specified for the algorithm
var ErrNoItems = errors.New("no items")

// RoundRobin iterates over the items in a round robin fashion
type RoundRobin struct {
	Options          Options
	itemsMap         map[string]struct{}
	items            []Item
	next             uint32
	currentItemCount uint32
}

// New returns a new RoundRobin structure
func New(items ...string) (*RoundRobin, error) {
	return NewWithOptions(DefaultOptions, items...)
}

// New returns a new RoundRobin structure with custom options
func NewWithOptions(options Options, items ...string) (*RoundRobin, error) {
	if len(items) == 0 {
		return nil, ErrNoItems
	}

	rb := &RoundRobin{itemsMap: make(map[string]struct{}), Options: options}
	rb.Add(items...)

	return rb, nil
}

// Next returns next item
func (r *RoundRobin) Add(items ...string) {
	for _, itemValue := range items {
		if _, ok := r.itemsMap[itemValue]; ok {
			continue
		}

		item := Item{
			value: itemValue,
		}

		r.items = append(r.items, item)
	}
}

// Next returns next item
func (r *RoundRobin) Next() Item {
	currentAmount := atomic.AddUint32(&r.currentItemCount, 1)
	if currentAmount > uint32(r.Options.RotateAmount) {
		atomic.StoreUint32(&r.currentItemCount, 0)
		n := atomic.AddUint32(&r.next, 1)
		return r.items[(int(n)-1)%len(r.items)]
	}
	return r.items[(int(r.next)-1)%len(r.items)]
}