package histogram

import (
	"sync"

	"golang.org/x/exp/maps"
)

type SyncHistogram[T comparable] struct {
	sync.RWMutex
	store map[T]int
}

func NewSync[T comparable]() *SyncHistogram[T] {
	return &SyncHistogram[T]{
		store: make(map[T]int),
	}
}

func (h *SyncHistogram[T]) Add(item T) {
	h.Lock()
	h.store[item]++
	h.Unlock()
}

func (h *SyncHistogram[T]) Count() int {
	h.RLock()
	total := len(h.store)
	h.RUnlock()
	return total
}

func (h *SyncHistogram[T]) Items() map[T]int {
	h.RLock()
	res := maps.Clone(h.store)
	h.RUnlock()
	return res
}
