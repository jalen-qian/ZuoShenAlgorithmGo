package class_07

import (
	"fmt"
	"strings"
	"testing"
)

// 测试手写堆的功能

// 堆中存放用户
type User struct {
	name string
	age  int
}

func TestMyHeapGreater(t *testing.T) {
	u1 := &User{"张三", 56}
	u2 := &User{"李四", 10}
	u3 := &User{"王五", 35}
	u4 := &User{"赵六", 18}
	u5 := &User{"宋庄", 46}
	u6 := &User{"赵云", 2}
	u7 := &User{"张辽", 90}
	u8 := &User{"关羽", 56}
	u9 := &User{"刘备", 60}
	u10 := &User{"卧龙", 23}
	u11 := &User{"武松", 48}
	// 初始化一个小根堆，堆中的用户排序按照年龄从小到大排序，如果年龄相等，则名字小的排前面
	myHeap := NewMyHeapGreater[*User](func(a *User, b *User) bool {
		// 年龄小的排前面
		if a.age < b.age {
			return true
		} else if a.age == b.age {
			// 年龄相等，根据姓名，小的排前面
			return a.name < b.name
		} else {
			return false
		}
	})
	myHeap.Push(u1)
	myHeap.Push(u2)
	myHeap.Push(u3)
	myHeap.Push(u4)
	myHeap.Push(u5)
	myHeap.Push(u6)
	myHeap.Push(u7)
	myHeap.Push(u8)
	myHeap.Push(u9)
	myHeap.Push(u10)
	myHeap.Push(u11)
	// 打印堆，堆结构正确
	// [{2,赵云},{18,赵六},{10,李四},{56,关羽},{23,卧龙},{35,王五},{90,张辽},{56,张三},{60,刘备},{46,宋庄},{48,武松}]
	if sprintGreaterHeap(myHeap) != "[{2,赵云},{18,赵六},{10,李四},{56,关羽},{23,卧龙},{35,王五},{90,张辽},{56,张三},{60,刘备},{46,宋庄},{48,武松}]" {
		t.Error("测试失败，生成的堆结构不正确！！！")
		t.Log(sprintGreaterHeap(myHeap))
		return
	}
	// 测试是否包含某个对象 true
	if !myHeap.Contains(u5) {
		t.Error("测试失败，堆应该包含u5")
		t.Logf("是否包含赵云：%v", myHeap.Contains(u5))
		return
	}
	// 测试是否包含某个未加入过的对象 false（注意，这个张三是新new出来的，虽然年龄和姓名都一样，但不包含）
	if myHeap.Contains(&User{"张三", 56}) {
		t.Error("测试失败，堆应该不包含这个用户")
		t.Logf("是否包含张三：%v", myHeap.Contains(&User{"张三", 56}))
		return
	}
	// 测试重新调整，将李四从10岁改成85岁
	u2.age = 85
	myHeap.Resign(u2)
	// [{2,赵云},{18,赵六},{35,王五},{56,关羽},{23,卧龙},{85,李四},{90,张辽},{56,张三},{60,刘备},{46,宋庄},{48,武松}]
	t.Logf("重新调整后的堆：%s", sprintGreaterHeap(myHeap))
	// 测试移除一个对象 23 卧龙
	myHeap.Remove(u10)
	// [{2,赵云},{18,赵六},{35,王五},{56,关羽},{46,宋庄},{85,李四},{90,张辽},{56,张三},{60,刘备},{48,武松}]
	t.Logf("重新调整后的堆：%s", sprintGreaterHeap(myHeap))
	// 测试弹出2赵云
	u := myHeap.Pop()
	if u != u6 {
		t.Error("错误，弹出的不是最小的user", "弹出的：", u)
		return
	}
	// [{18,赵六},{46,宋庄},{35,王五},{56,关羽},{48,武松},{85,李四},{90,张辽},{56,张三},{60,刘备}]
	t.Logf("弹出后的堆：%s", sprintGreaterHeap(myHeap))
}

// 打印堆
func sprintGreaterHeap(myHeap *MyHeapGreater[*User]) string {
	ans := make([]string, 0)
	for i := 0; i < myHeap.heapSize; i++ {
		user := myHeap.arr[i]
		ans = append(ans, fmt.Sprintf("{%d,%s}", user.age, user.name))
	}
	return fmt.Sprintf("[%s]", strings.Join(ans, ","))
}
