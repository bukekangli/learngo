package arithmetic

import (
	"crypto/sha256"
	"fmt"
)

// 动态规划
// 最长回文子串

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	longest := ""
	f := make([][]bool, len(s))
	for i := range f {
		f[i] = make([]bool, len(s))
		f[i][i] = true
		longest = string(s[i])
	}
	for step := 1; step < len(s); step++ {
		for i := 0; i+step < len(s); i++ {
			j := i + step
			if step == 1 {
				if s[i] == s[j] {
					f[i][j] = true
					longest = s[i : j+1]
				}
			} else {
				if f[i+1][j-1] && s[i] == s[j] {
					f[i][j] = true
					if j-i+1 > len(longest) {
						longest = s[i : j+1]
					}
				}
			}
		}
	}
	for _, i := range f {
		fmt.Println(i)
	}
	return longest
}

// WordBreak 字符串是否能被给定的单词集合拼接
// eg：Ilovebytedance 是否能被[I, love, byte, bytedance] 拼接成
func WordBreak(s string, d map[string]bool) bool {
	f := make([]bool, len(s)+1)
	f[0] = true
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if f[i] && d[s[i:j]] {
				f[j] = true
			}
		}
	}
	return f[len(s)]
}

// LongestValidParentheses 最长有效括号长度
func LongestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	longest := 0
	f := make([]int, len(s))
	f[0] = 0
	for i := 1; i < len(s); i++ {
		if s[i-1] == '(' && s[i] == ')' {
			if i >= 2 && f[i-2] > 0 {
				f[i] = f[i-2] + 2
			} else {
				f[i] = 2
			}
		} else if s[i] == ')' && f[i-1] > 0 && i-f[i-1]-1 >= 0 && s[i-f[i-1]-1] == '(' {
			if i-f[i-1]-2 >= 0 {
				f[i] = f[i-1] + f[i-f[i-1]-2] + 2
			} else {
				f[i] = f[i-1] + 2
			}
		}
		if f[i] > longest {
			longest = f[i]
		}
	}
	return longest
}

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	var f func(s string, lCount, rCount int)
	f = func(s string, lCount, rCount int) {
		if len(s) >= 2*n {
			res = append(res, s)
			return
		}
		if lCount > 0 {
			f(s+"(", lCount-1, rCount)
		}
		if rCount > 0 && (s[len(s)-1] == '(' || lCount < rCount) {
			f(s+")", lCount, rCount-1)
		}
	}
	s := "("
	f(s, n-1, n)
	return res
}

func Trap(height []int) int {
	heightLeft := make([]int, len(height))
	heightLeft[0] = height[0]
	heightRight := make([]int, len(height))
	heightRight[len(heightRight)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		heightLeft[i] = max(heightLeft[i-1], height[i-1])
	}
	for j := len(height) - 2; j >= 0; j-- {
		heightRight[j] = max(heightRight[j+1], height[j+1])
	}
	rain := 0
	for i := 1; i < len(height); i++ {
		minVal := min(heightLeft[i], heightRight[i])
		if minVal > height[i] {
			rain += minVal - height[i]
		}
	}
	return rain
}

func IsMatch(s string, p string) bool {
	match := make([]bool, len(s))
	index := 0
	for i := 0; i < len(s); {
		if index > len(p)-1 {
			break
		}
		switch p[index] {
		case '.':
			match[i] = true
			index++
		case '*':
			if s[i] == p[index-1] || p[index-1] == '.' {
				match[i] = true
			} else {
				index++
				continue
			}
		default:
			if s[i] == p[index] {
				match[i] = true
				index++
			} else if index+1 < len(p) && p[index+1] == '*' {
				index += 2
				continue
			} else {
				break
			}
		}
		i++
	}
	fmt.Println(match, index)
	return match[len(s)-1] && ((p[len(p)-1] == '*' && index == len(p)-1) || (p[len(p)-1] != '*' && index == len(p)))
}

func Jump(nums []int) int {
	f := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = make([]int, len(nums))
	}
	for i, num := range nums {
		for j := 1; j <= num && i+j < len(nums); j++ {
			f[i][i+j] = 1
		}
	}
	for i := range f {
		fmt.Println(f[i])
	}
	fmt.Println("")
	for i := 1; i < len(nums); i++ {
		for j := i - 1; j >= 0; j-- {
			step := f[0][j] + f[j][i]
			if f[0][i] == 0 {
				f[0][i] = step
			} else {
				f[0][i] = min(f[0][i], step)
			}
		}
	}
	for i := range f {
		fmt.Println(f[i])
	}
	return f[0][len(nums)-1]
}

func WordBreakV2(s string, wordDict []string) []string {
	s1 := "00091fc66b66c134ea99382c2690e90920df2c69"
	sha256.Sum224([]byte(s1))
}
