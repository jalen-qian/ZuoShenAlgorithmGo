package class_15

import "ZuoShenAlgorithmGo/class_06"

/**
并查集
并查集是一个数据接口，里面维护了多个元素的集合，我们可以通过构造函数将一个元素列表加入到并查集里。
通过并查集，我们很容易实现两个功能：
1. 给定两个元素，检查这两个元素是否在一个集合里
2. 给定两个元素，将这两个元素合并到一个集合
3. 给定一个元素，返回这个元素所在集合的大小

注意：这里是集合，所以在并查集中，元素相同的值只会出现一个
实现如下，这个并查集的优势是：不管是合并集合还是查询某两个元素是否在同一个集合，时间复杂度都是O(1)，这是经过二十年时间才证明出来的，
直接记住结论就好。
*/

// UnionFind 使用map实现的并查集
type UnionFind[V comparable] struct {
	// 存元素到Node的映射
	nodes map[V]*Node[V]
	// 存关联关系，Key的父亲是Value
	parents map[*Node[V]]*Node[V]
	// 如果给定一个代表节点（每个集合都有一个代表节点，代表节点的父亲是自己），记录这个代表节点所在集合的集合大小
	size map[*Node[V]]int
}

func NewUnionFind[V comparable](arr []V) *UnionFind[V] {
	nodes := make(map[V]*Node[V])
	parents := make(map[*Node[V]]*Node[V])
	size := make(map[*Node[V]]int)
	for _, v := range arr {
		node := &Node[V]{Value: v}
		nodes[v] = node
		parents[node] = node // 自己是自己的父亲节点
		size[node] = 1
	}
	return &UnionFind[V]{
		nodes:   nodes,
		parents: parents,
		size:    size,
	}
}

// IsSameSet 返回a,b两个元素是否在同一个集合
func (u *UnionFind[V]) IsSameSet(a, b V) bool {
	// 代表节点是同一个，则是同一个集合
	return u.findFather(u.nodes[a]) == u.findFather(u.nodes[b])
}

// findFather 往上到不能再往上，找到代表节点，找不到就返回空
func (u *UnionFind[V]) findFather(cur *Node[V]) *Node[V] {
	if cur == nil {
		return nil
	}
	stack := class_06.MyStackWithLinkedList[*Node[V]]{}
	// 往上到不能再往上，找到代表节点
	for u.parents[cur] != cur {
		stack.Push(cur)
		cur = u.parents[cur]
	}
	for !stack.IsEmpty() {
		u.parents[stack.Pop()] = cur
	}
	return cur
}

// Union 合并a 和 b代表的两个集合
func (u *UnionFind[V]) Union(a, b V) {
	aHead := u.findFather(u.nodes[a])
	bHead := u.findFather(u.nodes[b])
	if aHead != bHead && aHead != nil && bHead != nil {
		aSetSize := u.size[aHead]
		bSetSize := u.size[bHead]
		// 小集合挂到大集合上
		mainHead, subHead := aHead, bHead
		if aSetSize < bSetSize {
			mainHead, subHead = bHead, aHead
		}
		u.parents[subHead] = mainHead
		u.size[mainHead] += u.size[subHead]
		// 删除subHead的size记录，因为subHead不再是代表节点
		delete(u.size, subHead)
	}
}

type Node[T comparable] struct {
	Value T
}
