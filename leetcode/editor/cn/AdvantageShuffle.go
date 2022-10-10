package main

import (
	"fmt"
	"sort"
)

//给定两个大小相等的数组 nums1 和 nums2，nums1 相对于 nums2 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的
//数目来描述。
//
// 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [2,7,11,15], nums2 = [1,10,4,11]
//输出：[2,11,7,15]
//
//
// 示例 2：
//
//
//输入：nums1 = [12,24,8,32], nums2 = [13,25,32,11]
//输出：[24,32,8,12]
//
//
//
//
// 提示：
//
//
// 1 <= nums1.length <= 10⁵
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 10⁹
//
//
// Related Topics 贪心 数组 双指针 排序 👍 303 👎 0

func main() {
	fmt.Println(advantageCount([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func advantageCount(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	var ans []int
	var num int
	for _, n := range nums2 {
		num, nums1 = findOne(n, nums1)
		ans = append(ans, num)
	}
	return ans
}
func findOne(num int, nums []int) (int, []int) {
	if num >= nums[len(nums)-1] || num < nums[0] {
		ret := nums[0]
		nums = nums[1:]
		return ret, nums
	}
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = (l + r) / 2
		if nums[mid] > num {
			r = mid
			if r-l == 1 {
				break
			}
		} else {
			l = mid
			if r-l == 1 {
				mid = r
				break
			}
		}
	}
	ans := nums[mid]
	if mid == len(nums)-1 {
		nums = nums[:len(nums)-1]
		return ans, nums
	}
	nums = append(nums[:mid], nums[mid+1:]...)
	return ans, nums
}

//leetcode submit region end(Prohibit modification and deletion)
