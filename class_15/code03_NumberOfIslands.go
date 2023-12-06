package class_15

import (
	"ZuoShenAlgorithmGo/class_03"
)

// 岛问题
// https://leetcode.cn/problems/number-of-islands/
// 方法1：感染方法
// 假设有个黑盒方法，infect(i,j int)，如果 i,j 位置是'1', 调用之后会将{i,j}位置的数联通的1全部感染成'2'
// 那么遍历整个数组，每遇到一个'1'，就感染一次，总共感染的次数，就是最后联通的岛的个数。
// [ 1 0 0 0 1
//   1 0 0 1 1
//   1 1 0 1 1
//   0 1 0 0 1 ]
// 调用到{0,0}时，发现是1，则调用一次感染，会将 {0,0}联通的1全部改成2，记录岛数量1
// 如下：
// [ 2 0 0 0 1
//   2 0 0 1 1
//   2 2 0 1 1
//   0 2 0 0 1 ]
// 继续遍历，遍历到{0,4}时，又遇到1，继续调用一次感染，会将{0,4}联通的1全部变成2，记录岛的数量是2
// [ 2 0 0 0 2
//   2 0 0 2 2
//   2 2 0 2 2
//   0 2 0 0 2 ]
// 继续遍历，后序的遍历都不会遇到1了，结束，所以最后返回的结果是2
// 复杂度分析：矩阵大小是M*N的话，时间复杂度是O(M*N)，因为1，主函数每个位置只会碰1遍 2，感染的过程，一个位置只有上下左右4个
// 邻居可能访问这个位置，最多4次，所以时间复杂度O(M*N)

// 方法2：使用并查集，准备一个并查集，初始时遍历整个矩阵，每个1都初始化成自己的集合。
// 遍历矩阵，每遇到一个1，则看这个位置的左和上两个位置是否是1，是就合并
// 最后输出并查集的集合个数。

// 方法3：使用数组实现的并查集

// 使用感染的方式
func numIslands3(grid [][]byte) int {
	islandCount := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				infect(grid, i, j)
				islandCount++
			}
		}
	}
	return islandCount
}

// 感染过程
func infect(gird [][]byte, i, j int) {
	// 越界了，或者不是'1'，都不需要再感染了
	if i < 0 || j < 0 || i >= len(gird) || j >= len(gird[0]) || gird[i][j] != '1' {
		return
	}
	// 感染到了i,j位置，为了防止递归停不下来，将这个位置变为ASCII的0
	gird[i][j] = 0
	// 分别感染上下左右位置
	infect(gird, i-1, j)
	infect(gird, i+1, j)
	infect(gird, i, j-1)
	infect(gird, i, j+1)
}

// 实现思路2：使用map实现的并查集
func numIslands2(grid [][]byte) int {
	// 将grid矩阵，转换成 *Dot 类型的矩阵，每个有'1'的地方都填充一个 Dot 结构体的指针，用来区分每个不同的'1'
	dots := make([][]*Dot, len(grid))
	for i, _ := range dots {
		dots[i] = make([]*Dot, len(grid[0]))
	}
	dotList := make([]*Dot, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dots[i][j] = &Dot{}
				dotList = append(dotList, dots[i][j])
			}
		}
	}
	// 初始化一个并查集，并将矩阵中的dot列表传入
	uf := NewUnionFind1[*Dot](dotList)
	// 先合并最上面一排
	for j := 1; j < len(grid[0]); j++ {
		// 遍历的顺序： (0,1) (0,2) (0,3) 如果左边的点是'1',则合并
		if grid[0][j-1] == '1' && grid[0][j] == '1' {
			uf.Union(dots[0][j-1], dots[0][j])
		}
	}
	// 再合并最左边一列
	for i := 1; i < len(grid); i++ {
		if grid[i-1][0] == '1' && grid[i][0] == '1' {
			uf.Union(dots[i-1][0], dots[i][0])
		}
	}
	// 再合并剩下的
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			// 左边
			if grid[i][j] == '1' && grid[i][j-1] == '1' {
				uf.Union(dots[i][j], dots[i][j-1])
			}
			// 上面
			if grid[i][j] == '1' && grid[i-1][j] == '1' {
				uf.Union(dots[i][j], dots[i-1][j])
			}
		}
	}
	return uf.sets
}

type Dot struct {
	v int32
}

// UnionFind1 并查集
type UnionFind1[V comparable] struct {
	nodes   map[V]*Node[V]
	parents map[*Node[V]]*Node[V]
	size    map[*Node[V]]int
	sets    int // 集合个数
}

func NewUnionFind1[V comparable](items []V) *UnionFind1[V] {
	nodes := make(map[V]*Node[V])
	parents := make(map[*Node[V]]*Node[V])
	size := make(map[*Node[V]]int)
	sets := 0
	for _, v := range items {
		if _, ok := nodes[v]; ok {
			continue
		}
		item := &Node[V]{Value: v}
		nodes[v] = item
		parents[item] = item
		size[item] = 1
		sets++
	}
	return &UnionFind1[V]{
		nodes:   nodes,
		parents: parents,
		size:    size,
		sets:    sets,
	}
}

// 往上到不能再往上，找到代表节点，并返回，在这个过程中进行路径压缩
func (u *UnionFind1[V]) findFather(cur *Node[V]) *Node[V] {
	stack := class_03.NewMyStack[*Node[V]]()
	for u.parents[cur] != cur {
		cur = u.parents[cur]
		stack.Push(cur)
	}
	for !stack.IsEmpty() {
		u.parents[stack.Pop()] = cur
	}
	return cur
}

func (u *UnionFind1[V]) Union(a, b V) {
	aHead, bHead := u.findFather(u.nodes[a]), u.findFather(u.nodes[b])
	if aHead != bHead {
		big, small := aHead, bHead
		if u.size[aHead] < u.size[bHead] {
			big, small = bHead, aHead
		}
		u.parents[small] = big
		u.size[big] = u.size[big] + u.size[small]
		u.size[small] = 0
		u.sets--
	}
}

// 实现方法3：使用数组实现的并查集来实现
// 由于并查集需要元素的不可重复性（集合的概念，加入的元素不能重复），上面的方法2，是将字符的二维数组，转换成了
// Dot对象指针的二维数组，通过对象的指针不同，来区分不同的'1'
// 这里使用数组来实现的话，也需要区分不同的'1'，我们可以将不同的'1'的索引坐标，作为唯一值来加入到并查集中。
// grid[i][j] == '1' , 其一一对应的索引 可以标识成：i*列数 + j
func numIslands1(grid [][]byte) int {
	uf := NewUnionFind2(grid)
	row, col := len(grid), len(grid[0])
	// 先合并第一行
	for c := 1; c < col; c++ {
		// 如果当前位置是1，且当前位置左边也是1，则合并这两个索引位置
		if grid[0][c] == '1' && grid[0][c-1] == '1' {
			uf.Union(0, c, 0, c-1)
		}
	}
	// 再合并第一列
	for r := 1; r < row; r++ {
		// 如果当前位置是1，且当前位置上方也是1，则合并这两个索引位置
		if grid[r][0] == '1' && grid[r-1][0] == '1' {
			uf.Union(r, 0, r-1, 0)
		}
	}
	// 再合并除了0行 0列以外的其他位置，这些位置一定有左侧也一定有上方，所以不用判断边界
	for r := 1; r < row; r++ {
		for j := 1; j < col; j++ {
			// 如果当前坐标是'1', 当前坐标上方也是'1' ，则合并
			if grid[r][j] == '1' && grid[r-1][j] == '1' {
				uf.Union(r, j, r-1, j)
			}
			// 如果当前坐标是'1', 当前坐标左侧也是'1' ，则合并
			if grid[r][j] == '1' && grid[r][j-1] == '1' {
				uf.Union(r, j, r, j-1)
			}
		}
	}
	return uf.sets
}

type UnionFind2 struct {
	// 记录节点的父节点
	// parents[i] = j 表示i的父亲节点是j
	parents []int
	// 记录每个代表节点的大小
	size []int
	// help 用来做路径压缩
	help []int
	// sets 记录一共有多少个集合
	sets int
	// col 记录矩阵一共有多少列，用来根据坐标计算索引
	// 比如[3][2]（3行2列），一共5列的话，索引就是 3*5 + 2 = 17
	col int
}

func NewUnionFind2(grid [][]byte) *UnionFind2 {
	row := len(grid)
	col := len(grid[0])
	length := row * col
	sets := 0
	parents := make([]int, length)
	size := make([]int, length)
	help := make([]int, length)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == '1' {
				i := index(r, c, col)
				parents[i] = i // 初始时，自己是自己的父亲（各自是代表节点）
				size[i] = 1
				sets++
			}
		}
	}
	return &UnionFind2{
		parents: parents,
		size:    size,
		help:    help,
		sets:    sets,
		col:     col,
	}
}

func (u *UnionFind2) findFather(i int) int {
	hi := 0
	for u.parents[i] != i {
		u.help[hi] = i
		hi++
		i = u.parents[i]
	}
	// 做路径压缩
	for hi--; hi >= 0; hi-- {
		u.parents[u.help[hi]] = i
	}
	return i
}

// Union 传入两个坐标，将这两个坐标对应位置的集合合并
// r表示行 c表示列
func (u *UnionFind2) Union(r1, c1 int, r2, c2 int) {
	// 找到索引位置
	index1 := index(r1, c1, u.col)
	index2 := index(r2, c2, u.col)
	// 找到代表节点
	h1 := u.findFather(index1)
	h2 := u.findFather(index2)
	// 代表节点不是同一个，说明不是同一个集合，则合并
	if h1 != h2 {
		big, small := h1, h2
		if u.size[h1] < u.size[h2] {
			big, small = small, big
		}
		// 将较小的集合合并到较大的集合
		u.parents[small] = big
		u.size[big] = u.size[small] + u.size[big]
		u.size[small] = 0
		u.sets--
	}
}

// 根据矩阵坐标计算在并查集中的索引位置
// 比如3行2列，每行5列，则索引位置是 3 * 5 + 2 = 17
func index(r int, c int, col int) int {
	return r*col + c
}
