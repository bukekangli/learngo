package arithmetic_test

import (
	"github.com/bukekangli/learngo/arithmetic"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t.Logf("%s", arithmetic.LongestPalindrome("babad"))
	t.Logf("%s", arithmetic.LongestPalindrome("cbbd"))
	t.Logf("%s", arithmetic.LongestPalindrome("ccc"))
}

func TestWordCut(t *testing.T) {
	t.Logf("%t", arithmetic.WordBreak("Ilovebytedance", map[string]bool{"I": true, "love": true, "byte": true, "bytedance": true}))
	t.Logf("%t", arithmetic.WordBreak("applepenapple", map[string]bool{"apple": true, "pen": true}))
}

func TestLongestValidParentheses(t *testing.T) {
	t.Logf("%d", arithmetic.LongestValidParentheses(")()"))
	t.Logf("%d", arithmetic.LongestValidParentheses(")()()("))
	t.Logf("%d", arithmetic.LongestValidParentheses("()(())"))
	t.Logf("%d", arithmetic.LongestValidParentheses("(()())"))
}

func TestGenerateParenthesis(t *testing.T) {
	t.Logf("%v", arithmetic.GenerateParenthesis(3))
}

func TestTrap(t *testing.T) {
	t.Logf("%d", arithmetic.Trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	t.Logf("%d", arithmetic.Trap([]int{4, 2, 0, 3, 2, 5}))
}

func TestIsMatch(t *testing.T) {
	t.Logf("%t", arithmetic.IsMatch("aa", "a"))
	t.Logf("%t", arithmetic.IsMatch("aa", "a*"))
	t.Logf("%t", arithmetic.IsMatch("aa", ".*"))
	t.Logf("%t", arithmetic.IsMatch("aab", "c*a*b"))
}

func TestJump(t *testing.T) {
	t.Logf("%d", arithmetic.Jump([]int{2, 3, 1, 1, 4}))
}
