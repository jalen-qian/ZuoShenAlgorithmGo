package class_09

import (
	"math/rand"
	"strings"
	"testing"
	"time"
)

// 测试

// 使用哈希表的方式实现和前缀树相同的功能，用于对数器测试

type Right struct {
	box map[string]int // 存储字符串以及出现的次数
}

func NewRight() *Right {
	return &Right{
		box: make(map[string]int),
	}
}

// Insert 向前缀树中插入一个单词
func (r *Right) Insert(word string) {
	if word == "" {
		return
	}
	// 没加入过，则加入，次数是1
	if _, ok := r.box[word]; !ok {
		r.box[word] = 1
	} else {
		// 加入过，次数+1
		r.box[word]++
	}
}

// Delete 从前缀树中删除一个单词
func (r *Right) Delete(word string) {
	// 没加入过，则什么都不用做，直接返回
	if _, ok := r.box[word]; !ok {
		return
	}
	// 加入次数-1
	r.box[word]--
	// 如果减到0，则删除这个word
	if r.box[word] == 0 {
		delete(r.box, word)
	}
}

// Search 搜索单词在前缀树中加入过几次
func (r *Right) Search(word string) int {
	// 直接返回value，如果没加入过，返回的也是0，符合
	return r.box[word]
}

// PrefixNumber 返回前缀树中有多少个单词是以pre开头的
func (r *Right) PrefixNumber(pre string) int {
	count := 0
	// 遍历每个字符和加入的次数
	for word, times := range r.box {
		if strings.HasPrefix(word, pre) {
			count += times
		}
	}
	return count
}

// TestTrie 测试前缀树
func TestTrie(t *testing.T) {
	// 测试10万次
	t.Log("Start testing...")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 100000; i++ {
		// 初始化一个随机的字符串数组
		strArr := generateRandomStringArray(500, 100)
		// 初始化3种前缀树
		trie1 := NewTrie1()
		trie2 := NewTrie2()
		right := NewRight()
		for _, s := range strArr {
			decide := r.Float64()
			// 25%的概率同时加入这个字符串
			if decide < 0.25 {
				trie1.Insert(s)
				trie2.Insert(s)
				right.Insert(s)
			} else if decide < 0.5 {
				// 25%概率同时删除
				trie1.Delete(s)
				trie2.Delete(s)
				right.Delete(s)
			} else if decide < 0.75 {
				// 25%的概率同时查询字符串加入的次数
				ans1 := trie1.Search(s)
				ans2 := trie2.Search(s)
				ans3 := right.Search(s)
				if ans1 != ans2 || ans2 != ans3 {
					t.Errorf("Fucking Fucked!!同时搜索结果不一致：\n ans1:%d\n ans2:%d\n ans3:%d\n", ans1, ans2, ans3)
					return
				}
			} else {
				// 25%概率同时查询前缀出现的次数
				ans1 := trie1.PrefixNumber(s)
				ans2 := trie2.PrefixNumber(s)
				ans3 := right.PrefixNumber(s)
				if ans1 != ans2 || ans2 != ans3 {
					t.Errorf("Fucking Fucked!!同时查询前缀的结果不一致：\n ans1:%d\n ans2:%d\n ans3:%d\n", ans1, ans2, ans3)
					return
				}
			}
		}
	}
	t.Log("Great!!!")
}

// 初始化一个给定长度的随机的字符串，取值范围是小写字母
func generateRandomString(r *rand.Rand, maxLength int) string {
	length := r.Intn(maxLength) + 1 // [1, maxLength]
	s := make([]rune, length)
	// 从0-25之间取一个随机数
	for i := 0; i < length; i++ {
		s[i] = r.Int31n(26) + 'a'
	}
	return string(s)
}

// 初始化一个随机字符串数组
func generateRandomStringArray(maxArrLength, maxStrLength int) []string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	arrLen := r.Intn(maxArrLength) + 1 // [1, maxArrLength]
	var ans []string
	for i := 0; i < arrLen; i++ {
		ans = append(ans, generateRandomString(r, maxStrLength))
	}
	return ans
}
