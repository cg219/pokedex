package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
    const interval = time.Second * 5

    cache := *NewCache(interval)
    cache.Add("hi", []byte("There"))

    time.Sleep(time.Second * 10)
    _, ok := cache.Get("hi")

    if !ok {
        t.Errorf("expected cache to be set")
    }
}
