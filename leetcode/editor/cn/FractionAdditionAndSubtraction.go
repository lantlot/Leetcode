package main

import (
	"fmt"
	"strconv"
)

//给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
//
// 这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2 应该被转换为
// 2/1。
//
//
//
// 示例 1:
//
//
//输入: expression = "-1/2+1/2"
//输出: "0/1"
//
//
// 示例 2:
//
//
//输入: expression = "-1/2+1/2+1/3"
//输出: "1/3"
//
//
// 示例 3:
//
//
//输入: expression = "1/3-1/2"
//输出: "-1/6"
//
//
//
//
// 提示:
//
//
// 输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
// 输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
// 输入只包含合法的最简分数，每个分数的分子与分母的范围是 [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
// 输入的分数个数范围是 [1,10]。
// 最终结果的分子与分母保证是 32 位整数范围内的有效整数。
//
//
// Related Topics 数学 字符串 模拟 👍 108 👎 0

func main() {
	fmt.Println(fractionAddition("1/3-1/2"))
}

//leetcode submit region begin(Prohibit modification and deletion)

func fractionAddition(expression string) string {
	isnega := false
	ansn, ansd, nume, den := 0, 0, 0, 0
	now := &nume
	for _, v := range expression {
		if v == '-' {
			if nume != 0 {
				if isnega {
					nume = -nume
				}
				if ansd == 0 {
					ansn, ansd = nume, den
					nume, den = 0, 0
				} else {
					lcm := lcm(ansd, den)
					ansn = ansn*lcm/ansd + nume*lcm/den
					ansd = lcm
					nume, den = 0, 0
				}
			}
			isnega = true
			now = &nume
		}
		if v >= '0' && v <= '9' {
			*now *= 10
			*now += int(v - '0')
		}
		if v == '/' {
			now = &den
		}
		if v == '+' {
			if nume != 0 {
				if isnega {
					nume = -nume
				}
				if ansd == 0 {
					ansn, ansd = nume, den
					nume, den = 0, 0
				} else {
					lcm := lcm(ansd, den)
					ansn = ansn*lcm/ansd + nume*lcm/den
					ansd = lcm
					nume, den = 0, 0
				}
			}
			isnega = false
			now = &nume
		}
	}
	if nume != 0 {
		if isnega {
			nume = -nume
		}
		if ansd == 0 {
			ansn, ansd = nume, den
			nume, den = 0, 0
		} else {
			lcm := lcm(ansd, den)
			ansn = ansn*lcm/ansd + nume*lcm/den
			ansd = lcm
		}
	}
	gcd := gcd(ansn, ansd)
	if gcd < 0 {
		gcd = -gcd
	}
	return strconv.Itoa(ansn/gcd) + "/" + strconv.Itoa(ansd/gcd)
}
func gcd(a, b int) int {
	tmp := a % b
	if tmp == 0 {
		return b
	} else {
		return gcd(b, tmp)
	}
}
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

//leetcode submit region end(Prohibit modification and deletion)
