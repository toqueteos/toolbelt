package set

import (
	"iter"
	"maps"
	"slices"
	"sync"
)

type SyncSet[T comparable] struct {
	sync.RWMutex
	store map[T]struct{}
}

func NewSync[T comparable]() *SyncSet[T] {
	return &SyncSet[T]{
		store: make(map[T]struct{}),
	}
}

func NewSyncFrom[T comparable](inputs []T) *SyncSet[T] {
	s := NewSync[T]()
	for _, item := range inputs {
		s.store[item] = struct{}{}
	}
	return s
}

func (h *SyncSet[T]) Add(item T) {
	h.Lock()
	h.store[item] = struct{}{}
	h.Unlock()
}

func (h *SyncSet[T]) Remove(item T) {
	h.Lock()
	delete(h.store, item)
	h.Unlock()
}

func (h *SyncSet[T]) Count() int {
	h.RLock()
	total := len(h.store)
	h.RUnlock()
	return total
}

func (h *SyncSet[T]) Has(item T) bool {
	h.RLock()
	_, ok := h.store[item]
	h.RUnlock()
	return ok
}

func (h *SyncSet[T]) Items() iter.Seq[T] {
	h.RLock()
	res := slices.Collect(maps.Keys(h.store))
	h.RUnlock()
	return slices.Values(res)
}

func (h *SyncSet[T]) Clear() {
	h.Lock()
	clear(h.store)
	h.Unlock()
}
