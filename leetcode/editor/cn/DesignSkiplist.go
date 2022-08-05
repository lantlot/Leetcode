package main

//不使用任何库函数，设计一个 跳表 。
//
// 跳表 是在 O(log(n)) 时间内完成增加、删除、搜索操作的数据结构。跳表相比于树堆与红黑树，其功能与性能相当，并且跳表的代码长度相较下更短，其设计思
//想与链表相似。
//
// 例如，一个跳表包含 [30, 40, 50, 60, 70, 90] ，然后增加 80、45 到跳表中，以下图的方式操作：
//
// Artyom Kalinin [CC BY-SA 3.0], via Wikimedia Commons
//
// 跳表中有很多层，每一层是一个短的链表。在第一层的作用下，增加、删除和搜索操作的时间复杂度不超过 O(n)。跳表的每一个操作的平均时间复杂度是 O(log(
//n))，空间复杂度是 O(n)。
//
// 了解更多 : https://en.wikipedia.org/wiki/Skip_list
//
// 在本题中，你的设计应该要包含这些函数：
//
//
// bool search(int target) : 返回target是否存在于跳表中。
// void add(int num): 插入一个元素到跳表。
// bool erase(int num): 在跳表中删除一个值，如果 num 不存在，直接返回false. 如果存在多个 num ，删除其中任意一个即可。
//
//
//
// 注意，跳表中可能存在多个相同的值，你的代码需要处理这种情况。
//
//
//
// 示例 1:
//
//
//输入
//["Skiplist", "add", "add", "add", "search", "add", "search", "erase", "erase",
// "search"]
//[[], [1], [2], [3], [0], [4], [1], [0], [1], [1]]
//输出
//[null, null, null, null, false, null, true, false, true, false]
//
//解释
//Skiplist skiplist = new Skiplist();
//skiplist.add(1);
//skiplist.add(2);
//skiplist.add(3);
//skiplist.search(0);   // 返回 false
//skiplist.add(4);
//skiplist.search(1);   // 返回 true
//skiplist.erase(0);    // 返回 false，0 不在跳表中
//skiplist.erase(1);    // 返回 true
//skiplist.search(1);   // 返回 false，1 已被擦除
//
//
//
//
// 提示:
//
//
// 0 <= num, target <= 2 * 10⁴
// 调用search, add, erase操作次数不大于 5 * 10⁴
//
//
// Related Topics 设计 链表 👍 208 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
import (
	"fmt"
	"math/rand"
	"time"
)

const maxLevel = 32

func main() {
	seed := time.Now().UnixNano()
	fmt.Println(seed)
	rand.Seed(1658932193990310700)
	Skiplist := Constructor()
	Skiplist.Add(9)
	Skiplist.Add(4)
	Skiplist.Add(5)
	Skiplist.Add(6)
	Skiplist.Add(9)
	fmt.Println(Skiplist.Erase(2))
	fmt.Println(Skiplist.Erase(1))
	Skiplist.Add(2)
	fmt.Println(Skiplist.Search(7))
	fmt.Println(Skiplist.Search(4))
	Skiplist.Add(5)
	fmt.Println(Skiplist.Erase(6))
	fmt.Println(Skiplist.Search(5))
	Skiplist.Add(6)
	Skiplist.Add(7)
	Skiplist.Add(4)
	fmt.Println(Skiplist.Erase(3))
	fmt.Println(Skiplist.Search(6))
	fmt.Println(Skiplist.Erase(3))
	fmt.Println(Skiplist.Search(4))
	fmt.Println(Skiplist.Search(3))
	fmt.Println(Skiplist.Search(8))
	fmt.Println(Skiplist.Erase(7))
	fmt.Println(Skiplist.Erase(6))
	Skiplist.Search(7)
	fmt.Println(Skiplist.Erase(4))
	Skiplist.Add(1)
	Skiplist.Add(6)
	fmt.Println(Skiplist.Erase(3))
	Skiplist.Add(4)
	fmt.Println(Skiplist.Search(7))
	fmt.Println(Skiplist.Search(6))
	fmt.Println(Skiplist.Search(1))
	fmt.Println(Skiplist.Search(0))
	fmt.Println(Skiplist.Search(3))
}

type Skiplist struct {
	List []*Node
	Top  int
}
type Node struct {
	Val  int
	Next *Node
	Down *Node
}

func Constructor() Skiplist {
	return Skiplist{make([]*Node, maxLevel), 0}
}

func (this *Skiplist) Search(target int) bool {
	w := this.List[this.Top]
	for i := this.Top; i >= 0; i-- {
		if w.Val > target {
			w = this.List[i]
		}
		for w.Next != nil && w.Next.Val <= target {
			w = w.Next
		}
		if w.Val == target {
			return true
		}
		w = w.Down
	}
	return false
}

func (this *Skiplist) Add(num int) {
	n := &Node{num, nil, nil}
	w := this.List[0]
	if this.List[0] == nil {
		this.List[0] = n
	} else if this.List[0].Val > num {
		n.Next = this.List[0]
		this.List[0] = n
	} else {
		for w.Next != nil && w.Next.Val < num {
			w = w.Next
		}
		n.Next = w.Next
		w.Next = n
	}
	floor := 0
	dn := n
	for rand.Int()&1 == 0 {
		n = &Node{num, nil, dn} //这里逻辑还有问题
		floor++
		if this.List[floor] == nil {
			this.List[floor] = n
		} else if this.List[floor] != nil && this.List[floor].Val > num {
			n.Next = this.List[floor]
			this.List[floor] = n
		} else {
			w = this.List[floor]
			for w.Next != nil && w.Next.Val < num {
				w = w.Next
			}
			n.Next = w.Next
			w.Next = n
		}
		dn = n
	}
	if this.Top < floor {
		this.Top = floor
	}
}

func (this *Skiplist) Erase(num int) bool {
	w := this.List[this.Top]
	var dn *Node = nil
	if w.Val == num {
		dn = w
		for i := this.Top; i >= 0; i-- {
			if this.List[i].Next == nil {
				this.Top--
			}
			this.List[i] = this.List[i].Next
		}
	}
	for i := this.Top; i >= 0; i-- {
		if w.Val > num {
			w = this.List[i]
		}
		for w.Next != nil && w.Next.Val < num {
			w = w.Next
		}
		if w.Next != nil && w.Next.Val == num {
			if dn == nil {
				dn = w.Next
			} else {
				for w.Next != dn.Down {
					w = w.Next
				}
				dn = dn.Down
			}
			w.Next = w.Next.Next
		}
		w = w.Down
	}
	if dn == nil {
		return false
	}
	return true

}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
//leetcode submit region end(Prohibit modification and deletion)
