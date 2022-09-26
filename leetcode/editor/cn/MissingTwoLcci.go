package main

//给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？
//
// 以任意顺序返回这两个数字均可。
//
// 示例 1:
//
// 输入: [1]
//输出: [2,3]
//
// 示例 2:
//
// 输入: [2,3]
//输出: [1,4]
//
// 提示：
//
//
// nums.length <= 30000
//
//
// Related Topics 位运算 数组 哈希表 👍 197 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func missingTwo(nums []int) []int {
	l := len(nums) + 2
	sum := (l + 1) * l / 2
	for _, n := range nums {
		sum -= n
	}
	s2 := 0
	t := sum / 2
	for _, n := range nums {
		if n <= t {
			s2 += n
		}
	}
	one := ((t + 1) * t / 2) - s2
	return []int{one, sum - one}
}

//leetcode submit region end(Prohibit modification and deletion)
