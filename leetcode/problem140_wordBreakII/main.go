package main

import "fmt"

func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]struct{})
	for _, word := range wordDict {
		wordMap[word] = struct{}{}
	}
	return wordProcess(s, wordDict, wordMap, 0)
}

func wordProcess(s string, wordDict []string, wordMap map[string]struct{}, i int) []string {
	if len(wordDict) == 0 || i >= len(s) {
		return nil
	}
	var word string
	cur := i
	i++
	var ans []string
	for ; i <= len(s); i++ {
		word = s[cur:i]
		if _, ok := wordMap[word]; ok {
			if i == len(s) {
				ans = append(ans, word)
			}
			subAns := wordProcess(s, wordDict, wordMap, i)
			if len(subAns) > 0 {
				for _, sub := range subAns {
					ans = append(ans, word+" "+sub)
				}
			}
		}
	}
	return ans
}

func main() {
	ans := wordBreak("aaaaaaa", []string{"aaaa", "aa", "a"})
	fmt.Println(ans)
}
