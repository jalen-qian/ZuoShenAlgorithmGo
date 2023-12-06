package class_15

/**
LeetCode第547题：省份数量
https://leetcode.cn/problems/number-of-provinces/
美版：
https://leetcode.com/problems/friend-circles/
题目描述：
有 n 个城市，其中一些彼此相连，另一些没有相连。如果城市 a 与城市 b 直接相连，且城市 b 与城市 c 直接相连，那么城市 a 与城市 c 间接相连。

省份 是一组直接或间接相连的城市，组内不含其他没有相连的城市。

给你一个 n x n 的矩阵 isConnected ，其中 isConnected[i][j] = 1 表示第 i 个城市和第 j 个城市直接相连，而 isConnected[i][j] = 0 表示二者不直接相连。

返回矩阵中 省份 的数量。
例子1：
输入：isConnected =
[
	[1,1,0],
	[1,1,0],
	[0,0,1]
]
输出：2
因为0认识1，1认识0；但是2不认识0和1，所以有两个省份

例子2：
isConnected = [
	[1,0,0],
	[0,1,0],
	[0,0,1]
]
输出：3
因为没有互相认识的城市

解释：如果 isConnected[i][j] = 1, 则 isConnected[j][i] 也必然等于1，不存在A认识B，但是B不认识A的情况。

思路：联通性问题，都可以用并查集来做。
1. 准备一个并查集，先将所有的城市都加入并查集，比如例子1中，有城市 0 1 2 共3个城市
2. 遍历 isConnected 的上半部分，发现1的话，就将对应的索引位置的集合合并
3. 返回并查集中的集合数量就是要求的值
解释：为什么只遍历上半部分？因为isConnected[i][j] = 1, 则 isConnected[j][i] 也必然等于1，只需要遍历上半部分，就足够知道城市之间的认识关系了。

补充：这里的并查集不需要完整的并查集，只需要提供合并集合和查询集合数量就行了，不需要提供检查两个元素是否在同一个集合的函数
由于都是数字，可以使用数组实现并查集

[
	[1,0,0,1],
	[0,1,1,0],
	[0,1,1,1],
	[1,0,1,1],
]
*/

func findCircleNum(isConnected [][]int) int {
	mUnionFind := newUnionFind(len(isConnected))
	// 只遍历矩阵的右上半部分
	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected); j++ {
			// 如果i认识j,则将i和j所在的集合合并
			if isConnected[i][j] == 1 {
				mUnionFind.union(i, j)
			}
		}
	}
	// 返回最终的集合个数
	return mUnionFind.setsNum()
}

// unionFind 使用数组实现的并查集
type unionFind struct {
	parents []int // parent[i] = j 表示 i 的父亲是 j
	size    []int // 集合大小 size[i] = 10 表示以i为代表节点的
	sets    int   // 集合个数
	help    []int // 充当findFather中的栈的作用，用于做路径压缩
}

func newUnionFind(n int) *unionFind {
	parents := make([]int, n)
	size := make([]int, n)
	help := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i // 刚开始时，都是自己是自己的父亲，单独一个集合
		size[i] = 1    // 刚开始时，集合大小都是1
	}
	return &unionFind{
		parents: parents,
		size:    size,
		sets:    n,
		help:    help,
	}
}

// 往上到不能再往上找到代表节点返回，并在这个过程中压缩路径
func (u *unionFind) findFather(i int) int {
	hi := 0
	for u.parents[i] != i {
		u.help[hi] = i
		i = u.parents[i]
		hi++
	}
	for hi--; hi >= 0; hi-- {
		u.parents[u.help[hi]] = i
	}
	return i
}

// 合并a b所在的集合
func (u *unionFind) union(a, b int) {
	aHead := u.findFather(a)
	bHead := u.findFather(b)
	if aHead != bHead {
		// 分出较大集合和较小集合
		maxHead, minHead := aHead, bHead
		if u.size[aHead] < u.size[bHead] {
			maxHead, minHead = bHead, aHead
		}
		// 将较小集合合并到较大集合
		u.parents[minHead] = maxHead
		// 将较小集合的size清零
		u.size[minHead] = 0
		// 总集合数减1
		u.sets--
	}
}

// 返回集合数量
func (u *unionFind) setsNum() int {
	return u.sets
}
