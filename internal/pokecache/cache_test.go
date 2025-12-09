package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://test.com?limit=20",
			value: []byte("here are some tests"),
		},
	}
	const interval = 5 * time.Second
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.value)
			val, err := cache.Get(c.key)
			if err != nil {
				t.Fatalf("%s", err.Error())
			}
			if string(val) != string(c.value) {
				t.Fatalf("expected to have the same value: %s, %s", string(val), string(c.value))
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = time.Millisecond * 10
	const waitTime = baseTime + time.Millisecond*10
	cache := NewCache(baseTime)
	key, value := "http://test.com?offset=20", []byte("second page response")
	cache.Add(key, value)

	if _, err := cache.Get(key); err != nil {
		t.Fatal(err.Error())
	}
	time.Sleep(waitTime)
	if _, err := cache.Get(key); err == nil {
		t.Fatal("Not expected to find a key")
	}
}
