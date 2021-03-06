package dp

/*
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

    拆分时可以重复使用字典中的单词。
    你可以假设字典中没有重复的单词。

示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。

示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
*/

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordMap := make(map[string]struct{})
	for _, v := range wordDict {
		wordMap[v] = struct{}{}
	}
	contains := func(s string) bool {
		_, ok := wordMap[s]
		return ok
	}
	dp := make([]bool, n+1) //dp[i]表示s[0,...,i]能否被拆分
	dp[0] = true            //空串可以被拆分
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && contains(s[j:i]) {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
