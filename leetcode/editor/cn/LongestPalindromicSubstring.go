package main

import "fmt"

//给你一个字符串 s，找到 s 中最长的回文子串。
//
//
//
// 示例 1：
//
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//
//
// 示例 2：
//
//
//输入：s = "cbbd"
//输出："bb"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
//
// Related Topics 字符串 动态规划 👍 5510 👎 0

func main() {
	fmt.Println(longestPalindrome("bb"))
}

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s //长度小于2直接返回
	}
	anss, anse := 0, 0
	dp := make([][]bool, l)
	for i, _ := range dp {
		dp[i] = make([]bool, l)
	}
	for i := 0; i < l; i++ {
		dp[i][i] = true
	}
	for i := 2; i <= l; i++ { //字串长度
		for st := 0; st < l; st++ { //起始点
			en := st + i - 1
			if en >= l {
				break
			}
			if s[st] != s[en] {
				dp[st][en] = false
			} else {
				if i < 3 {
					dp[st][en] = true
				} else {
					dp[st][en] = dp[st+1][en-1]
				}
			}
			if dp[st][en] {
				anss = st
				anse = en
			}
		}
	}
	return s[anss : anse+1]
}

//leetcode submit region end(Prohibit modification and deletion)
