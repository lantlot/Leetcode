package main

//给你一个大小为 n x n 的二元矩阵 grid ，其中 1 表示陆地，0 表示水域。
//
// 岛 是由四面相连的 1 形成的一个最大组，即不会与非组内的任何其他 1 相连。grid 中 恰好存在两座岛 。
//
//
//
// 你可以将任意数量的 0 变为 1 ，以使两座岛连接起来，变成 一座岛 。
//
//
//
// 返回必须翻转的 0 的最小数目。
//
//
//
// 示例 1：
//
//
//输入：grid = [[0,1],[1,0]]
//输出：1
//
//
// 示例 2：
//
//
//输入：grid = [[0,1,0],[0,0,0],[0,0,1]]
//输出：2
//
//
// 示例 3：
//
//
//输入：grid = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//输出：1
//
//
//
//
// 提示：
//
//
// n == grid.length == grid[i].length
// 2 <= n <= 100
// grid[i][j] 为 0 或 1
// grid 中恰有两个岛
//
//
// Related Topics 深度优先搜索 广度优先搜索 数组 矩阵 👍 394 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func shortestBridge(grid [][]int) (step int) {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			island := []pair{}
			grid[i][j] = -1
			q := []pair{{i, j}}
			for len(q) > 0 {
				p := q[0]
				q = q[1:]
				island = append(island, p)
				for _, d := range dirs {
					x, y := p.x+d.x, p.y+d.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = -1
						q = append(q, pair{x, y})
					}
				}
			}

			q = island
			for {
				tmp := q
				q = nil
				for _, p := range tmp {
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
