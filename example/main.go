package main

import (
	"fmt"

	"github.com/projectdiscovery/roundrobin"
)

func main() {

	ips := []string{
		"192.168.29.58",
		"192.168.29.1",
		"192.168.29.92",
	}

	// create new roundrobin iterator
	rb, err := roundrobin.New(ips...)
	if err != nil {
		panic(err)
	}

	// rb.Add adds given item to sequence
	rb.Add("192.168.29.86")

	for i := 0; i < 10; i++ {
		// rb.Next returns next item in roundrobin fashion
		val := rb.Next().String()
		fmt.Println(val)
	}

	/*
		Output:
		192.168.29.58
		192.168.29.1
		192.168.29.92
		192.168.29.86
		192.168.29.58
		192.168.29.1
		192.168.29.92
		192.168.29.86
		192.168.29.58
		192.168.29.1
	*/

}
