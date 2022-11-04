package main

import "fmt"

//给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。
//
// 返回执行此操作后，grid 中最大的岛屿面积是多少？
//
// 岛屿 由一组上、下、左、右四个方向相连的 1 形成。
//
//
//
// 示例 1:
//
//
//输入: grid = [[1, 0], [0, 1]]
//输出: 3
//解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
//
//
// 示例 2:
//
//
//输入: grid = [[1, 1], [1, 0]]
//输出: 4
//解释: 将一格0变成1，岛屿的面积扩大为 4。
//
// 示例 3:
//
//
//输入: grid = [[1, 1], [1, 1]]
//输出: 4
//解释: 没有0可以让我们变成1，面积依然为 4。
//
//
//
// 提示：
//
//
// n == grid.length
// n == grid[i].length
// 1 <= n <= 500
// grid[i][j] 为 0 或 1
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 矩阵 👍 269 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
var lm map[int]map[int]bool
var ll map[int]int

func largestIsland(grid [][]int) (ans int) {
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dsum := make([][]int, len(grid))
	n := len(grid)
	dn := 1
	lm = make(map[int]map[int]bool)
	ll = make(map[int]int)
	for i, row := range grid {
		dsum[i] = make([]int, len(grid))
		for j := range row {
			if row[j] == 1 {
				//是0，0或者虽然不是0，0 但是上方和左方都是0的话给这里标记一个新的块标记
				if (i == 0 || grid[i-1][j] == 0) && (j == 0 || row[j-1] == 0) {
					dsum[i][j] = dn
					dn++
					//上方联通 左方不连通
				} else if i != 0 && grid[i-1][j] == 1 && (j == 0 || row[j-1] == 0) {
					dsum[i][j] = dsum[i-1][j]
					//左联通 上不连
				} else if j != 0 && row[j-1] == 1 && (i == 0 || grid[i-1][j] == 0) {
					dsum[i][j] = dsum[i][j-1]
					//全连
				} else {
					//上左是同一块
					if dsum[i][j-1] == dsum[i-1][j] {
						dsum[i][j] = dsum[i][j-1]
						//不是的话链接左边 并且记录两个块的链接情况
					} else {
						dsum[i][j] = dsum[i][j-1]
						if lm[dsum[i][j-1]] == nil {
							lm[dsum[i][j-1]] = map[int]bool{dsum[i-1][j]: true}
						} else {
							lm[dsum[i][j-1]][dsum[i-1][j]] = true
						}
					}
				}
				ll[dsum[i][j]]++
			}
		}
	}
	fmt.Println(lm)
	for i := range lm {
		if !lm[i][i] {
			lm[i] = dfs(lm[i], i, true)
		}
	}
	fmt.Println(ll)
	fmt.Println(lm)
	fmt.Println(dsum)
	for i, row := range grid {
		for j, x := range row {
			if x == 0 { // 枚举可以添加陆地的位置
				newArea := 1
				conn := map[int]bool{0: true}
				for _, d := range dir4 {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < n && 0 <= y && y < n && !conn[dsum[x][y]] {
						newArea += ll[dsum[x][y]]
						conn[dsum[x][y]] = true
						fmt.Println(dsum[x][y])
						if lm[dsum[x][y]] != nil {
							for l := range lm[dsum[x][y]] {
								fmt.Println(l)
								if !conn[l] {
									newArea += ll[l]
									conn[l] = true
								}
							}
						}

					}
				}
				fmt.Println(newArea, i, j)
				ans = max(ans, newArea)
			}
		}
	}
	if ans == 0 {
		return ll[1]
	}
	return
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func dfs(ma map[int]bool, t int, isr bool) map[int]bool {
	res := make(map[int]bool)
	if ma == nil {
		if isr {
			return nil
		}
		res[t] = true
		return res
	}
	flag := true
	for i := range ma {
		if lm[i] != nil {
			m := dfs(lm[i], i, false)
			for k := range m {
				res[k] = true
			}
			lm[i] = nil
			flag = false
		} else {
			res[i] = true
		}
	}
	if flag && !isr {
		res[t] = true
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
