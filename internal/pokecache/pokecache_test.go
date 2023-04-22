package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Minute * 15)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGet(t *testing.T) {
	cache := NewCache(time.Minute * 15)

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

func TestReapCache(t *testing.T) {
	cache := NewCache(time.Second * 3)
	cache.Add("key1", []byte("val"))

	if _, ok := cache.cache["key1"]; !ok {
		t.Error("added entry does not exist")
	}

	time.Sleep(5 * time.Second)

	if _, ok := cache.cache["key1"]; ok {
		t.Error("added entry did not get reaped")
	}

}

func TestReapCacheFail(t *testing.T) {
	cache := NewCache(time.Second * 3)
	cache.Add("key1", []byte("val"))

	if _, ok := cache.cache["key1"]; !ok {
		t.Error("added entry does not exist")
	}

	time.Sleep(2 * time.Second)

	if _, ok := cache.cache["key1"]; !ok {
		t.Error("added entry should not have been reaped")
	}

}
