package roundrobin

// Options for RoundRobin
type Options struct {
	RotateAmount int32
}

var DefaultOptions = Options{
	RotateAmount: 1,
}
