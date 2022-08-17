package main

import "fmt"

//给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//输出：15
//
//
// 示例 2：
//
//
//输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
//输出：19
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 10⁴] 之间。
// 1 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 134 👎 0

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
var deepest = 0
var sum int

func deepestLeavesSum(root *TreeNode) int { //递归 DFS 有优化空间
	sum = 0
	deepest = 0
	trave(root, 0)
	return sum
}
func trave(node *TreeNode, l int) {
	if node.Left == nil && node.Right == nil && l == deepest {
		sum += node.Val
		fmt.Println(sum)
		return
	}
	if l == deepest {
		deepest++
		sum = 0
	}
	if node.Left != nil {
		trave(node.Left, l+1)
	}
	if node.Right != nil {
		trave(node.Right, l+1)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
