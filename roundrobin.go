package roundrobin

// code adapted from https://github.com/hlts2/round-robin

import (
	"errors"
	"sync"
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
	mutex            sync.Mutex
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
	rb.next = 1

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
	defer r.mutex.Unlock()
	r.mutex.Lock()

	currentAmount := atomic.LoadUint32(&r.currentItemCount)
	if currentAmount >= uint32(r.Options.RotateAmount) {
		atomic.StoreUint32(&r.currentItemCount, 1)
		atomic.AddUint32(&r.next, 1)
		return r.getNextItem()
	}
	atomic.AddUint32(&r.currentItemCount, 1)
	return r.getNextItem()
}

func (r *RoundRobin) getNextItem() Item {
	nextItemIndex := (int(r.next) - 1) % len(r.items)
	if nextItemIndex < 0 || nextItemIndex > len(r.items) {
		r.items[0].Stats.Inc(1) // Increment stats by 1 everytime item is retrieved
		return r.items[0]
	}
	r.items[nextItemIndex].Stats.Inc(1)
	return r.items[nextItemIndex]
}
