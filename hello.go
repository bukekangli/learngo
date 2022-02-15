package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func longestPalindrome(s string) string {
	startEndIndex := []int{0, 0}
	fmt.Println(startEndIndex)
	for i, _ := range s {
		for j := 1; j <= i; j++ {
			fmt.Println(i + j)
			if len(s) <= i+j {
				break
			}
			if s[i-j] == s[i+j] {
				if startEndIndex[1]-startEndIndex[0] < 2*j {
					startEndIndex[0], startEndIndex[1] = i-j, i+j
				}
			} else {
				break
			}
		}

	}
	for i, _ := range s {
		for z := 1; z <= i+1; z++ {
			if len(s) <= i+z {
				break
			}
			if s[i-z-1] == s[i+z] {
				if startEndIndex[1]-startEndIndex[0] < 2*z-1 {
					startEndIndex[0], startEndIndex[1] = i-z+1, i+z
				}
			} else {
				break
			}
		}
	}
	return s[startEndIndex[0] : startEndIndex[1]+1]
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	length := len(nums)
	var res [][]int
	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := length - 1
		for second := first + 1; second < length; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[first]+nums[second]+nums[third] > 0 {
				third--
			}
			if second >= third {
				break
			}
			if nums[first]+nums[second]+nums[third] == 0 {
				var tempRes []int
				tempRes = append(tempRes, nums[first])
				tempRes = append(tempRes, nums[second])
				tempRes = append(tempRes, nums[third])
				res = append(res, tempRes)
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for first := 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		four := length - 1
		for {
			third := four - 1
			for second := first + 1; second < length; second++ {
				if second > first+1 && nums[second] == nums[second-1] {
					continue
				}
				for second < third && nums[first]+nums[second]+nums[third]+nums[four] > target {
					third--
				}
				if second >= third {
					break
				}
				if nums[first]+nums[second]+nums[third]+nums[four] == target {
					res = append(res, []int{nums[first], nums[second], nums[third], nums[four]})
				}
			}
			for {
				four--
				if four-first < 3 {
					break
				}
				if nums[four] == nums[four+1] {
					continue
				} else {
					break
				}
			}
			if four-first < 3 {
				break
			}
		}
	}
	return res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var resCursor *ListNode
	resCursor = &ListNode{Val: 0, Next: nil}
	res = resCursor
	for {
		if l1 == nil {
			resCursor.Next = l2
			break
		}
		if l2 == nil {
			resCursor.Next = l1
			break
		}
		if l1.Val < l2.Val {
			resCursor.Next = l1
			l1 = l1.Next
		} else {
			resCursor.Next = l2
			l2 = l2.Next
		}
		resCursor = resCursor.Next
	}
	return res.Next
}

func generateParenthesis(n int) []string {
	var res []string
	var generate func(str []string, left, right int)
	generate = func(str []string, left, right int) {
		if len(str) == 2*n {
			res = append(res, strings.Join(str, ""))
			return
		}
		if left < n {
			str = append(str, "(")
			generate(str, left+1, right)
			str = str[:len(str)-1]
		}
		if right < left {
			str = append(str, ")")
			generate(str, left, right+1)
			str = str[:len(str)-1]
		}
	}
	generate([]string{}, 0, 0)
	return res
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	originHead := head
	var h []*ListNode
	for i := 0; i < k; i++ {
		if head == nil {
			return originHead
		}
		h = append(h, head)
		head = head.Next
	}
	var cursor *ListNode
	newHead := &ListNode{0, nil}
	cursor = newHead
	for j := len(h) - 1; j >= 0; j-- {
		cursor.Next = h[j]
		cursor = cursor.Next
	}
	cursor.Next = reverseKGroup(head, k)
	return newHead.Next
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast := 1
	slow := 1
	length := len(nums)
	for fast < length {
		if nums[slow-1] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeElement(nums []int, val int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	head := 0
	tail := length - 1
	for head < tail {
		if nums[head] != val {
			head++
			continue
		}
		if nums[tail] != val {
			nums[head] = nums[tail]
			head++
		}
		tail--
	}
	return head
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	length := len(haystack)
	lenNeedle := len(needle)
	for i := 0; i <= length-lenNeedle; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}

func divide(dividend int, divisor int) int {
	isNeg := true
	if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
		isNeg = false
	}
	if divisor == -1 {
		if dividend == -1<<31 {
			return 1<<31 - 1
		} else {
			return -dividend
		}
	}
	if dividend < 0 {
		dividend = 0 - dividend
	}
	if divisor < 0 {
		divisor = 0 - divisor
	}
	if dividend < divisor {
		return 0
	}
	count := 1
	preCount := 1
	doubleDivisor := divisor
	preDoubleDivisor := doubleDivisor
	for doubleDivisor < dividend {
		preCount = count
		count += count
		preDoubleDivisor = doubleDivisor
		doubleDivisor += doubleDivisor
	}
	val := preCount + divide(dividend-preDoubleDivisor, divisor)
	if isNeg {
		return 0 - val
	} else {
		return val
	}
}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return nil
	}
	lenS := len(s)
	lenWord := len(words[0])
	lenW := len(words) * lenWord
	if lenS < lenW {
		return nil
	}
	var findWord func(s string, resWords []string, lenWord int) bool
	findWord = func(s string, resWords []string, lenWord int) bool {
		if len(resWords) == 0 || len(s) == 0 {
			return true
		}
		var resWords2 []string
		var match bool
		for _, word := range resWords {
			if word == s[0:lenWord] {
				// 单词存在重复
				if match {
					resWords2 = append(resWords2, word)
				}
				match = true
			} else {
				resWords2 = append(resWords2, word)
			}
		}
		if match && findWord(s[lenWord:], resWords2, lenWord) {
			return true
		} else {
			return false
		}
	}
	var res []int
	for i := 0; i+lenW <= lenS; i++ {
		if findWord(s[i:], words, lenWord) {
			res = append(res, i)
		}
	}
	return res
}

func nextPermutation(nums []int) {
	if len(nums) == 0 {
		return
	}
	var reverse func(nums []int)
	reverse = func(num []int) {
		for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
			num[i], num[j] = num[j], num[i]
		}
	}
	length := len(nums) - 1
	i := length - 1
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	for j := length; j > i && i >= 0; j-- {
		if nums[i] < nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	reverse(nums[i+1:])
}

func longestValidParentheses(s string) int {
	var strIndex []int
	var pairIndex []int
	for index, c := range s {
		if string(c) == "(" {
			strIndex = append(strIndex, index)
		} else if len(strIndex) > 0 {
			startIndex := strIndex[len(strIndex)-1]
			strIndex = strIndex[:len(strIndex)-1]
			s = s[:len(s)-1]
			pairIndex = append(pairIndex, startIndex)
			pairIndex = append(pairIndex, index)
		}
	}
	if pairIndex == nil {
		return 0
	}
	sort.Ints(pairIndex)
	var count, maxCount, preVal int
	count = 1
	preVal = -9
	for i := 0; i < len(pairIndex); i++ {
		if preVal+1 == pairIndex[i] {
			count += 1
			preVal = pairIndex[i]
		} else {
			preVal = pairIndex[i]
			count = 1
		}
		if count > maxCount {
			maxCount = count
		}
	}
	if maxCount <= 1 {
		maxCount = 0
	}
	return maxCount
}

func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[left] {
			if nums[0] <= target && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[right] {
				left = mid
			} else {
				right = mid + 1
			}
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 || target < nums[0] || target > nums[len(nums)-1] {
		return []int{-1, -1}
	}
	var binarySearch func(nums []int, alwaysLeft bool) int
	binarySearch = func(nums []int, alwaysLeft bool) int {
		left, right := 0, len(nums)-1
		ans := -1
		for left <= right {
			mid := (left + right) / 2
			if alwaysLeft {
				if target <= nums[mid] {
					right = mid - 1
					ans = mid
				} else {
					left = mid + 1
				}
			} else {
				if target < nums[mid] {
					right = mid - 1
				} else {
					left = mid + 1
					ans = mid
				}
			}
		}
		return ans
	}
	res := []int{binarySearch(nums, true), binarySearch(nums, false)}
	if res[0] > res[1] {
		return []int{-1, -1}
	}
	return res
}

func searchInsert(nums []int, target int) int {
	lenNums := len(nums)
	if lenNums == 0 {
		return 0
	}
	var ans int
	left, right := 0, lenNums-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
			ans = left
		} else {
			right = mid - 1
		}
	}
	return ans
}

func isValidSudoku(board [][]byte) bool {
	var rows, columns, boxes [9]map[int]int
	for i := 0; i < 9; i++ {
		rows[i] = make(map[int]int)
		columns[i] = make(map[int]int)
		boxes[i] = make(map[int]int)
	}
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		for columnIndex := 0; columnIndex < 9; columnIndex++ {
			num := board[rowIndex][columnIndex]
			if string(num) != "." {
				num2 := int(num)
				boxIndex := (rowIndex/3)*3 + columnIndex%3
				rows[rowIndex][num2] += 1
				columns[columnIndex][num2] += 1
				boxes[boxIndex][num2] += 1
				if rows[rowIndex][num2] > 1 || columns[columnIndex][num2] > 1 || boxes[boxIndex][num2] > 1 {
					return false
				}
			}
		}
	}
	return true
}

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int
	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j})
			} else {
				digit := b - '1'
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var resCount []int
	var res []string
	for i, s := range countAndSay(n - 1) {
		char := string(s)
		if i > 0 && char == res[len(res)-1] {
			resCount[len(res)-1] += 1
		} else {
			resCount = append(resCount, 1)
			res = append(res, char)
		}
	}
	var tmpList []string
	for i, char := range res {
		tmpList = append(tmpList, strconv.Itoa(resCount[i]))
		tmpList = append(tmpList, char)
	}
	tmp := strings.Join(tmpList, "")
	return tmp
}

func combinationSum(candidates []int, target int) [][]int {
	var comb []int
	var ans [][]int
	var dfs func(target, index int)
	dfs = func(target, index int) {
		if index == len(candidates) {
			return
		}
		if target == candidates[index] {
			ans = append(ans, append([]int{target}, comb...))
			return
		}
		dfs(target, index+1)
		if target > candidates[index] {
			comb = append(comb, candidates[index])
			dfs(target-candidates[index], index+1)
		}
	}
	dfs(target, 0)
	return ans
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var comb []int
	var ans [][]int
	var dfs func(targetS, Index int)
	var freq [][2]int
	var min func(a, b int) int
	min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for _, num := range candidates {
		if len(freq) == 0 || freq[len(freq)-1][0] != num {
			freq = append(freq, [2]int{num, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}
	dfs = func(targetS, Index int) {
		if targetS == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		if Index == len(freq) {
			return
		}
		dfs(targetS, Index+1)
		minN := min(targetS/freq[Index][0], freq[Index][1])
		for i := 1; i <= minN; i++ {
			comb = append(comb, freq[Index][0])
			dfs(targetS-freq[Index][0]*i, Index+1)
		}
		comb = comb[:len(comb)-minN]
	}
	dfs(target, 0)
	return ans
}

func findMinLackPositiveInteger(nums []int) int {
	// 时间复杂度O(n)
	// 空间复杂度O(n)
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}
	var lastNum int
	for i := 0; i < len(nums); i++ {
		if m[i+1] == false {
			return i + 1
		}
		lastNum = i + 1
	}
	return lastNum
}

func findMinLackPositiveIntegerV2(nums []int) int {
	// 时间复杂度O(n)
	// 空间复杂度O(1)
	length := len(nums)
	for i := 0; i < length; i++ {
		for nums[i] > 0 && nums[i] <= length && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i, n := range nums {
		if n-i != 1 {
			return i + 1
		}
	}
	return length + 1
}

func trap(height []int) int {
	// 动态规划法
	// ans = min(left, right) - current
	n := len(height)
	if n < 1 {
		return 0
	}
	var max func(int, int) int
	var min func(int, int) int
	max = func(i int, i2 int) int {
		if i > i2 {
			return i
		}
		return i2
	}
	min = func(i int, i2 int) int {
		if i > i2 {
			return i2
		}
		return i
	}
	leftMax := make([]int, n, n)
	rightMax := make([]int, n, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	var ans int
	for i, h := range height {
		minH := min(leftMax[i], rightMax[i])
		if minH > h {
			ans += minH - h
		}
	}
	return ans
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1, len2 := len(num1), len(num2)
	list := make([]int, len1+len2, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			index := i + j + 1
			list[index] += int(num1[i]-'0') * int(num2[j]-'0')
		}
	}
	var carry int
	for i := len1 + len2 - 1; i >= 0; i-- {
		tmp := list[i] + carry
		list[i] = tmp % 10
		carry = tmp / 10
	}
	var res string
	startZero := true
	for _, n := range list {
		if startZero && n == 0 {
			continue
		}
		startZero = false
		res += strconv.Itoa(n)
	}
	return res
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j] || dp[i-1][j-1]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

func isBigger(aStr, bStr string) bool {
	return aStr+bStr > bStr+aStr
}

func removeDuplicateLetters(s string) string {
	// 记录每个字母最后出现的下表
	lastIndex := [26]int{}
	for i, c := range s {
		lastIndex[c-'a'] = i
	}
	stack := make([]rune, 0)
	inStack := make(map[rune]bool)
	for i, c := range s {
		if inStack[c] {
			continue
		}
		for len(stack) != 0 && stack[len(stack)-1] > c && lastIndex[stack[len(stack)-1]-'a'] > i {
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, c)
		inStack[c] = true
	}
	return string(stack)
}

//func maxNumber(nums1 []int, nums2 []int, k int) []int {
//	num1Max := -1
//	num1MaxIndex := -1
//	for i, num := range nums1 {
//		if num > num1Max {
//			num1Max = num
//			num1MaxIndex = i
//		}
//	}
//	num2Max := -1
//	num2MaxIndex := -1
//	for i, num := range nums2 {
//		if num > num2Max {
//			num2Max = num
//			num2MaxIndex = i
//		}
//	}
//	stack := make([]int, 0, k)
//	isInStack := make(map[int]bool)
//	if num1Max > num2Max && len(nums1) - num1MaxIndex - 1 + len(nums2) > k {
//		stack = append(stack, num1Max)
//		isInStack[num1Max] = true
//	} else if
//}

func main() {
	//fmt.Printf("hello, world\n")
	fmt.Println(isBigger("3", "35"))
	fmt.Println(isBigger("3", "32"))
	fmt.Println(isBigger("111311", "1113"))
	bInt, _ := strconv.Atoi("b")
	aInt, _ := strconv.Atoi("a")
	fmt.Println(bInt - aInt)
	fmt.Println("removeDuplicateLetters: ", removeDuplicateLetters("cbacdcbc"))
	fmt.Println("removeDuplicateLetters: ", removeDuplicateLetters("bcabc"))

	//l := []int{4343, 1, 7, 13, 4, 234, 3234, 5, 6, 3, 8, 9, 11, 321, 224, 999, 123, 444, 333}
	//fmt.Println(arithmetic.SortOddEvenNum([]int{1,2,3,4,5,6,7,8,9,11,321,224,999,123,444,333}))
	//arithmetic.InsertSort(l)
	//arithmetic.BubbleSort(l)
	//arithmetic.SelectSort(l)
	//fmt.Println(arithmetic.MyMergeSort(l))
	//arithmetic.MergeSort(l)
	//fmt.Println(l)
	//fmt.Println(l[:3])
	//str := "cbbd"
	//fmt.Println(longestPalindrome(str))
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//fmt.Println(threeSum(nums))
	//fmt.Println(generateParenthesis(4))
	//fmt.Println(removeDuplicates([]int{1,1,2}))
	//fmt.Println(strStr("hello", "ll"))
	//s := "hello"
	//fmt.Println(s[2:4]=="ll")
	//fmt.Println(divide(2147483648, 1))
	//nextPermutation([]int{1, 2})
	//fmt.Println(longestValidParentheses(")()())"))
	//fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	//fmt.Println(countAndSay(4))
	//candidates := []int{10, 1, 2, 7, 6, 1, 5}
	//target := 8
	//fmt.Println(combinationSum2(candidates, target))
	//nums := []int{2,1,0}
	//nums := []int{3,2,-1}
	//nums := []int{7,8,9,10}
	//nums := []int{3, 4, -1, 1}
	//nums := []int{1, 2, 6, 3, 5, 4}
	//fmt.Println(findMinLackPositiveIntegerV2(nums))
	//trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	//fmt.Println(multiply("123", "456"))
	//fmt.Println(isMatch("aa", "a"))
	//fmt.Println(isMatch("aa", "*"))
	//fmt.Println(isMatch("cd", "?a"))
	//fmt.Println(isMatch("adceb", "*a*b"))
	//fmt.Println(isMatch("acdcb", "a*c?b"))
	//fmt.Println(isMatch("abcabczzzde", "*abc???de*"))
}
