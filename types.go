package roundrobin

import "sync/atomic"

type Item struct {
	value string
	Stats Stats
}

type Stats struct {
	Count int32
}

func (s *Stats) Inc(value int32) {
	atomic.AddInt32(&s.Count, value)
}

func (s *Stats) Reset() {
	atomic.StoreInt32(&s.Count, 0)
}
