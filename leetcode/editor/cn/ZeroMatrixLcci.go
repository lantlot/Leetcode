package main

//编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
//
//
//
// 示例 1：
//
// 输入：
//[
//  [1,1,1],
//  [1,0,1],
//  [1,1,1]
//]
//输出：
//[
//  [1,0,1],
//  [0,0,0],
//  [1,0,1]
//]
//
//
// 示例 2：
//
// 输入：
//[
//  [0,1,2,0],
//  [3,4,5,2],
//  [1,3,1,5]
//]
//输出：
//[
//  [0,0,0,0],
//  [0,4,5,0],
//  [0,3,1,0]
//]
//
//
// Related Topics 数组 哈希表 矩阵 👍 129 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func setZeroes(matrix [][]int) {
	r, c := make([]int, len(matrix)), make([]int, len(matrix[0]))
	for i := range matrix {
		for j, n := range matrix[i] {
			if n == 0 {
				r[i] = 1
				c[j] = 1
			}
		}
	}
	for i, n := range r {
		if n == 1 {
			matrix[i] = make([]int, len(matrix[0]))
		}
	}
	for i, n := range c {
		if n == 1 {
			for j := range matrix {
				matrix[j][i] = 0
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
