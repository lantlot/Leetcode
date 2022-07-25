package main

import "math/bits"

//完全二叉树 是每一层（除最后一层外）都是完全填充（即，节点数达到最大）的，并且所有的节点都尽可能地集中在左侧。
//
// 设计一种算法，将一个新节点插入到一个完整的二叉树中，并在插入后保持其完整。
//
// 实现 CBTInserter 类:
//
//
// CBTInserter(TreeNode root) 使用头节点为 root 的给定树初始化该数据结构；
// CBTInserter.insert(int v) 向树中插入一个值为 Node.val == val的新节点 TreeNode。使树保持完全二叉树的状态
//，并返回插入节点 TreeNode 的父节点的值；
// CBTInserter.get_root() 将返回树的头节点。
//
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
//输入
//["CBTInserter", "insert", "insert", "get_root"]
//[[[1, 2]], [3], [4], []]
//输出
//[null, 1, 2, [1, 2, 3, 4]]
//
//解释
//CBTInserter cBTInserter = new CBTInserter([1, 2]);
//cBTInserter.insert(3);  // 返回 1
//cBTInserter.insert(4);  // 返回 2
//cBTInserter.get_root(); // 返回 [1, 2, 3, 4]
//
//
//
// 提示：
//
//
// 树中节点数量范围为 [1, 1000]
// 0 <= Node.val <= 5000
// root 是完全二叉树
// 0 <= val <= 5000
// 每个测试用例最多调用 insert 和 get_root 操作 10⁴ 次
//
//
// Related Topics 树 广度优先搜索 设计 二叉树 👍 137 👎 0
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
type CBTInserter struct {
	root *TreeNode
	cnt  int
}

func Constructor(r *TreeNode) CBTInserter {
	q := []*TreeNode{r}
	w := r
	cnt := 0
	for len(q) > 0 {
		cnt++
		w = q[0]
		q = q[1:]
		if w.Left != nil {
			q = append(q, w.Left)
		}
		if w.Right != nil {
			q = append(q, w.Right)
		}
	}
	return CBTInserter{r, cnt}
}

func (this *CBTInserter) Insert(val int) int {
	this.cnt++
	node := &TreeNode{val, nil, nil}
	w := this.root
	for i := bits.Len(uint(this.cnt)) - 2; i > 0; i-- { //-2是因为不需要最高位的1
		if this.cnt>>i&1 == 0 { //从最高位的1开始的第二位起 出现0向左 1向右 找到插入点的父节点
			w = w.Left
		} else {
			w = w.Right
		}
	}
	if this.cnt&1 == 0 {
		w.Left = node
	} else {
		w.Right = node
	}
	return w.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
//leetcode submit region end(Prohibit modification and deletion)
