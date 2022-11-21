package roundrobin

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundrobin(t *testing.T) {
	roundRobin, err := New("a", "b", "c", "d")
	require.Nil(t, err)
	for i := 0; i < 1000; i++ {
		next := roundRobin.Next()
		require.NotEmpty(t, next.value)
	}
}

func TestRoundrobinWithNoItems(t *testing.T) {
	roundRobin, err := New()
	require.ErrorIs(t, err, ErrNoItems)
	require.Nil(t, roundRobin)
}

func TestRoundrobinWithOneItem(t *testing.T) {
	roundRobin, err := New("a")
	require.Nil(t, err)
	for i := 0; i < 1000; i++ {
		require.Equal(t, roundRobin.Next().String(), "a")
	}
}

func TestRoundrobinWithGrowingItems(t *testing.T) {
	var items []string
	for i := 0; i < 500; i++ {
		items = append(items, fmt.Sprint(i))
		roundRobin, err := New(items...)
		require.Nil(t, err)
		for i := 0; i < 10000; i++ {
			next := roundRobin.Next()
			require.NotEmpty(t, next.value)
		}
	}
}

func TestRoundrobinWithRotate(t *testing.T) {
	for rotateAmount := 1; rotateAmount < 100; rotateAmount++ {
		roundRobin, err := NewWithOptions(Options{RotateAmount: int32(rotateAmount)}, "a", "b", "c")
		require.Nil(t, err)
		c := 0
		expected := "a"
		for i := 0; i < 1000; i++ {
			if c == rotateAmount {
				switch expected {
				case "a":
					expected = "b"
				case "b":
					expected = "c"
				case "c":
					expected = "a"
				}
				c = 0
			}
			c++
			next := roundRobin.Next()
			require.Equal(t, expected, next.String(), "i=%d c=%d", i, c)
		}
	}
}

func TestRoundrobinwithGoroutinesUsingStats(t *testing.T) {
	roundRobin, err := New("a", "b", "c", "d")
	require.Nil(t, err)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(rbx *RoundRobin, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				rbx.Next()
			}
		}(roundRobin, &wg)
	}

	wg.Wait()

	/*
		In Roundrobin algo all items have same priority and
		are assinged in circular order Hence test results for 100
		access with 3 iterations from different goroutines should be
		a=75,b=75,c=75,d=75
	*/

	for _, v := range roundRobin.items {
		require.Equal(t, int32(75), v.Stats.Count, "Test Failed for item %v", v.value)
	}

}
