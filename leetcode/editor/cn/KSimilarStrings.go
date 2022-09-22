package main

//对于某些非负整数 k ，如果交换 s1 中两个字母的位置恰好 k 次，能够使结果字符串等于 s2 ，则认为字符串 s1 和 s2 的 相似度为 k 。
//
// 给你两个字母异位词 s1 和 s2 ，返回 s1 和 s2 的相似度 k 的最小值。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab", s2 = "ba"
//输出：1
//
//
// 示例 2：
//
//
//输入：s1 = "abc", s2 = "bca"
//输出：2
//
//
//
//
// 提示：
//
//
// 1 <= s1.length <= 20
// s2.length == s1.length
// s1 和 s2 只包含集合 {'a', 'b', 'c', 'd', 'e', 'f'} 中的小写字母
// s2 是 s1 的一个字母异位词
//
//
// Related Topics 广度优先搜索 字符串 👍 242 👎 0

func main() {
}

// leetcode submit region begin(Prohibit modification and deletion)
func kSimilarity(s1, s2 string) (step int) {
	type pair struct {
		s string
		i int
	}
	q := []pair{{s1, 0}}
	vis := map[string]bool{s1: true}
	for n := len(s1); ; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			s, i := p.s, p.i
			if s == s2 {
				return
			}
			for i < n && s[i] == s2[i] {
				i++
			}
			t := []byte(s)
			for j := i + 1; j < n; j++ {
				if s[j] == s2[i] && s[j] != s2[j] { // 剪枝，只在 s[j] != s2[j] 时去交换
					t[i], t[j] = t[j], t[i]
					if t := string(t); !vis[t] {
						vis[t] = true
						q = append(q, pair{t, i + 1})
					}
					t[i], t[j] = t[j], t[i]
				}
			}
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
