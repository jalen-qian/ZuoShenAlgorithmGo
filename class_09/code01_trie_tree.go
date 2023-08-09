package class_09

// 前缀树

type node1 struct {
	pass  int        // 经过这个节点的字符串个数
	end   int        // 以当前节点结尾的字符串个数
	nexts [26]*node1 // 子节点，如果为null，则说明i方向的路不存在
}

func newNode1() *node1 {
	return &node1{
		nexts: [26]*node1{},
	}
}

// Trie1 前缀树
type Trie1 struct {
	root *node1
}

func NewTrie1() *Trie1 {
	return &Trie1{
		root: newNode1(),
	}
}

// Insert 向前缀树中插入一个单词
func (t *Trie1) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root // node先来到头节点的位置
	// 所有单词必然经过头节点，头节点的pass++
	node.pass++
	// 遍历这个单词的每个字符
	var path int32
	for _, c := range word {
		// 前缀树的nexts[i]表示ascii码中的第i个字符，'a'是第0个字符，这里 c-'a' 就找到了c字符的通路
		path = c - 'a'
		// 判断通向c的路存不存在，如果不存在，就创建这条路径
		if node.nexts[path] == nil {
			node.nexts[path] = newNode1()
		}
		// 通向c的这条路的节点的pass++，表示多了一个字符串经过这条路
		node.nexts[path].pass++
		// node跳到c字符所在的节点，继续遍历
		node = node.nexts[path]
	}
	// 遍历完后，node肯定在word字符串的最后字符通向的节点，则end++
	node.end++
}

// Delete 从前缀树中删除一个单词
func (t *Trie1) Delete(word string) {
	// 如果没加入过这个单词，则什么都不需要做
	if t.Search(word) == 0 {
		return
	}
	// 加入过，沿途pass--，如果某个沿途节点的pass减到0了，
	// 说明没有任何其他字符串通过这条路了，则直接从这个节点开始丢弃，避免内存泄漏
	node := t.root
	// 头节点的pass--
	node.pass--
	var path int32
	for _, c := range word {
		path = c - 'a'
		node.nexts[path].pass--
		// 途中减到0了，丢弃这条路，golang垃圾回收器会自动回收
		if node.nexts[path].pass == 0 {
			node.nexts[path] = nil
			return
		}
		node = node.nexts[path]
	}
	// 遍历到了最后一个节点，这个节点的end--
	node.end--
}

// Search 搜索单词在前缀树中加入过几次
func (t *Trie1) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root
	// 遍历字符
	var path int32
	for _, c := range word {
		path = c - 'a' // 找到路
		// 如果找不到路了，则说明没有添加过这个字符串，直接返回0
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	// 遍历到了最后一个字符，返回这个节点的end数量，就是加入过几次
	return node.end
}

// PrefixNumber 返回前缀树中有多少个单词是以pre开头的
func (t *Trie1) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root
	// 遍历字符
	var path int32
	for _, c := range pre {
		path = c - 'a' // 找到路
		// 如果提前找不到路了，则说明没有字符串有这个前缀，返回0
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	// 遍历到了最后一个字符，返回这个节点的pass数量，就是以当前字符串为前缀的字符串的数量
	return node.pass
}

// =======================
// Trie1 只能加入26个小写的英文字母组成的字符串，如果希望能加入任何字符串，则使用另一种方式实现
type node2 struct {
	pass int // 经过这个节点的字符串个数
	end  int // 以当前节点结尾的字符串个数
	// 子节点使用map存储，key就是字符的值 在golang中是rune类型，也就是 int32 类型
	// 比如"你好"的'你'对应的值是 20320
	nexts map[rune]*node2
}

func newNode2() *node2 {
	return &node2{
		nexts: make(map[rune]*node2),
	}
}

// Trie2 前缀树，能加入任何字符串
type Trie2 struct {
	root *node2
}

func NewTrie2() *Trie2 {
	return &Trie2{
		root: newNode2(),
	}
}

// Insert 向前缀树中插入一个单词
func (t *Trie2) Insert(word string) {
	if word == "" {
		return
	}
	node := t.root // node先来到头节点的位置
	// 所有单词必然经过头节点，所有头节点的pass++
	node.pass++
	// 遍历这个单词的每个字符
	for _, c := range word {
		// c 本身就是路径
		// 判断通向c的路存不存在，如果不存在，就创建这条路
		if _, ok := node.nexts[c]; !ok {
			node.nexts[c] = newNode2()
		}
		// 通向c的这条路的节点的pass++，表示多了一个字符串经过这条路
		node.nexts[c].pass++
		// node跳到这个节点
		node = node.nexts[c]
	}
	// 遍历完后，node肯定在word字符串的最后字符的节点，则end++
	node.end++
}

// Delete 从前缀树中删除一个单词
func (t *Trie2) Delete(word string) {
	// 如果没加入过这个单词，则什么都不需要做
	if t.Search(word) == 0 {
		return
	}
	// 加入过，沿途pass--，如果某个沿途节点的pass减到0了，
	// 说明没有任何其他字符串通过这条路了，则直接从这个节点开始丢弃，避免内存泄漏
	node := t.root
	// 头节点的pass--
	node.pass--
	for _, c := range word {
		node.nexts[c].pass--
		// 途中减到0了，丢弃这条路，golang垃圾回收器会自动回收
		if node.nexts[c].pass == 0 {
			delete(node.nexts, c)
			return
		}
		node = node.nexts[c]
	}
	// 遍历到了最后一个节点，这个节点的end--
	node.end--
}

// Search 搜索单词在前缀树中加入过几次
func (t *Trie2) Search(word string) int {
	if word == "" {
		return 0
	}
	node := t.root
	// 遍历字符
	for _, c := range word {
		// 如果找不到路了，则说明没有添加过这个字符串，直接返回0
		if _, ok := node.nexts[c]; !ok {
			return 0
		}
		node = node.nexts[c]
	}
	// 遍历到了最后一个字符，返回这个节点的end数量，就是加入过几次
	return node.end
}

// PrefixNumber 返回前缀树中有多少个单词是以pre开头的
func (t *Trie2) PrefixNumber(pre string) int {
	if pre == "" {
		return 0
	}
	node := t.root
	// 遍历字符
	for _, c := range pre {
		// 如果提前找不到路了，则说明没有字符串有这个前缀，返回0
		if _, ok := node.nexts[c]; !ok {
			return 0
		}
		node = node.nexts[c]
	}
	// 遍历到了最后一个字符，返回这个节点的pass数量，就是以当前字符串为前缀的字符串的数量
	return node.pass
}
