package roundrobin

type Options struct {
	RotateAmount int32
}

var DefaultOptions = Options{
	RotateAmount: 1,
}
