package main

func main() {

}

/**

1234. 替换子串得到平衡字符串
https://leetcode.cn/problems/replace-the-substring-for-balanced-string/

**思考

示例代码是最简单的一种情况。


*/
func balancedString(s string) int {

	cnt := map[byte]int{}

	for _, c := range s {
		cnt[byte(c)]++
	}

	partial := len(s) / 4

	check := func() bool {
		if cnt['Q'] > partial ||
			cnt['W'] > partial ||
			cnt['E'] > partial ||
			cnt['R'] > partial {
			return false
		} else {
			return true
		}
	}

	if check() {
		return 0
	}

	res := len(s)
	r := 0

	for l, c := range s {
		for r < len(s) && !check() {
			cnt[s[r]]--
			r += 1
		}
		if !check() {
			break
		}
		res = min(res, r-l)
		cnt[byte(c)]++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
