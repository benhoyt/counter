// Package counter is a hash table for counting short strings.

package counter

import (
	"bytes"
)

// Counter is a hash table for counting the frequencies of short strings. The
// zero value is ready to use.
type Counter struct {
	items []CounterItem // hash buckets, linearly probed
	size  int           // number of active items in items slice
}

// CounterItem is the type returned by Counter.Items.
type CounterItem struct {
	Key   []byte // key, nil in Counter.items if not yet used
	Count int    // number of occurences
}

const (
	// FNV-1 64-bit constants from hash/fnv.
	offset64 = 14695981039346656037
	prime64  = 1099511628211

	// Initial length for Counter.items.
	initialLen = 1024
)

// Inc increments the count of the given key by n. If the key is not yet
// present, set its count to n (making a copy of the key slice).
func (c *Counter) Inc(key []byte, n int) {
	// Like hash/fnv New64, Write, Sum64 -- but inlined without extra code.
	hash := uint64(offset64)
	for _, c := range key {
		hash *= prime64
		hash ^= uint64(c)
	}

	// Make 64-bit hash in range for items slice.
	index := int(hash & uint64(len(c.items)-1))

	// If current items more than half full, double length and reinsert items.
	if c.size >= len(c.items)/2 {
		newLen := len(c.items) * 2
		if newLen == 0 {
			newLen = initialLen
		}
		newC := Counter{items: make([]CounterItem, newLen)}
		for _, item := range c.items {
			if item.Key != nil {
				newC.Inc(item.Key, item.Count)
			}
		}
		c.items = newC.items
		index = int(hash & uint64(len(c.items)-1))
	}

	// Look up key, using direct match and linear probing if not found.
	for {
		if c.items[index].Key == nil {
			// Found empty slot, add new item (copying key).
			keyCopy := make([]byte, len(key))
			copy(keyCopy, key)
			c.items[index] = CounterItem{keyCopy, n}
			c.size++
			return
		}
		if bytes.Equal(c.items[index].Key, key) {
			// Found matching slot, increment existing count.
			c.items[index].Count += n
			return
		}
		// Slot already holds another key, try next slot (linear probe).
		index++
		if index >= len(c.items) {
			index = 0
		}
	}
}

// Items returns a copy of the incremented items.
func (c *Counter) Items() []CounterItem {
	var items []CounterItem
	for _, item := range c.items {
		if item.Key != nil {
			items = append(items, item)
		}
	}
	return items
}
