package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
    entries map[string]entry
    lock *sync.RWMutex
} 

type entry struct {
    createdAt time.Time
    val []byte
}

func NewCache(interval time.Duration) *Cache {
    c := &Cache{
        entries: map[string]entry{},
        lock: &sync.RWMutex{},
    }

    go c.reapLoop(time.Second * 30) 

    return c
}

func (c *Cache) Add(k string, v []byte) {
    c.lock.Lock()
    c.entries[k] = entry{ createdAt: time.Now(), val: v }
    c.lock.Unlock()
}

func (c *Cache) Get(k string) ([]byte, bool) {
    c.lock.RLock()
    v, ok := c.entries[k]
    c.lock.RUnlock()
    
    if !ok {
        return nil, false
    }

    return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
    t := time.NewTicker(interval)
    defer t.Stop()

    for {
        select {

        case <- t.C:
            c.lock.Lock()
            for _, e := range c.entries {
                fmt.Println(e)
            }
            c.lock.Unlock()
        }
    }

}
