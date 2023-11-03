package utils

type HashSet[T comparable] struct {
	set map[T]struct{}
}

// NewHashSet 构造函数
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{make(map[T]struct{})}
}

func (set *HashSet[T]) Add(item T) {
	set.set[item] = struct{}{}
}

func (set *HashSet[T]) Contains(item T) bool {
	_, found := set.set[item]
	return found
}

func (set *HashSet[T]) Remove(item T) {
	delete(set.set, item)
}
