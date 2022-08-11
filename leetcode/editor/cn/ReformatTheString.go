package main

//给你一个混合了数字和字母的字符串 s，其中的字母均为小写英文字母。
//
// 请你将该字符串重新格式化，使得任意两个相邻字符的类型都不同。也就是说，字母后面应该跟着数字，而数字后面应该跟着字母。
//
// 请你返回 重新格式化后 的字符串；如果无法按要求重新格式化，则返回一个 空字符串 。
//
//
//
// 示例 1：
//
// 输入：s = "a0b1c2"
//输出："0a1b2c"
//解释："0a1b2c" 中任意两个相邻字符的类型都不同。 "a0b1c2", "0a1b2c", "0c2a1b" 也是满足题目要求的答案。
//
//
// 示例 2：
//
// 输入：s = "leetcode"
//输出：""
//解释："leetcode" 中只有字母，所以无法满足重新格式化的条件。
//
//
// 示例 3：
//
// 输入：s = "1229857369"
//输出：""
//解释："1229857369" 中只有数字，所以无法满足重新格式化的条件。
//
//
// 示例 4：
//
// 输入：s = "covid2019"
//输出："c2o0v1i9d"
//
//
// 示例 5：
//
// 输入：s = "ab123"
//输出："1a2b3"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 500
// s 仅由小写英文字母和/或数字组成。
//
//
// Related Topics 字符串 👍 57 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func reformat(s string) string {
	var n, c, ans []rune
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			c = append(c, r)
		} else {
			n = append(n, r)
		}
	}
	if (len(n)-len(c))*(len(n)-len(c)) > 1 {
		return ""
	}
	if len(n) > len(c) {
		i := 0
		for ; i < len(c); i++ {
			ans = append(ans, n[i])
			ans = append(ans, c[i])
		}
		return string(append(ans, n[i]))
	} else if len(n) < len(c) {
		i := 0
		for ; i < len(n); i++ {
			ans = append(ans, c[i])
			ans = append(ans, n[i])
		}
		return string(append(ans, c[i]))
	} else {
		i := 0
		for ; i < len(n); i++ {
			ans = append(ans, c[i])
			ans = append(ans, n[i])
		}
		return string(ans)
	}
}

// leetcode submit region end(Prohibit modification and deletion)
func reformat2(s string) string { //复杂度略高
	i, j := 0, 0
	l := len(s)
	ans := ""
	nflag := false
	for i < l && j < l { //尽可能拼接
		if !nflag && s[i] <= 'z' && s[i] >= 'a' {
			ans += string(s[i])
			nflag = true
			i++
		} else if !nflag {
			i++
			continue
		}
		if nflag && s[j] >= '0' && s[j] <= '9' {
			ans += string(s[j])
			nflag = false
			j++
		} else if nflag {
			j++
			continue
		}
	}
	if l-len(ans) > 1 { //数字或者字母到头了还差一个以上的字符没拼进去的话就代表无解
		return ""
	}
	if l-len(ans) == 1 {
		if !nflag && i < l { //应该拼字母 剩字母 直接找到剩下那个返回即可
			for i < l {
				if s[i] <= 'z' && s[i] >= 'a' {
					ans += string(s[i])
					return ans
				}
				i++
			}
			return ""
		} else if i < l { //应该拼数字 剩字母 无解
			return ""
		} else if j < l { //需要拼数字的话只需要拼在第一个前面即可
			for j < l {
				if s[j] <= '9' && s[j] >= '0' {
					ans = string(s[j]) + ans
					return ans
				}
				i++
			}
		}
	}
	return ans
}
