package main

import (
	"sort"
)

//给出 n 个数对。 在每一个数对中，第一个数字总是比第二个数字小。
//
// 现在，我们定义一种跟随关系，当且仅当 b < c 时，数对(c, d) 才可以跟在 (a, b) 后面。我们用这种形式来构造一个数对链。
//
// 给定一个数对集合，找出能够形成的最长数对链的长度。你不需要用到所有的数对，你可以以任何顺序选择其中的一些数对来构造。
//
//
//
// 示例：
//
//
//输入：[[1,2], [2,3], [3,4]]
//输出：2
//解释：最长的数对链是 [1,2] -> [3,4]
//
//
//
//
// 提示：
//
//
// 给出数对的个数在 [1, 1000] 范围内。
//
//
// Related Topics 贪心 数组 动态规划 排序 👍 318 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	ans, t := 1, pairs[0][1]
	for _, pair := range pairs[1:] {
		if pair[0] > t {
			t = pair[1]
			ans++
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
