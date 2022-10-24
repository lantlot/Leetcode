package main

//给定一个数组 nums ，将其划分为两个连续子数组 left 和 right， 使得：
//
//
// left 中的每个元素都小于或等于 right 中的每个元素。
// left 和 right 都是非空的。
// left 的长度要尽可能小。
//
//
// 在完成这样的分组后返回 left 的 长度 。
//
// 用例可以保证存在这样的划分方法。
//
//
//
// 示例 1：
//
//
//输入：nums = [5,0,3,8,6]
//输出：3
//解释：left = [5,0,3]，right = [8,6]
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1,0,6,12]
//输出：4
//解释：left = [1,1,1,0]，right = [6,12]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// 0 <= nums[i] <= 10⁶
// 可以保证至少有一种方法能够按题目所描述的那样对 nums 进行划分。
//
//
// Related Topics 数组 👍 178 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func partitionDisjoint(nums []int) int {
	max := nums[0]
	tm := 0
	ans := 1
	f := false
	for i, n := range nums {
		if n < max {
			ans = i + 1
			max = tm
		} else if n > max {
			if !f {
				f = true
				ans = i
			}
		}
		if n > tm {
			tm = n
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
