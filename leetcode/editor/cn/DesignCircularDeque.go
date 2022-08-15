package main

import "fmt"

//设计实现双端队列。
//
// 实现 MyCircularDeque 类:
//
//
// MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
// boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
// boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
// boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
// boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
// int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
// int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
// boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false 。
// boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
//
//
//
//
// 示例 1：
//
//
//输入
//["MyCircularDeque", "insertLast", "insertLast", "insertFront", "insertFront",
//"getRear", "isFull", "deleteLast", "insertFront", "getFront"]
//[[3], [1], [2], [3], [4], [], [], [], [4], []]
//输出
//[null, true, true, true, false, 2, true, true, true, 4]
//
//解释
//MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
//circularDeque.insertLast(1);			        // 返回 true
//circularDeque.insertLast(2);			        // 返回 true
//circularDeque.insertFront(3);			        // 返回 true
//circularDeque.insertFront(4);			        // 已经满了，返回 false
//circularDeque.getRear();  				// 返回 2
//circularDeque.isFull();				        // 返回 true
//circularDeque.deleteLast();			        // 返回 true
//circularDeque.insertFront(4);			        // 返回 true
//circularDeque.getFront();				// 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= k <= 1000
// 0 <= value <= 1000
// insertFront, insertLast, deleteFront, deleteLast, getFront, getRear, isEmpty,
// isFull 调用次数不大于 2000 次
//
//
// Related Topics 设计 队列 数组 链表 👍 184 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
type MyCircularDeque struct {
	data    []int
	f       int
	r       int
	isEmpty bool
	l       int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{make([]int, k), 0, 0, true, k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.isEmpty {
		this.data[this.f] = value
		this.isEmpty = false
		return true
	}
	i := (this.f + this.l - 1) % this.l
	if this.r != i {
		this.data[i] = value
		this.f = i
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.isEmpty { //实际上这里f=r
		this.data[this.r] = value
		this.isEmpty = false
		return true
	}
	i := (this.r + 1) % this.l
	if this.f != i {
		this.data[i] = value
		this.r = i
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.isEmpty {
		return false
	}
	if this.f == this.r {
		this.isEmpty = true
		return true
	}
	this.f = (this.f + 1) % this.l
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.isEmpty {
		return false
	}
	if this.f == this.r {
		this.isEmpty = true
		return true
	}
	this.r = (this.r + this.l - 1) % this.l
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.isEmpty {
		return -1
	}
	return this.data[this.f]
}

func (this *MyCircularDeque) GetRear() int {
	if this.isEmpty {
		return -1
	}
	return this.data[this.r]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.isEmpty
}

func (this *MyCircularDeque) IsFull() bool {
	if (this.r+1)%this.l == this.f {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
//leetcode submit region end(Prohibit modification and deletion)
