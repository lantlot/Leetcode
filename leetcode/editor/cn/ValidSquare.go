package main

//给定2D空间中四个点的坐标 p1, p2, p3 和 p4，如果这四个点构成一个正方形，则返回 true 。
//
// 点的坐标 pi 表示为 [xi, yi] 。输入 不是 按任何顺序给出的。
//
// 一个 有效的正方形 有四条等边和四个等角(90度角)。
//
//
//
// 示例 1:
//
//
//输入: p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,1]
//输出: True
//
//
// 示例 2:
//
//
//输入：p1 = [0,0], p2 = [1,1], p3 = [1,0], p4 = [0,12]
//输出：false
//
//
// 示例 3:
//
//
//输入：p1 = [1,0], p2 = [-1,0], p3 = [0,1], p4 = [0,-1]
//输出：true
//
//
//
//
// 提示:
//
//
// p1.length == p2.length == p3.length == p4.length == 2
// -10⁴ <= xi, yi <= 10⁴
//
//
// Related Topics 几何 数学 👍 152 👎 0

func main() {
}

//leetcode submit region begin(Prohibit modification and deletion)

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	kv := make(map[int]int, 6)
	kv[cal(p1, p2)]++
	kv[cal(p1, p3)]++
	kv[cal(p1, p4)]++
	kv[cal(p2, p3)]++
	kv[cal(p2, p4)]++
	kv[cal(p3, p4)]++
	if len(kv) > 2 {
		return false
	}
	s := "1234fff"
	for i, c := range s {

	}
	f1, f2 := false, false
	for k, v := range kv {
		if k == 0 {
			return false
		}
		if v == 2 {
			f2 = true
		}
		if v == 4 {
			f1 = true
		}
	}
	return f1 && f2
}
func cal(a, b []int) int {
	l1, l2 := b[0]-a[0], b[1]-a[1]
	return l1*l1 + l2*l2
}

//leetcode submit region end(Prohibit modification and deletion)
