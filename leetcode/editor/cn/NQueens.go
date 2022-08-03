package main

import "fmt"

//按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。
//
// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
//
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
//
//
//
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4
//输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
//解释：如上图所示，4 皇后问题存在两个不同的解法。
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[["Q"]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 9
//
//
// Related Topics 数组 回溯 👍 1438 👎 0

func main() {
	fmt.Println(solveNQueens(8))
}

// leetcode submit region begin(Prohibit modification and deletion)
var ans [][]string

func solveNQueens(n int) [][]string {
	ans = [][]string{}
	queens := make([]int, n)
	row, diag1, diag2 := make(map[int]bool, n), make(map[int]bool, n), make(map[int]bool, n)
	backtrack(n, 0, queens, row, diag1, diag2)
	return ans
}
func backtrack(n, col int, queens []int, row, diag1, diag2 map[int]bool) {
	if col == n {
		print(queens, n)
	} else {
		for i := 0; i < n; i++ { //穷举该列的8种情况
			if row[i] {
				continue
			}
			dia1 := i + col
			if diag1[dia1] {
				continue
			}
			dia2 := i - col
			if diag2[dia2] {
				continue
			}
			row[i] = true
			diag1[dia1] = true
			diag2[dia2] = true
			queens[col] = i
			backtrack(n, col+1, queens, row, diag1, diag2) //参数传递下去递归查下一列
			row[i] = false
			diag1[dia1] = false
			diag2[dia2] = false
			queens[col] = -1
		}
	}
}
func print(queens []int, n int) {
	var res []string //单个答案
	for _, queen := range queens {
		tmp := ""
		for i := 0; i < n; i++ {
			if i == queen {
				tmp += "Q"
			} else {
				tmp += "."
			}
		}
		res = append(res, tmp)
	}
	ans = append(ans, res)
}

//leetcode submit region end(Prohibit modification and deletion)
