package set

import "sync"

type SyncSet[T comparable] struct {
	sync.RWMutex
	store map[T]struct{}
}

func NewSync[T comparable]() *SyncSet[T] {
	return &SyncSet[T]{
		store: make(map[T]struct{}),
	}
}

func (h *SyncSet[T]) Add(item T) {
	h.Lock()
	h.store[item] = struct{}{}
	h.Unlock()
}

func (h *SyncSet[T]) Count() int {
	h.RLock()
	total := len(h.store)
	h.RUnlock()
	return total
}

func (h *SyncSet[T]) Items() []T {
	h.RLock()
	res := make([]T, len(h.store))
	idx := 0
	for item := range h.store {
		res[idx] = item
		idx++
	}
	h.RUnlock()
	return res
}
