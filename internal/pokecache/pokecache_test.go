package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache()

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("value"),
		},
		{
			inputKey: "key2",
			inputVal: []byte(""),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		data, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(data) != string(cas.inputVal) {
			t.Errorf("%s doesn't match %s",
				string(data),
				string(cas.inputVal),
			)
		}

		if len(data) != len(cas.inputVal) {
			t.Errorf("%s is not the same size as %s",
				data,
				cas.inputVal,
			)
		}

	}
}
