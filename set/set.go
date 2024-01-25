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

func NewFrom[T comparable](inputs []T) *Set[T] {
	s := New[T]()
	for _, item := range inputs {
		s.Add(item)
	}
	return s
}

func (h *Set[T]) Add(item T) {
	h.store[item] = struct{}{}
}

func (h *Set[T]) Remove(item T) {
	delete(h.store, item)
}

func (h *Set[T]) Count() int {
	total := len(h.store)
	return total
}

func (h *Set[T]) Has(item T) bool {
	_, ok := h.store[item]
	return ok
}

func (h *Set[T]) Items() []T {
	return maps.Keys(h.store)
}

func (h *Set[T]) Clear() {
	maps.Clear(h.store)
}