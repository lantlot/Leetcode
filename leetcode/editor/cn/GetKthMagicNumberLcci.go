package main

//有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，
//5，7，9，15，21。
//
// 示例 1:
//
// 输入: k = 5
//
//输出: 9
//
//
// Related Topics 哈希表 数学 动态规划 堆（优先队列） 👍 209 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func getKthMagicNumber(k int) int {
	if k == 1 {
		return 1
	}
	a, b, c := 0, 0, 0
	arr := make([]int, k+1)
	arr[0] = 1
	for i := 0; i < k; i++ {
		arr[i+1] = min(min(arr[a]*3, arr[b]*5), arr[c]*7)
		if arr[i+1] == arr[a]*3 {
			a++
		}
		if arr[i+1] == arr[b]*5 {
			b++
		}
		if arr[i+1] == arr[c]*7 {
			c++
		}
	}
	return arr[k-1]
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
