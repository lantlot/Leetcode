package main

//给定一棵二叉树 root，返回所有重复的子树。
//
// 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
// 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]]
//
// 示例 2：
//
//
//
//
//输入：root = [2,1,1]
//输出：[[1]]
//
// 示例 3：
//
//
//
//
//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]]
//
//
//
// 提示：
//
//
// 树中的结点数在[1,10^4]范围内。
// -200 <= Node.val <= 200
//
//
// Related Topics 树 深度优先搜索 哈希表 二叉树 👍 608 👎 0

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
type pair struct {
	node *TreeNode
	idx  int
}

var seen map[[3]int]pair
var repeat map[*TreeNode]bool
var idx = 0

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	idx = 0
	repeat = map[*TreeNode]bool{}
	seen = map[[3]int]pair{}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	tri := [3]int{node.Val, dfs(node.Left), dfs(node.Right)}
	if p, ok := seen[tri]; ok {
		repeat[p.node] = true
		return p.idx
	}
	idx++
	seen[tri] = pair{node, idx}
	return idx
}

//leetcode submit region end(Prohibit modification and deletion)
