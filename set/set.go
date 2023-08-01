package set

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
	res := make([]T, len(h.store))
	idx := 0
	for item := range h.store {
		res[idx] = item
		idx++
	}
	return res
}
