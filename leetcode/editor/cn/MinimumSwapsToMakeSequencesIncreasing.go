package main

//我们有两个长度相等且不为空的整型数组 nums1 和 nums2 。在一次操作中，我们可以交换 nums1[i] 和 nums2[i]的元素。
//
//
// 例如，如果 nums1 = [1,2,3,8] ， nums2 =[5,6,7,4] ，你可以交换 i = 3 处的元素，得到 nums1 =[1,2,3
//,4] 和 nums2 =[5,6,7,8] 。
//
//
// 返回 使 nums1 和 nums2 严格递增 所需操作的最小次数 。
//
// 数组 arr 严格递增 且 arr[0] < arr[1] < arr[2] < ... < arr[arr.length - 1] 。
//
// 注意：
//
//
// 用例保证可以实现操作。
//
//
//
//
// 示例 1:
//
//
//输入: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
//输出: 1
//解释:
//交换 A[3] 和 B[3] 后，两个数组如下:
//A = [1, 3, 5, 7] ， B = [1, 2, 3, 4]
//两个数组均为严格递增的。
//
// 示例 2:
//
//
//输入: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
//输出: 1
//
//
//
//
// 提示:
//
//
// 2 <= nums1.length <= 10⁵
// nums2.length == nums1.length
// 0 <= nums1[i], nums2[i] <= 2 * 10⁵
//
//
// Related Topics 数组 动态规划 👍 393 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func minSwap(nums1 []int, nums2 []int) int {
	l := len(nums1)
	a, b := 0, 1
	for i := range nums1 {
		if i == 0 {
			continue
		}
		at, bt := a, b
		a, b = l, l
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			a = min(a, at)
			b = min(b, bt+1)
		}
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			a = min(a, bt)
			b = min(b, at+1)
		}
	}
	return min(a, b)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
