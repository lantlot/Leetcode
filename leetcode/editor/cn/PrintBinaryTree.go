package main

import (
	"fmt"
	"strconv"
)

//给你一棵二叉树的根节点 root ，请你构造一个下标从 0 开始、大小为 m x n 的字符串矩阵 res ，用以表示树的 格式化布局 。构造此格式化布局矩
//阵需要遵循以下规则：
//
//
// 树的 高度 为 height ，矩阵的行数 m 应该等于 height + 1 。
// 矩阵的列数 n 应该等于 2ʰᵉⁱᵍʰᵗ⁺¹ - 1 。
// 根节点 需要放置在 顶行 的 正中间 ，对应位置为 res[0][(n-1)/2] 。
// 对于放置在矩阵中的每个节点，设对应位置为 res[r][c] ，将其左子节点放置在 res[r+1][c-2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] ，右子节点放置在
//res[r+1][c+2ʰᵉⁱᵍʰᵗ⁻ʳ⁻¹] 。
// 继续这一过程，直到树中的所有节点都妥善放置。
// 任意空单元格都应该包含空字符串 "" 。
//
//
// 返回构造得到的矩阵 res 。
//
//
//
//
//
// 示例 1：
//
//
//输入：root = [1,2]
//输出：
//[["","1",""],
// ["2","",""]]
//
//
// 示例 2：
//
//
//输入：root = [1,2,3,null,4]
//输出：
//[["","","","1","","",""],
// ["","2","","","","3",""],
// ["","","4","","","",""]]
//
//
//
//
// 提示：
//
//
// 树中节点数在范围 [1, 2¹⁰] 内
// -99 <= Node.val <= 99
// 树的深度在范围 [1, 10] 内
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 189 👎 0

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maxd = 0

func printTree(root *TreeNode) [][]string {
	maxd = 0
	dfs(root, 0)
	n := 2<<maxd - 1
	ans := make([][]string, maxd+1)
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]string, n)
	}
	fmt.Println(n)
	n = (n - 1) >> 1
	ans[0][n] = strconv.Itoa(root.Val)
	dfsl(root.Left, 1, ans, n)
	dfsr(root.Right, 1, ans, n)
	return ans
}
func dfsl(node *TreeNode, deep int, ans [][]string, l int) {
	if node != nil {
		fmt.Println(node.Val, maxd, deep, l)
		i := 1 << (maxd - deep)
		l = l - i
		fmt.Println(i)
		ans[deep][l] = strconv.Itoa(node.Val)
		dfsl(node.Left, deep+1, ans, l)
		dfsr(node.Right, deep+1, ans, l)
	}
}

func dfsr(node *TreeNode, deep int, ans [][]string, l int) {
	if node != nil {
		fmt.Println(node.Val, maxd, deep)
		l = l + (1 << (maxd - deep))
		ans[deep][l] = strconv.Itoa(node.Val)
		dfsl(node.Left, deep+1, ans, l)
		dfsr(node.Right, deep+1, ans, l)
	}
}
func dfs(node *TreeNode, deep int) {

	if node != nil {
		dfs(node.Right, deep+1)
		dfs(node.Left, deep+1)
		if deep > maxd {
			maxd = deep
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
