package main

import "fmt"

//给定一个由不同正整数的组成的非空数组 nums ，考虑下面的图：
//
//
// 有 nums.length 个节点，按从 nums[0] 到 nums[nums.length - 1] 标记；
// 只有当 nums[i] 和 nums[j] 共用一个大于 1 的公因数时，nums[i] 和 nums[j]之间才有一条边。
//
//
// 返回 图中最大连通组件的大小 。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入：nums = [4,6,15,35]
//输出：4
//
//
// 示例 2：
//
//
//
//
//输入：nums = [20,50,9,63]
//输出：2
//
//
// 示例 3：
//
//
//
//
//输入：nums = [2,3,6,7,4,12,21,39]
//输出：8
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 2 * 10⁴
// 1 <= nums[i] <= 10⁵
// nums 中所有值都 不同
//
//
// Related Topics 并查集 数组 数学 👍 138 👎 0

func main() {
	fmt.Println(largestComponentSize([]int{65, 35, 43, 76, 15, 11, 81, 22, 55, 92, 31}))
}

// leetcode submit region begin(Prohibit modification and deletion)
func largestComponentSize(nums []int) int {
	m1, m2 := make(map[int]int, 32), make(map[int]map[int]int, 32)
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			l1 := m1[nums[i]]
			if l1 == 0 {
				l1 = i + 1
				m1[nums[i]] = l1
				m2[l1] = make(map[int]int, 32)
				m2[l1][nums[i]] = 1
			}
			if m1[nums[i]] == m1[nums[j]] {
				continue
			}
			if gcd1(nums[i], nums[j]) != 1 {
				m1[nums[j]] = l1
				m2[l1][nums[j]] = 1
				if len(m2[l1]) > ans {
					ans = len(m2[l1])
				}
			}
		}
	}
scan:
	newLink := false
	for i, _ := range m2 {
		if ans == len(nums) {
			return ans
		}
		for j, _ := range m2 {
			if i == j {
				continue
			}
		inner:
			for num, _ := range m2[i] {
				for num2, _ := range m2[j] {
					if gcd1(num, num2) != 1 {
						newLink = true
						for k, _ := range m2[i] {
							m2[j][k] = 1
						}
						delete(m2, i)
						if len(m2[j]) > ans {
							ans = len(m2[j])
						}
						break inner
					}
				}
			}
		}
	}
	if newLink {
		newLink = false
		goto scan
	}
	return ans
}
func gcd1(a, b int) int {
	tmp := a % b
	if tmp == 0 {
		return b
	} else {
		return gcd1(b, tmp)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
