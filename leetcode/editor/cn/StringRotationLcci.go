package main

import "strings"

//字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。
//
// 示例1:
//
//  输入：s1 = "waterbottle", s2 = "erbottlewat"
// 输出：True
//
//
// 示例2:
//
//  输入：s1 = "aa", s2 = "aba"
// 输出：False
//
//
//
//
//
// 提示：
//
//
// 字符串长度在[0, 100000]范围内。
//
//
// 说明:
//
//
// 你能只调用一次检查子串的方法吗？
//
//
// Related Topics 字符串 字符串匹配 👍 208 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1 += s1
	if strings.Index(s1, s2) != -1 {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
