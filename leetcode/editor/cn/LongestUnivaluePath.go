package main

//给定一个二叉树的
// root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
//
// 两个节点之间的路径长度 由它们之间的边数表示。
//
//
//
// 示例 1:
//
//
//
//
//输入：root = [5,4,5,1,1,5]
//输出：2
//
//
// 示例 2:
//
//
//
//
//输入：root = [1,4,5,4,4,5]
//输出：2
//
//
//
//
// 提示:
//
//
// 树的节点数的范围是
// [0, 10⁴]
// -1000 <= Node.val <= 1000
// 树的深度将不超过 1000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 720 👎 0

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
var max = 0

func longestUnivaluePath(root *TreeNode) int {
	max = 0
	dfs(root)
	return max
}
func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	if node.Left == nil && node.Right == nil {
		return 0
	}
	l, r := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		l = dfs(node.Left) + 1
	} else if node.Left != nil {
		dfs(node.Left)
	}
	if node.Right != nil && node.Right.Val == node.Val {
		r = dfs(node.Right) + 1
	} else if node.Right != nil {
		dfs(node.Right)
	}
	if l+r > max {
		max = l + r
	}
	if l > r {
		return l
	} else {
		return r
	}
}

//leetcode submit region end(Prohibit modification and deletion)
