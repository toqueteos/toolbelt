package histogram

import "golang.org/x/exp/maps"

type Histogram[T comparable] struct {
	store map[T]int
}

func New[T comparable]() *Histogram[T] {
	return &Histogram[T]{
		store: make(map[T]int),
	}
}

func (h *Histogram[T]) Add(item T) {
	h.store[item]++
}

func (h *Histogram[T]) Count() int {
	total := len(h.store)
	return total
}

func (h *Histogram[T]) Items() map[T]int {
	return maps.Clone(h.store)
}
