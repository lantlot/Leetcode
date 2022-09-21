package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//给定一个整数数组 nums 和一个正整数 k，找出是否有可能把这个数组分成 k 个非空子集，其总和都相等。
//
//
//
// 示例 1：
//
//
//输入： nums = [4, 3, 2, 3, 5, 2, 1], k = 4
//输出： True
//说明： 有可能将其分成 4 个子集（5），（1,4），（2,3），（2,3）等于总和。
//
// 示例 2:
//
//
//输入: nums = [1,2,3,4], k = 3
//输出: false
//
//
//
// 提示：
//
//
// 1 <= k <= len(nums) <= 16
// 0 < nums[i] < 10000
// 每个元素的频率在 [1,4] 范围内
//
//
// Related Topics 位运算 记忆化搜索 数组 动态规划 回溯 状态压缩 👍 795 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
var nums []int
var n, tval, k int
var hi, lo, fa = 1e9, 1e-4, 0.95
var tc = 600
var ans bool

func calc() int {
	diff := tval * k
	for i, j := 0, 0; i < n && j < k; j++ {
		cur := 0
		for i < n && cur+nums[i] <= tval {
			cur += nums[i]
			i++
		}
		diff -= cur
	}
	if diff == 0 {
		ans = true
	}
	return diff
}
func sa() {
	shuffle(nums)
	for t := hi; t > lo && !ans; t *= fa {
		a, b := rand.Intn(n), rand.Intn(n)
		if a == b {
			continue
		}
		prev := calc()
		swap(nums, a, b)
		cur := calc()
		diff := cur - prev
		if float64(diff)/t > rand.Float64() {
			swap(nums, a, b)
		}
	}
}
func shuffle(nums []int) {
	for i := n; i > 0; i-- {
		swap(nums, rand.Intn(i), i-1)
	}
}
func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
func canPartitionKSubsets(ns []int, tk int) bool {
	k = tk
	nums = ns
	tc = 600
	ans = false
	rand.Seed(568884)
	sum := 0
	for _, tn := range nums {
		sum += tn
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	n = len(nums)
	tval = sum / k
	fmt.Println(tc)
	for !ans && tc > 0 {
		tc--
		sa()
	}
	fmt.Println(tc)
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
