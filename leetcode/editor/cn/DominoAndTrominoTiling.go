package main

import "fmt"

//有两种形状的瓷砖：一种是 2 x 1 的多米诺形，另一种是形如 "L" 的托米诺形。两种形状都可以旋转。
//
//
//
// 给定整数 n ，返回可以平铺 2 x n 的面板的方法的数量。返回对 10⁹ + 7 取模 的值。
//
// 平铺指的是每个正方形都必须有瓷砖覆盖。两个平铺不同，当且仅当面板上有四个方向上的相邻单元中的两个，使得恰好有一个平铺有一个瓷砖占据两个正方形。
//
//
//
// 示例 1:
//
//
//
//
//输入: n = 3
//输出: 5
//解释: 五种不同的方法如上所示。
//
//
// 示例 2:
//
//
//输入: n = 1
//输出: 1
//
//
//
//
// 提示：
//
//
// 1 <= n <= 1000
//
//
// Related Topics 动态规划 👍 211 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func numTilings(n int) int {
	var mod int64 = 1000000007
	dp := make([][4]int64, n+1)
	dp[0] = [4]int64{0, 0, 0, 1}
	for i := 1; i <= n; i++ {
		d := [4]int64{}
		d[0] = dp[i-1][3] % mod
		d[1] = (dp[i-1][0] + dp[i-1][2]) % mod
		d[2] = (dp[i-1][0] + dp[i-1][1]) % mod
		d[3] = (dp[i-1][3] + dp[i-1][1] + dp[i-1][2] + dp[i-1][0]) % mod
		dp[i] = d
	}
	fmt.Println(dp)
	return int(dp[n][3] % int64(mod))
}

//leetcode submit region end(Prohibit modification and deletion)
