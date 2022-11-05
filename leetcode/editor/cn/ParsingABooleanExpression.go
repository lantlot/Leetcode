package main

import "fmt"

//给你一个以字符串形式表述的 布尔表达式（boolean） expression，返回该式的运算结果。
//
// 有效的表达式需遵循以下约定：
//
//
// "t"，运算结果为 True
// "f"，运算结果为 False
// "!(expr)"，运算过程为对内部表达式 expr 进行逻辑 非的运算（NOT）
// "&(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 与的运算（AND）
// "|(expr1,expr2,...)"，运算过程为对 2 个或以上内部表达式 expr1, expr2, ... 进行逻辑 或的运算（OR）
//
//
//
//
// 示例 1：
//
// 输入：expression = "!(f)"
//输出：true
//
//
// 示例 2：
//
// 输入：expression = "|(f,t)"
//输出：true
//
//
// 示例 3：
//
// 输入：expression = "&(t,f)"
//输出：false
//
//
// 示例 4：
//
// 输入：expression = "|(&(t,f,t),!(t))"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= expression.length <= 20000
// expression[i] 由 {'(', ')', '&', '|', '!', 't', 'f', ','} 中的字符组成。
// expression 是以上述形式给出的有效表达式，表示一个布尔值。
//
//
// Related Topics 栈 递归 字符串 👍 120 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func parseBoolExpr(expression string) bool {
	if expression[0] == 'f' {
		return false
	}
	if expression[0] == 't' {
		return true
	}
	switch expression[0] {
	case '&':
		return and(expression[1:])
	case '|':
		return or(expression[1:])
	case '!':
		return not(expression[1:])
	}
	return true
}
func findExp(s string) (string, int) {
	c := 1
	i := 1
	for i < len(s) {
		if s[i] == ')' {
			c--
		}
		if s[i] == '(' {
			c++
		}
		if s[i] == ',' && c == 1 {
			break
		}
		i++
	}
	return s[:i], i
}
func and(s string) bool {
	exp, i := findExp(s[1 : len(s)-1])
	for {
		fmt.Println("and", exp, i, s)
		if !parseBoolExpr(exp) {
			return false
		}
		//表达式如果未结束至少需要有一个逗号
		if i < len(s)-2 {
			expt, t := findExp(s[i+2:])
			exp = expt
			i += t + 1
		} else {
			break
		}
	}
	return true
}
func or(s string) bool {
	exp, i := findExp(s[1 : len(s)-1])
	for {
		fmt.Println("or", i, exp)
		if parseBoolExpr(exp) {
			return true
		}
		if i < len(s)-2 {
			expt, t := findExp(s[i+2:])
			exp = expt
			i += t + 1
		} else {
			break
		}
	}
	return false
}

func not(s string) bool {
	return !parseBoolExpr(s[1 : len(s)-1])
}

//leetcode submit region end(Prohibit modification and deletion)
