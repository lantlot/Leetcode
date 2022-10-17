package main

//给定一组 n 人（编号为 1, 2, ..., n）， 我们想把每个人分进任意大小的两组。每个人都可能不喜欢其他人，那么他们不应该属于同一组。
//
// 给定整数 n 和数组 dislikes ，其中 dislikes[i] = [ai, bi] ，表示不允许将编号为 ai 和 bi的人归入同一组。当可以用
//这种方法将所有人分进两组时，返回 true；否则返回 false。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 4, dislikes = [[1,2],[1,3],[2,4]]
//输出：true
//解释：group1 [1,4], group2 [2,3]
//
//
// 示例 2：
//
//
//输入：n = 3, dislikes = [[1,2],[1,3],[2,3]]
//输出：false
//
//
// 示例 3：
//
//
//输入：n = 5, dislikes = [[1,2],[2,3],[3,4],[4,5],[1,5]]
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2000
// 0 <= dislikes.length <= 10⁴
// dislikes[i].length == 2
// 1 <= dislikes[i][j] <= n
// ai < bi
// dislikes 中每一组都 不同
//
//
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 256 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
// 染色解 需要增加一个待定集
//
//	func possibleBipartition(n int, dislikes [][]int) bool {
//		m := make(map[int]int)
//		sort.Slice(dislikes, func(i, j int) bool {
//			return dislikes[i][0] < dislikes[j][0]
//		})
//		for _, dislike := range dislikes {
//			n1l := m[dislike[0]]
//			n2l := m[dislike[1]]
//			if n1l == 0 && n2l == 0 {
//				m[dislike[0]] = 1
//				m[dislike[1]] = -1
//			} else if n1l != 0 {
//				if n2l == 0 {
//					m[dislike[1]] = -n1l
//				} else if n2l == n1l {
//					return false
//				}
//			} else if n2l != 0 {
//				if n1l == 0 {
//					m[dislike[0]] = -n2l
//				} else if n2l == n1l {
//					return false
//				}
//			}
//			fmt.Println(m)
//		}
//		return true
//	}
var t [4020]int

func find(num int) int {
	if t[num] != num {
		t[num] = find(t[num])
	}
	return t[num]
}
func union(a, b int) {
	t[find(a)] = t[find(b)]
}
func query(a, b int) bool {
	return find(a) == find(b)
}
func possibleBipartition(n int, dislikes [][]int) bool {
	t = [4020]int{}
	for i := 0; i < n*2; i++ {
		t[i] = i
	}
	for _, dislike := range dislikes {
		a, b := dislike[0], dislike[1]
		if query(a, b) {
			return false
		}
		union(a, b+n)
		union(b, a+n)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
