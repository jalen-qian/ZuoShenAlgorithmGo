package class_15

import "strconv"

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
//
// 上面的解法，会在一开始时，就初始化 m * n 长度的数组，那么如果 m 和 n 特别大（比如m=10亿，n=10亿），positions 又特别小（比如position只有几个坐标）
// 这种情况下，初始化一个 10亿*10亿的数组，就很不划算了。如何优化呢？
// 由于用一维数组对应矩阵的坐标，导致需要先初始化。如果我们将索引改成字符串，比如(3,5) 这个坐标，变成"3_5"
// 同时使用map来存储 parents["3_5"] = "7_8" 表示(3,5) 这个坐标 是'1'且代表节点是(7,8) 这个坐标

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

// ==== 优化：当 m 和 n 特别大的情况下，如何不需要一开始就初始化一个 m * n 大小的数组
// 这个方法中，不需要m和n参数了
func numIslands22(m int, n int, positions [][]int) []int {
	uf := NewUnionFind22()
	var ans []int
	for _, position := range positions {
		uf.connect(position[0], position[1])
		ans = append(ans, uf.sets)
	}
	return ans
}

type UnionFind22 struct {
	parents map[string]string
	size    map[string]int
	help    []string
	sets    int
}

func NewUnionFind22() *UnionFind22 {
	return &UnionFind22{
		parents: make(map[string]string),
		size:    make(map[string]int),
		help:    nil,
		sets:    0,
	}
}

func (u *UnionFind22) key(row, col int) string {
	return strconv.Itoa(row) + "_" + strconv.Itoa(col)
}

func (u *UnionFind22) connect(r int, c int) {
	key := u.key(r, c)
	// 如果这个key（也就是矩阵中的对应坐标）没有初始化过，则初始化成一个单独的集合
	if _, ok := u.parents[key]; !ok {
		u.parents[key] = key // 自己是自己的父亲
		u.size[key] = 1
		u.sets++
		// 同时尝试和上下左右4个邻居融合
		u.tryUnion(r, c, r-1, c)
		u.tryUnion(r, c, r+1, c)
		u.tryUnion(r, c, r, c-1)
		u.tryUnion(r, c, r, c+1)
	}
}

func (u *UnionFind22) tryUnion(r1 int, c1 int, r2 int, c2 int) {
	// 如果有任意一个坐标没有被初始化过，直接返回
	key1 := u.key(r1, c1)
	key2 := u.key(r2, c2)
	if u.parents[key1] == "" || u.parents[key2] == "" {
		return
	}
	// 走到这里，说明坐标都没越界，且两个坐标初始化过，则合并集合
	head1 := u.findFather(key1)
	head2 := u.findFather(key2)
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

func (u *UnionFind22) findFather(key string) string {
	u.help = u.help[0:0]
	hi := 0
	for key != u.parents[key] {
		u.help = append(u.help, key)
		hi++
		key = u.parents[key]
	}
	for hi--; hi >= 0; hi-- {
		u.parents[u.help[hi]] = key
	}
	return key
}
