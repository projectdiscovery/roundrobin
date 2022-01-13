package roundrobin

import (
	"fmt"
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

// func TestActionWaitVisible(t *testing.T) {
// 	response := `
// 		<html>
// 			<head>
// 				<title>Nuclei Test Page</title>
// 			</head>
// 			<button style="display:none" id="test">Wait for me!</button>
// 			<script>
// 				setTimeout(() => document.querySelector('#test').style.display = '', 1000);
// 			</script>
// 		</html>`

// 	actions := []*Action{
// 		{ActionType: ActionTypeHolder{ActionType: ActionNavigate}, Data: map[string]string{"url": "{{BaseURL}}"}},
// 		{ActionType: ActionTypeHolder{ActionType: ActionWaitVisible}, Data: map[string]string{"by": "x", "xpath": "//button[@id='test']"}},
// 	}

// 	t.Run("wait for an element being visible", func(t *testing.T) {
// 		testHeadlessSimpleResponse(t, response, actions, 2*time.Second, func(page *Page, err error, out map[string]string) {
// 			require.Nil(t, err, "could not run page actions")

// 			page.Page().MustElement("button").MustVisible()
// 		})
// 	})

// 	t.Run("timeout because of element not visible", func(t *testing.T) {
// 		testHeadlessSimpleResponse(t, response, actions, time.Second/2, func(page *Page, err error, out map[string]string) {
// 			require.Error(t, err)
// 			require.Contains(t, err.Error(), "Element did not appear in the given amount of time")
// 		})
// 	})
// }
