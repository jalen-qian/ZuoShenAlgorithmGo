package class_11_12

import (
	"ZuoShenAlgorithmGo/class_03"
	"strconv"
	"strings"
)

// 序列化和反序列化二叉树

// 给定一个二叉树的头节点，返回序列化后的字符串
// 规则：节点值之间用逗号隔开，空节点用#表示

// SerializeAndReconstructBT 实现二叉树的序列化和反序列化
type SerializeAndReconstructBT struct{}

// PreSerialize 先序方式序列化成字符串
func (s *SerializeAndReconstructBT) PreSerialize(head *Node) string {
	queue := class_03.NewMyQueue[string]()
	// 序列化
	s.preSerialize(head, queue)
	// 将队列转换成字符串
	return s.queueToStr(queue)
}

// 递归方式实现先序序列化
func (s *SerializeAndReconstructBT) preSerialize(head *Node, queue *class_03.MyQueue[string]) {
	if head == nil {
		queue.Push("#")
	} else {
		// 先入队
		queue.Push(strconv.Itoa(head.Value))
		// 再递归左子树和右子树
		s.preSerialize(head.Left, queue)
		s.preSerialize(head.Right, queue)
	}
}

// InSerialize 中序方式序列化成字符串
func (s *SerializeAndReconstructBT) InSerialize(head *Node) string {
	queue := class_03.NewMyQueue[string]()
	// 序列化
	s.inSerialize(head, queue)
	// 将队列转换成字符串
	return s.queueToStr(queue)
}

// 递归方式实现中序序列化
func (s *SerializeAndReconstructBT) inSerialize(head *Node, queue *class_03.MyQueue[string]) {
	if head == nil {
		queue.Push("#")
	} else {
		// 先执行左子树
		s.inSerialize(head.Left, queue)
		// 再入队
		queue.Push(strconv.Itoa(head.Value))
		// 再执行右子树
		s.inSerialize(head.Right, queue)
	}
}

// PosSerialize 后序方式序列化成字符串
func (s *SerializeAndReconstructBT) PosSerialize(head *Node) string {
	queue := class_03.NewMyQueue[string]()
	// 序列化
	s.posSerialize(head, queue)
	// 将队列转换成字符串
	return s.queueToStr(queue)
}

// 递归方式实现后序序列化
func (s *SerializeAndReconstructBT) posSerialize(head *Node, queue *class_03.MyQueue[string]) {
	if head == nil {
		queue.Push("#")
	} else {
		// 先执行左子树
		s.posSerialize(head.Left, queue)
		// 再执行右子树
		s.posSerialize(head.Right, queue)
		// 再入队
		queue.Push(strconv.Itoa(head.Value))
	}
}

func (s *SerializeAndReconstructBT) queueToStr(queue *class_03.MyQueue[string]) string {
	if queue == nil || queue.IsEmpty() {
		return ""
	}
	var ans string
	isFirst := true
	for !queue.IsEmpty() {
		if isFirst {
			ans += queue.Poll()
			isFirst = false
		} else {
			ans += "," + queue.Poll()
		}
	}
	return ans
}

// LevelSerialize 按层遍历序列化
func (s *SerializeAndReconstructBT) LevelSerialize(head *Node) string {
	queue := class_03.NewMyQueue[string]()
	// 序列化
	// 准备一个Node队列，用来实现按层遍历的
	nodeQueue := NewNodeQueue()
	// 头节点先入队列
	nodeQueue.Push(head)
	for !nodeQueue.IsEmpty() {
		// 先出队列一个
		head = nodeQueue.Poll()
		// 出队列就加入序列化，可能是空，要判断，空就序列化成#
		if head == nil {
			queue.Push("#")
		} else {
			queue.Push(strconv.Itoa(head.Value))
			// 左右子树入队列，不判断空，因为空也要序列化
			nodeQueue.Push(head.Left)
			nodeQueue.Push(head.Right)
		}
	}
	// 将队列转换成字符串
	return s.queueToStr(queue)
}

// 反序列化

// BuildByPreSerialize 通过先序遍历序列化后的字符串反序列化成树
func (s *SerializeAndReconstructBT) BuildByPreSerialize(preSer string) *Node {
	// 空树
	if preSer == "" || preSer == "#" {
		return nil
	}
	// 还原成序列化队列
	preQueue := s.getQueue(preSer)
	// 根据队列来构建目标树
	ans := s.buildByPreQueue(preQueue)
	return ans
}

func (s *SerializeAndReconstructBT) buildByPreQueue(queue *class_03.MyQueue[string]) *Node {
	// 从队列中取出一个，作为头节点
	head := s.buildNodeByQueue(queue)
	// 递归
	if head != nil {
		head.Left = s.buildByPreQueue(queue)
		head.Right = s.buildByPreQueue(queue)
	}
	return head
}

// BuildByPosSerialize 通过后序遍历序列化后的字符串反序列化成树
//        1
//       / \
//      2   3
//         /
//        4
// #,#,2,#,#,4,#,3,1
func (s *SerializeAndReconstructBT) BuildByPosSerialize(posSer string) *Node {
	// 空树
	if posSer == "" || posSer == "#" {
		return nil
	}
	// 后序遍历队列顺序：左 右 头， 先是左子树，再是右子树，最后是头节点
	// 而先序遍历我们已经实现了，先序遍历为：头，左，右。

	// 在改后序遍历非递归方法的代码时，我们是先将先序遍历的 头，左，右 改成 头，右，左 （这个改动很简单，交换两行代码顺序就能做到，本质是一样的）
	// 然后再用一个栈来逆序，改成了 左，右，头的顺序，也就是实现了后序遍历（每次要打印时，不打印，而是入栈，最后弹出）

	// 这里也是一样，如果将posSer的值按顺序压入一个栈中，则变成了 头，右，左的顺序，然后我们再按照先序遍历的方式处理

	// 还原成序列化栈
	posStack := s.getStack(posSer)
	// 根据队列来构建目标树
	ans := s.buildByPosStack(posStack)
	return ans
}

func (s *SerializeAndReconstructBT) buildByPosStack(posStack *class_03.MyStack[string]) *Node {
	// 当前栈是 头 右 左的顺序，先弹出的是头，再弹出的是右子树，再弹出的是左子树
	// 1. 弹出头的值
	strValue := posStack.Pop()
	if strValue == "#" {
		return nil
	}
	// 先构建头节点
	headValue, _ := strconv.Atoi(strValue)
	head := &Node{Value: headValue}
	// 递归，构建左右子节点，注意栈中的顺序是 头 右 左，所以构建时，也要遵循先右后左的顺序
	if head != nil {
		head.Right = s.buildByPosStack(posStack)
		head.Left = s.buildByPosStack(posStack)
	}
	return head
}

// BuildByLevelSerialize 按层遍历的方式反序列化
//        1
//       / \
//      2   3
//         /
//        4
// 1,2,3,#,#,4,#,#,#
func (s *SerializeAndReconstructBT) BuildByLevelSerialize(preSer string) *Node {
	// 空树
	if preSer == "" || preSer == "#" {
		return nil
	}
	// 还原成序列化队列
	preQueue := s.getQueue(preSer)
	// 根据队列来构建目标树
	ans := s.buildByLevelQueue(preQueue)
	return ans
}

func (s *SerializeAndReconstructBT) buildByLevelQueue(queue *class_03.MyQueue[string]) *Node {
	head := s.buildNodeByQueue(queue)
	// 构建一个Node队列
	nodeQueue := NewNodeQueue()
	// 头节点入队列
	nodeQueue.Push(head)
	for !nodeQueue.IsEmpty() {
		// 头节点出队列
		node := nodeQueue.Poll()
		// 构建当前head的左子树 queue中的下两个一定是当前节点的左右子树
		left := s.buildNodeByQueue(queue)
		right := s.buildNodeByQueue(queue)
		if node != nil {
			node.Left = left
			node.Right = right
			if left != nil {
				nodeQueue.Push(left)
			}
			if right != nil {
				nodeQueue.Push(right)
			}
		}
	}
	return head
}

func (s *SerializeAndReconstructBT) getQueue(serialized string) *class_03.MyQueue[string] {
	queue := class_03.NewMyQueue[string]()
	for _, str := range strings.Split(serialized, ",") {
		queue.Push(str)
	}
	return queue
}

func (s *SerializeAndReconstructBT) getStack(serialized string) *class_03.MyStack[string] {
	stack := class_03.NewMyStack[string]()
	for _, str := range strings.Split(serialized, ",") {
		stack.Push(str)
	}
	return stack
}

func (s *SerializeAndReconstructBT) buildNodeByQueue(queue *class_03.MyQueue[string]) *Node {
	if queue.IsEmpty() {
		return nil
	}
	v := queue.Poll()
	if v == "#" {
		return nil
	} else {
		nodeValue, _ := strconv.Atoi(v)
		return &Node{Value: nodeValue}
	}
}
