package utils

type HashSet[T comparable] struct {
	set map[T]struct{}
}

// NewHashSet 构造函数
func NewHashSet[T comparable](values ...T) *HashSet[T] {
	set := make(map[T]struct{})
	for _, value := range values {
		set[value] = struct{}{}
	}
	return &HashSet[T]{set: set}
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

func (set *HashSet[T]) GetRandomValue() T {
	var ans T
	for key, _ := range set.set {
		ans = key
		return ans
	}
	return ans
}

func (set *HashSet[T]) Length() int {
	return len(set.set)
}
