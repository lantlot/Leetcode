package main

import (
	"strconv"
)

//求解一个给定的方程，将x以字符串 "x=#value" 的形式返回。该方程仅包含 '+' ， '-' 操作，变量 x 和其对应系数。
//
// 如果方程没有解，请返回 "No solution" 。如果方程有无限解，则返回 “Infinite solutions” 。
//
// 题目保证，如果方程中只有一个解，则 'x' 的值是一个整数。
//
//
//
// 示例 1：
//
//
//输入: equation = "x+5-3+x=6+x-2"
//输出: "x=2"
//
//
// 示例 2:
//
//
//输入: equation = "x=x"
//输出: "Infinite solutions"
//
//
// 示例 3:
//
//
//输入: equation = "2x=x"
//输出: "x=0"
//
//
//
//
// 提示:
//
//
// 3 <= equation.length <= 1000
// equation 只有一个 '='.
// equation 方程由整数组成，其绝对值在 [0, 100] 范围内，不含前导零和变量 'x' 。
//
//
// Related Topics 数学 字符串 模拟 👍 141 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func solveEquation(equation string) string {
	nc, xc, pe := 0, 0, false    //pe:是否经过equal符 nc:所有常量值合并后的值 xc:所有未知量合并
	n, op, zx := 0, false, false //操作符 false为加
	for _, c := range equation {
		switch {
		case c >= '0' && c <= '9':
			n = n*10 + int(c-'0')
			if n == 0 {
				zx = true
			}
		case c == '-':
			if op == pe {
				nc += n
			} else {
				nc -= n
			}
			op = true
			zx = false
			n = 0
		case c == '+':
			if op == pe {
				nc += n
			} else {
				nc -= n
			}
			n = 0
			zx = false
			op = false
		case c == 'x':
			if zx {
				zx = false
				continue
			}
			if n == 0 {
				n = 1
			}
			if op == pe {
				xc += n
			} else {
				xc -= n
			}
			n = 0
		case c == '=':
			if op == pe {
				nc += n
			} else {
				nc -= n
			}
			zx = false
			n = 0
			pe = true
			op = false //默认为+号
		}
	}
	if op == pe {
		nc += n
	} else {
		nc -= n
	}
	if xc == 0 && nc == 0 {
		return "Infinite solutions" //无穷解 通过合并后可以使得x的系数变为0 常数项合并为0
	}
	if xc == 0 {
		return "No solution" // 无解 x系数为0但是常数项合并不为0
	}
	return "x=" + strconv.Itoa(-nc/xc)
}

//leetcode submit region end(Prohibit modification and deletion)
