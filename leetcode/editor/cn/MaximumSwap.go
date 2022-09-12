package main

import "fmt"

//给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
//
// 示例 1 :
//
//
//输入: 2736
//输出: 7236
//解释: 交换数字2和数字7。
//
//
// 示例 2 :
//
//
//输入: 9973
//输出: 9973
//解释: 不需要交换。
//
//
// 注意:
//
//
// 给定数字的范围是 [0, 10⁸]
//
//
// Related Topics 贪心 数学 👍 265 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func maximumSwap(num int) int {
	nums, c := []int{}, [10]int{}
	for i := 0; num > 0; i++ {
		nums = append(nums, num%10)
		c[nums[i]]++
		num /= 10
	}
	max := 9
	for c[max] == 0 {
		max--
	}
	l := len(nums)
	fmt.Println(max, c[max])
mainLoop:
	for i := range nums[:l-1] {
		if nums[l-1-i] != max {
			for j, n := range nums {
				fmt.Println(n)
				if n == max {
					nums[l-1-i], nums[j] = nums[j], nums[l-1-i]
					break mainLoop
				}
			}
		} else {
			c[max]--
			for c[max] == 0 {
				max--
				if max == 0 {
					break
				}
			}
		}
	}
	ans, t := 0, 1
	for _, n := range nums {
		ans += t * n
		t *= 10
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
