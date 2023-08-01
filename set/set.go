package set

import (
	"golang.org/x/exp/maps"
)

type Set[T comparable] struct {
	store map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		store: make(map[T]struct{}),
	}
}

func (h *Set[T]) Add(item T) {
	h.store[item] = struct{}{}
}

func (h *Set[T]) Count() int {
	total := len(h.store)
	return total
}

func (h *Set[T]) Items() []T {
	return maps.Keys(h.store)
}
