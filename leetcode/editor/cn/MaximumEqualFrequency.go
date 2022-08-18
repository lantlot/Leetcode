package main

//给你一个正整数数组 nums，请你帮忙从该数组中找出能满足下面要求的 最长 前缀，并返回该前缀的长度：
//
//
// 从前缀中 恰好删除一个 元素后，剩下每个数字的出现次数都相同。
//
//
// 如果删除这个元素后没有剩余元素存在，仍可认为每个数字都具有相同的出现次数（也就是 0 次）。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,2,1,1,5,3,3,5]
//输出：7
//解释：对于长度为 7 的子数组 [2,2,1,1,5,3,3]，如果我们从中删去 nums[4] = 5，就可以得到 [2,2,1,1,3,3]，里面每个数
//字都出现了两次。
//
//
// 示例 2：
//
//
//输入：nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
//输出：13
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁵
//
//
// Related Topics 数组 哈希表 👍 63 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func maxEqualFreq(nums []int) int {
	nm, fm := make(map[int]int), make(map[int]int) //mn 出现次数map fm 出现次数的次数map
	maxf, ans := 0, 0
	for i, num := range nums {
		if fm[nm[num]] > 0 {
			fm[nm[num]]--
		}
		nm[num]++
		fm[nm[num]]++
		if nm[num] > maxf {
			maxf = nm[num]
		}
		//最多出现一次满足要求 所有出现最多+1个任意满足 1个出现最多次 其他全都出现最多-1次
		if maxf == 1 || maxf*fm[maxf] == i || (maxf-1)*fm[maxf-1]+maxf == i+1 {
			ans = i + 1
		}
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
