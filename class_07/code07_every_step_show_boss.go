package class_07

import (
	"sort"
)

// 手动改写堆练习题

// customer 顾客
type customer struct {
	id        int // id是顾客唯一标识
	buy       int // 当前买的商品数
	enterTime int // 进入得奖区或者候选区的时间
}

// 候选区的排序比较器，排前面的，优先有进入爹区的权限
func candsComparator(c1 *customer, c2 *customer) bool {
	// 排序规则1，如果c1的购买数大于c2的购买数，则c1排前面
	if c1.buy > c2.buy {
		return true
	} else if c1.buy == c2.buy {
		// 排序规则2，如果c1和c2的购买数相等，则最先进入候选区的用户排前面（机会优先给最早进入候选区的用户）
		if c1.enterTime <= c2.enterTime {
			return true
		} else {
			return false
		}
	} else {
		// c1的购买数小于c2，则c1排后面
		return false
	}
}

// 爹区比较器，排前面的，是最可能被候选区淘汰的
func daddyComparator(c1 *customer, c2 *customer) bool {
	// 排序规则1，如果c1的购买数小于c2的购买数，则c1排前面，c1更可能淘汰
	if c1.buy < c2.buy {
		return true
	} else if c1.buy == c2.buy {
		// 排序规则2，如果c1和c2的购买数相等，则最先进入候选区的用户排前面（优先淘汰最早进入爹区的用户）
		if c1.enterTime <= c2.enterTime {
			return true
		} else {
			return false
		}
	} else {
		// c1的购买数大于c2，则c1排后面，c2更可能被淘汰
		return false
	}
}

// TopK 使用手写堆来解决
func TopK(arr []int, op []bool, k int) [][]int {
	// 生成 NewWhoIsYourDaddy 处理器
	var ans [][]int
	d := NewWhoIsYourDaddy(k)
	// 处理每个事件
	for i, id := range arr {
		d.operator(id, op[i], i)
		// 每个事件处理完，都获取一次当前的爹数组
		ans = append(ans, d.getDaddies())
	}
	return ans
}

// WhoIsYourDaddy 使用一个结构体来处理，使用加强堆实现
type WhoIsYourDaddy struct {
	indexMap  map[int]*customer
	daddyHeap *MyHeapGreater[*customer] // 爹区加强堆
	candsHeap *MyHeapGreater[*customer] // 候选区加强堆
	limit     int
}

func NewWhoIsYourDaddy(limit int) *WhoIsYourDaddy {
	return &WhoIsYourDaddy{
		indexMap:  make(map[int]*customer),
		daddyHeap: NewMyHeapGreater[*customer](daddyComparator),
		candsHeap: NewMyHeapGreater[*customer](candsComparator),
		limit:     limit,
	}
}

// operator 处理一个事件
// id 用户id
// isBuy 当前是在买还是退货
// currentTime 当前时间点
func (d *WhoIsYourDaddy) operator(id int, isBuy bool, currentTime int) {
	// 如果没加入过爹区和候选区，且当前是退货，则跳过
	if _, ok := d.indexMap[id]; !ok && !isBuy {
		return
	}
	// 走到这里，要么是购买，要么加入过
	// 没加入过，则加入
	if _, ok := d.indexMap[id]; !ok {
		d.indexMap[id] = &customer{id: id, buy: 0, enterTime: 0}
	}
	// 拿出当前顾客，如果是买，就购买数+1否则购买数-1
	c := d.indexMap[id]
	if isBuy {
		c.buy++
	} else {
		c.buy--
	}
	// 如果购买数减到了0，则删除
	if c.buy == 0 {
		delete(d.indexMap, c.id)
	}
	// 如果爹区和候选区都没加入过，则加入其中一个区，如果爹区没满，就放爹区
	if !d.daddyHeap.Contains(c) && !d.candsHeap.Contains(c) {
		c.enterTime = currentTime
		if d.daddyHeap.Size() == d.limit {
			d.candsHeap.Push(c)
		} else {
			d.daddyHeap.Push(c)
		}
	} else if d.daddyHeap.Contains(c) {
		// 如果爹区有c，则0就删除，不是0就从c的位置调整堆
		if c.buy == 0 {
			d.daddyHeap.Remove(c)
		} else {
			d.daddyHeap.Resign(c)
		}
	} else {
		// 如果在候选区有c，也是0就删除，不是0就从c位置重新调整堆
		if c.buy == 0 {
			d.candsHeap.Remove(c)
		} else {
			d.candsHeap.Resign(c)
		}
	}
	// 看是否候选区的堆顶的顾客能淘汰爹区的堆顶
	d.move(currentTime)
}

func (d *WhoIsYourDaddy) getDaddies() []int {
	return getCurAns(d.daddyHeap.GetAllElements())
}

// move 看是否候选区的堆顶的顾客能淘汰爹区的堆顶
func (d *WhoIsYourDaddy) move(currentTime int) {
	// 候选区堆是空的，则不用淘汰
	if d.candsHeap.IsEmpty() {
		return
	}
	// 看爹区是否是满的，如果不是满的，则直接将候选区的堆顶给爹区
	if d.daddyHeap.Size() < d.limit {
		d.candsHeap.Peek().enterTime = currentTime
		d.daddyHeap.Push(d.candsHeap.Pop())
	} else {
		// 爹区是满的，看候选区的堆顶能否淘汰爹区的堆顶
		// 爹区的堆顶
		oldDaddy := d.daddyHeap.Peek()
		// 候选区的堆顶
		newDaddy := d.candsHeap.Peek()
		// 新爹的购买数 > 旧爹，则新爹把旧爹淘汰
		if newDaddy.buy > oldDaddy.buy {
			// 弹出堆顶
			d.daddyHeap.Pop()
			d.candsHeap.Pop()
			// 压入新爹
			newDaddy.enterTime = currentTime
			d.daddyHeap.Push(newDaddy)
			// 旧爹压入候选区
			oldDaddy.enterTime = currentTime
			d.candsHeap.Push(oldDaddy)
		}
	}
}

// TopKCompare 使用暴力方法，同时作为对数器的比较函数
func TopKCompare(arr []int, op []bool, k int) [][]int {
	var ans [][]int
	// 记录顾客是否现在已经在得奖区或者候选区之一
	m := make(map[int]*customer)
	// 使用切片存储候选区和得奖区的顾客
	cands := make([]*customer, 0) // 候选区
	daddy := make([]*customer, 0) // 得奖区
	// arr数组的下标可以作为时间点，依次遍历每个事件
	for i, id := range arr {
		isBuy := op[i] // 当前事件是买了商品还是退货了商品
		// 如果当前是退货，且之前没加入过到任何区域（可能是退货被删除了，也可能第一次出现就是退货），按照规则，这个事件忽略
		if _, ok := m[id]; !isBuy && !ok {
			ans = append(ans, getCurAns(daddy))
			continue
		}
		// 走到这里，要么是购买，要么加入过
		// 没加入过，则加入
		if _, ok := m[id]; !ok {
			m[id] = &customer{id: id, buy: 0, enterTime: 0} // 初始化，enterTime后面会调整，这里先统一写成0
		}
		// 取出这个顾客
		c := m[id]
		// 如果是购买，则购买数+1，退货则购买数-1
		if isBuy {
			c.buy++
		} else {
			c.buy--
		}
		// 如果购买数减到了0，则从map中删除
		if c.buy == 0 {
			delete(m, id)
		}
		// 如果既不在候选区，也不在爹区，则判断爹区是否满了，没满就放爹区，否则放候选区
		if !contains(daddy, c) && !contains(cands, c) {
			c.enterTime = i // 记录放入的时间
			// 满了，放候选区
			if len(daddy) == k {
				cands = append(cands, c)
			} else {
				// 爹区没满，直接放爹区
				daddy = append(daddy, c)
			}
		}
		// 走到这里，则当前顾客必在一个区，要么两个区都没有，刚刚放到了爹区和候选区中的一个，要么之前就在某个区
		// 不管哪种情况，c顾客的购买数变了，可能需要调整候选区的顾客到爹区，或者淘汰爹区中的顾客到候选区
		// 先删除所有为0的
		deleteZeroBuy(&daddy)
		deleteZeroBuy(&cands)
		// 重新按照题目要求排序
		sort.SliceStable(cands, func(i, j int) bool {
			return candsComparator(cands[i], cands[j])
		})
		sort.SliceStable(daddy, func(i, j int) bool {
			return daddyComparator(daddy[i], daddy[j])
		})
		// 重新调整，爹区需要淘汰的，就淘汰，候选区能进爹区的就进
		move(&daddy, &cands, k, i)
		// 获取这次事件后的topK
		ans = append(ans, getCurAns(daddy))
	}
	return ans
}

func move(daddy *[]*customer, cands *[]*customer, k, currentTime int) {
	// 如果候选区是空的，则不需要调整
	if len(*cands) == 0 {
		return
	}
	// 候选区不是空的，看爹区是否满了，如果爹区没满（说明刚刚删除了一个爹区的），则直接将候选区的0号给爹区
	if len(*daddy) < k {
		c := (*cands)[0]
		c.enterTime = currentTime // 每次进入新区，都要重新记录时间
		// 加入爹区
		*daddy = append(*daddy, c)
		// 候选区删除0位置
		*cands = (*cands)[1:]
	} else {
		// 爹区是满的，且候选区也有，则需要用候选区的0号淘汰爹区的0号
		oldDaddy := (*daddy)[0]
		newDaddy := (*cands)[0]
		// 如果候选区的0号购买数>爹区的0号，则两者交换
		if newDaddy.buy > oldDaddy.buy {
			oldDaddy.enterTime = currentTime
			newDaddy.enterTime = currentTime
			(*daddy)[0] = newDaddy
			(*cands)[0] = oldDaddy
		}
	}
}

func deleteZeroBuy(arr *[]*customer) {
	var noZero []*customer
	for _, c := range *arr {
		if c.buy != 0 {
			noZero = append(noZero, c)
		}
	}
	// 先清空
	*arr = (*arr)[:0]
	// 再把非0的填充回来
	for _, c := range noZero {
		*arr = append(*arr, c)
	}
}

// 判断顾客是否在某个区域
func contains(arr []*customer, c *customer) bool {
	for _, currentCustomer := range arr {
		if currentCustomer.id == c.id {
			return true
		}
	}
	return false
}

// 从daddy中获取当前的top K 用户的id
func getCurAns(daddy []*customer) []int {
	var ans []int
	for _, c := range daddy {
		ans = append(ans, c.id)
	}
	return ans
}
