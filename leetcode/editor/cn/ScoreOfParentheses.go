package main

import "fmt"

//给定一个平衡括号字符串 S，按下述规则计算该字符串的分数：
//
//
// () 得 1 分。
// AB 得 A + B 分，其中 A 和 B 是平衡括号字符串。
// (A) 得 2 * A 分，其中 A 是平衡括号字符串。
//
//
//
//
// 示例 1：
//
// 输入： "()"
//输出： 1
//
//
// 示例 2：
//
// 输入： "(())"
//输出： 2
//
//
// 示例 3：
//
// 输入： "()()"
//输出： 2
//
//
// 示例 4：
//
// 输入： "(()(()))"
//输出： 6
//
//
//
//
// 提示：
//
//
// S 是平衡括号字符串，且只含有 ( 和 ) 。
// 2 <= S.length <= 50
//
//
// Related Topics 栈 字符串 👍 433 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func scoreOfParentheses(s string) int {
	stack := [26]int{}
	t := 0
	for i, c := range s {
		if c == '(' {
			t++
		} else {
			if i > 0 && s[i-1] == '(' {
				stack[t]++
			} else {
				stack[t] += stack[t+1] * 2
				stack[t+1] = 0
			}
			t--
		}
	}
	fmt.Println(stack)
	return stack[1]
}

//leetcode submit region end(Prohibit modification and deletion)
