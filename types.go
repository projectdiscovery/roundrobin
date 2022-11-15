package roundrobin

import "sync/atomic"

// Item stores value & stats of each element in RoundRobin
type Item struct {
	value string
	Stats Stats
}

// String returns value of item
func (i Item) String() string {
	return i.value
}

// Stats track frequency of item in RoundRobin
type Stats struct {
	Count int32
}

// Inc increments Stats counter
func (s *Stats) Inc(value int32) {
	atomic.AddInt32(&s.Count, value)
}

// Resets Stats counter
func (s *Stats) Reset() {
	atomic.StoreInt32(&s.Count, 0)
}
