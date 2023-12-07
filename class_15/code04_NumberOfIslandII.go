package class_15

// LeetCode 305题：岛屿数量2
// https://leetcode.cn/problems/number-of-islands-ii/
// 与岛屿问题1不同的是，初始时矩阵全是0（没有岛屿），矩阵中的'1'是一个一个传入的，并且要返回每个岛屿传入时，当时的岛屿数量
// 思路：
// 这种情况，就不能用感染的方式了，只能用并查集。
// 1. 使用并查集时，先只初始化parents、size、help数组，但是不初始化任何集合
// 2. 每加入一个'1'字符，判断这个'1'字符所在的坐标是否加入过，如果没加入过，则初始化为一个单独的集合
// 3. 在2的基础上，尝试将这个坐标和上下左右4个方向上的集合合并（合并条件是：1.不越界 2.集合存在）
// 4. 每次加入一个'1',都记录一次当前并查集中的集合个数
// 5. 如何知道某个索引位置是否初始化过集合呢？答案是：合并时不清空size数组对应的值，判断size[i]是否>0
//    如果size[i]>0，说明i索引位置对应的坐标，加入过'1'（初始化过集合）

func numIslands21(m int, n int, positions [][]int) []int {
	uf := NewUnionFind21(m, n)
	var ans []int
	for _, position := range positions {
		uf.connect(position[0], position[1])
		ans = append(ans, uf.sets)
	}
	return ans
}

type UnionFind21 struct {
	parents []int
	size    []int
	help    []int
	sets    int
	row     int
	col     int
}

func NewUnionFind21(m int, n int) *UnionFind21 {
	length := m * n
	return &UnionFind21{
		parents: make([]int, length),
		size:    make([]int, length),
		help:    make([]int, length),
		row:     m,
		col:     n,
	}
}
func (u *UnionFind21) connect(r int, c int) {
	i := u.index(r, c)
	if u.size[i] == 0 {
		u.parents[i] = i
		u.size[i] = 1
		u.sets++
		// 尝试和上下左右的邻居合并
		// 可以合并的条件：1.当前坐标是'1'（已经满足）2.要合并的坐标存在且是'1'
		u.tryUnion(r, c, r-1, c) // 上面邻居
		u.tryUnion(r, c, r+1, c) // 下面邻居
		u.tryUnion(r, c, r, c-1) // 左边邻居
		u.tryUnion(r, c, r, c+1) // 右边邻居
	}
}

func (u *UnionFind21) index(r, c int) int {
	return r*u.col + c
}

func (u *UnionFind21) tryUnion(r1 int, c1 int, r2 int, c2 int) {
	// 如果坐标越界，直接返回
	if r1 < 0 || r1 >= u.row || c1 < 0 || c1 >= u.col || r2 < 0 || r2 >= u.row || c2 < 0 || c2 >= u.col {
		return
	}
	// 如果有任意一个坐标没有被初始化过，直接返回
	i1 := u.index(r1, c1)
	i2 := u.index(r2, c2)
	if u.size[i1] == 0 || u.size[i2] == 0 {
		return
	}
	// 走到这里，说明坐标都没越界，且两个坐标初始化过，则合并集合
	head1 := u.findFather(i1)
	head2 := u.findFather(i2)
	// 当两个坐标的代表节点不是同一个（说明两个集合没合并过），则合并
	if head1 != head2 {
		big, small := head1, head2
		if u.size[big] < u.size[small] {
			big, small = small, big
		}
		u.parents[small] = u.parents[big]
		u.sets--
	}
}

func (u *UnionFind21) findFather(index int) int {
	hi := 0
	for u.parents[index] != index {
		u.help[hi] = index
		hi++
		index = u.parents[index]
	}
	for hi--; hi >= 0; hi-- {
		u.parents[u.help[hi]] = index
	}
	return index
}
