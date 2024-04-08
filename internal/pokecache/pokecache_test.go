package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.store == nil {
		t.Error("cache store should not be nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey   string
		inputValue []byte
	}{
		{
			inputKey:   "key1",
			inputValue: []byte("value1"),
		},
		{
			inputKey:   "key2",
			inputValue: []byte("value2"),
		},
		{
			inputKey:   "",
			inputValue: []byte("value3"),
		},
	}

	for _, c := range cases {
		cache.Add(c.inputKey, c.inputValue)
		actual, exist := cache.Get(c.inputKey)
		if !exist {
			t.Errorf("%s not found", c.inputKey)
			continue
		}

		if string(actual) != string(c.inputValue) {
			t.Errorf("%s value does not match %s",
				string(actual),
				string(c.inputValue),
			)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)
	key := "key"

	c.Add(key, []byte("value"))
	time.Sleep(interval + time.Millisecond)

	_, ok := c.Get(key)
	if ok {
		t.Errorf("%q should have been reaped", key)
	}
}

func TestReapFail(t *testing.T) {
	interval := 10 * time.Millisecond
	c := NewCache(interval)
	key := "key"

	c.Add(key, []byte("value"))
	time.Sleep(interval / 2)

	_, ok := c.Get(key)
	if !ok {
		t.Errorf("%q should not have been reaped", key)
	}
}
