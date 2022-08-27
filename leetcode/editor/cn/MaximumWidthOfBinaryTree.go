package main

import "fmt"

//给你一棵二叉树的根节点 root ，返回树的 最大宽度 。
//
// 树的 最大宽度 是所有层中最大的 宽度 。
//
//
//
// 每一层的 宽度 被定义为该层最左和最右的非空节点（即，两个端点）之间的长度。将这个二叉树视作与满二叉树结构相同，两端点间会出现一些延伸到这一层的
//null 节点，这些 null 节点也计入长度。
//
//
//
// 题目数据保证答案将会在 32 位 带符号整数范围内。
//
//
//
// 示例 1：
//
//
//输入：root = [1,3,2,5,3,null,9]
//输出：4
//解释：最大宽度出现在树的第 3 层，宽度为 4 (5,3,null,9) 。
//
//
// 示例 2：
//
//
//输入：root = [1,3,2,5,null,null,9,6,null,7]
//输出：7
//解释：最大宽度出现在树的第 4 层，宽度为 7 (6,null,null,null,null,null,7) 。
//
//
// 示例 3：
//
//
//输入：root = [1,3,2,5]
//输出：2
//解释：最大宽度出现在树的第 2 层，宽度为 2 (3,2) 。
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [1, 3000]
// -100 <= Node.val <= 100
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 446 👎 0

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
var lm map[int]int
var max, ml = 0, 0

func widthOfBinaryTree(root *TreeNode) int {
	lm = make(map[int]int, 16)
	max = 0
	dfs(root, 0, 1)
	return max + 1
}
func dfs(node *TreeNode, l, loc int) {
	if lm[l] == 0 || lm[l] > loc {
		lm[l] = loc
	} else if loc-lm[l] > max {
		max = loc - lm[l]
	}
	if l > ml {
		ml = l
	}
	if node.Left != nil {
		dfs(node.Left, l+1, loc*2-1)
	}
	if node.Right != nil {
		dfs(node.Right, l+1, loc*2)
	}
}

//leetcode submit region end(Prohibit modification and deletion)