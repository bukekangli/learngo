package main

import (
	"container/heap"
	"fmt"
	"os"
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

func removeKdigits(num string, k int) string {
	res := make([]rune, 0, len(num))
	for _, n := range num {
		for len(res) > 0 && k > 0 && res[len(res)-1] > n {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, n)
	}
	firstElemetNotZero := false
	s := ""
	for i := 0; i < len(res)-k; i++ {
		if res[i] == '0' && !firstElemetNotZero {
			continue
		}
		firstElemetNotZero = true
		s += string(res[i])
	}
	if len(s) == 0 {
		return "0"
	}
	return s
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, 0, k)
	for k1 := 0; k1 < k; k1++ {
		k2 := k - k1
		if k1 > len(nums1) || k2 > len(nums2) {
			continue
		}
		nums1Part := maxList(nums1, k1)
		nums2Part := maxList(nums2, k2)
		resTemp := merge(nums1Part, nums2Part)
		if isLessThan(res, resTemp) {
			res = resTemp
		}
	}
	return res
}

func maxList(nums []int, k int) []int {
	if k == 0 {
		return nil
	}
	res := make([]int, 0, k)
	for index, n := range nums {
		for len(res) > 0 && res[len(res)-1] < n && len(res)+len(nums)-index > k {
			res = res[:len(res)-1]
		}
		res = append(res, n)
	}
	return res[:k]
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, len(nums1)+len(nums2))
	for i := range res {
		if isLessThan(nums1, nums2) {
			res[i], nums2 = nums2[0], nums2[1:]
		} else {
			res[i], nums1 = nums1[0], nums1[1:]
		}
	}
	return res
}

func isLessThan(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] != nums2[i] {
			return nums1[i] < nums2[i]
		}
	}
	return len(nums1) < len(nums2)
}

func increasingTriplet(nums []int) bool {
	// 两个数组都保存递增的
	triple1 := make([]int, 0, 2)
	triple2 := make([]int, 0, 2)
	for _, n := range nums {
		if len(triple1) == 0 {
			triple1 = append(triple1, n)
			continue
		}
		switch len(triple1) {
		case 0:
			triple1 = append(triple1, n)
		case 1:
			if n > triple1[0] {
				triple1 = append(triple1, n)
			} else if n < triple1[0] {
				triple1[0] = n
			}
		case 2:
			if n > triple1[1] {
				return true
			}
			if len(triple2) == 1 && n > triple2[0] {
				triple2 = append(triple2, n)
				triple1 = triple2
				triple2 = make([]int, 0, 2)
				continue
			}
			if triple1[0] < n && n < triple1[1] {
				triple1[1] = n
			} else if n < triple1[0] && n < triple1[1] {
				switch len(triple2) {
				case 0:
					triple2 = append(triple2, n)
				case 1:
					triple2[0] = n
				}
			}
		}
	}
	return false
}

func integerReplacement(n int) int {
	times := 0
	for n != 1 {
		switch n % 2 {
		case 1:
			times++
			timesAdd := integerReplacement(n - 1)
			timesDel := integerReplacement(n + 1)
			if timesAdd > timesDel {
				return times + timesDel
			} else {
				return times + timesAdd
			}
		case 0:
			n = n / 2
			times++
		}
	}
	return times
}

func longestPalindromeCount(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	calSingle := false
	length := 0
	for _, v := range m {
		switch v % 2 {
		case 0:
			length += v
		case 1:
			if !calSingle {
				length += v
				calSingle = true
			} else {
				length += v - 1
			}
		}

	}
	return length
}

func splitArray(nums []int, m int) int {
	sumNum := 0
	for _, num := range nums {
		sumNum += num
	}
	avg := sumNum / m
	//s := make([]int, 0, len(nums))
	val, maxVal := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if val+nums[i] > avg && m > 1 {
			//s = append(s, i)
			m--
			val = nums[i]
			if val > maxVal {
				maxVal = val
			}
			continue
		}
		val += nums[i]
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func longestPalindromeV2(s string) string {
	sub := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		sub[i] = make([]bool, len(s))
		sub[i][i] = true
	}
	maxLen, startIndex, endIndex := 0, 0, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if i+1 == j && s[i] == s[j] {
				sub[i][j] = true
			}
			if s[i] == s[j] && sub[i+1][j-1] {
				sub[i][j] = true
			}
			if sub[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				startIndex, endIndex = i, j
			}
		}
	}
	return s[startIndex : endIndex+1]
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		} else {
			return intervals[i][1] < intervals[j][1]
		}
	})
	nums := make([][]int, 0, len(intervals))
	nums = append(nums, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < nums[len(nums)-1][1] {
			if intervals[i][1] < nums[len(nums)-1][1] {
				nums[len(nums)-1] = intervals[i]
			}
			continue
		}
		nums = append(nums, intervals[i])
	}
	return len(intervals) - len(nums)
}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		} else {
			return points[i][1] < points[j][1]
		}
	})
	s := make([][]int, 0, len(points))
	s = append(s, points[0])
	for i := 1; i < len(points); i++ {
		if s[len(s)-1][1] >= points[i][0] {
			left := points[i][0]
			right := s[len(s)-1][1]
			if points[i][1] < s[len(s)-1][1] {
				right = points[i][1]
			}
			s[len(s)-1] = []int{left, right}
			continue
		}
		s = append(s, points[i])
	}
	return len(s)
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var count int
	for gIndex, sIndex := 0, 0; gIndex < len(g)-1 && sIndex < len(s)-1; sIndex++ {
		if s[sIndex] >= g[gIndex] {
			gIndex++
			count++
		}
	}
	return count
}

func findMaximizedCapital(k, w int, profits, capital []int) int {
	n := len(profits)
	type pair struct{ c, p int }
	arr := make([]pair, n)
	for i, p := range profits {
		arr[i] = pair{capital[i], p}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].c < arr[j].c })

	h := &hp{}
	for cur := 0; k > 0; k-- {
		for cur < n && arr[cur].c <= w {
			heap.Push(h, arr[cur].p)
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(int)
	}
	return w
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { fmt.Println("i am here"); return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func findMinMoves(machines []int) int {
	var avg, total, max int
	for _, val := range machines {
		total += val
		if val > max {
			max = val
		}
	}
	if total%len(machines) != 0 {
		return -1
	}
	avg = total / len(machines)
	return max - avg
}

func arrayPairSum(nums []int) int {
	quickSort(nums, 0, len(nums)-1)
	var sum int
	for i := 0; i < len(nums); {
		sum += nums[i]
		i += 2
	}
	return sum
}

func quickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q, r)
	}

}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	minValList := make([]int, 0, len(nums))
	minVal := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < minVal {
			minVal = nums[i]
		}
		minValList = append(minValList, minVal)
	}
	wrongSlice := make([]int, 0, len(nums))
	maxVal := nums[0]
	for i, num := range nums {
		if num != minValList[len(nums)-i-1] {
			wrongSlice = append(wrongSlice, i)
		}
		if maxVal > num {
			wrongSlice = append(wrongSlice, i)
		} else {
			maxVal = num
		}
	}
	if len(wrongSlice) < 2 {
		return 0
	}
	return wrongSlice[len(wrongSlice)-1] - wrongSlice[0] + 1
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	var count int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}
		left := 0
		if i != 0 {
			left = flowerbed[i-1]
		}
		right := 0
		if i != len(flowerbed)-1 {
			right = flowerbed[i+1]
		}
		if left == 0 && right == 0 {
			count++
			flowerbed[i] = 1
		}
	}
	return count >= n
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	var count int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			validSlice := make([]int, 0, len(nums))
			for k := j + 1; k < len(nums); k++ {
				if nums[k] >= nums[j] && nums[k] < nums[i]+nums[j] {
					validSlice = append(validSlice, k)
					continue
				} else {
					break
				}
			}
			for _, index := range validSlice {
				count += m[nums[index]]
			}
		}
	}
	return count
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	isNeg := false
	if num < 0 {
		isNeg = true
		num = -num
	}
	var s []string
	for num > 0 {
		s = append(s, strconv.Itoa(num%7))
		num /= 7
	}
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i]
	}
	if isNeg {
		res = "-" + res
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traversal(res *[]int, child *TreeNode) {
	if child == nil {
		return
	}
	traversal(res, child.Left)
	*res = append(*res, child.Val)
	traversal(res, child.Right)
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(&res, root)
	return res
}
func TestSlice(res []int) []int {
	res = append(res, 1)
	res[0] = 2
	return res
}

func TestMap(m map[int]bool) {
	m[1] = true
}

type Node struct {
	Val  int64
	Next *Node
}

func (n *Node) Print() {
	if n == nil {
		return
	}
	fmt.Printf("%d->", n.Val)
	n.Next.Print()
}

func reverse(n *Node) *Node {
	stack := make([]*Node, 0)
	curNode := n
	for curNode != nil {
		stack = append(stack, curNode)
		curNode = curNode.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		if i == 0 {
			stack[i].Next = nil
		} else {
			stack[i].Next = stack[i-1]
		}
	}
	return stack[len(stack)-1]
}

func reverseV2(n *Node) (*Node, *Node) {
	if n.Next == nil {
		return n, n
	}
	beginNode, nextNode := reverseV2(n.Next)
	nextNode.Next = n
	n.Next = nil
	return beginNode, n
}

func funcAppend(s []int64) {
	s2 := append(s, 10)
	s2[0] = 100
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if nums[0] >= len(nums)-1 {
		return 1
	}
	stack := make([][]int, 0, len(nums))
	stack = append(stack, []int{nums[0], 0})
	stack = append(stack, []int{nums[1], 1})
	for i := 2; i < len(nums)-1; i++ {
		stackTop := stack[len(stack)-1]
		stackSecondTop := stack[len(stack)-2]
		if stackTop[0]+stackTop[1] >= len(nums)-1 {
			break
		}
		if stackSecondTop[0]+stackSecondTop[1] >= i {
			if nums[i]+i > stackTop[0]+stackTop[1] {
				stack[len(stack)-1] = []int{nums[i], i}
			}
		} else {
			stack = append(stack, []int{nums[i], i})
		}
	}
	return len(stack)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		count := len(nums) - 1
		for count > 0 {
			l := []int{nums[i]}
			res = append(res, l)
			count--
		}
	}
	var index int
	var getElem func(int) int
	getElem = func(i int) int {
		for {
			if index == len(nums) {
				index = 0
			}
			if nums[index] != i {
				r := nums[index]
				index++
				return r
			}
			index++
		}
	}
	for i := 1; i < len(nums); i++ {
		index = i
		for j := 0; j < len(res); j++ {
			elem := getElem(res[j][0])
			res[j] = append(res[j], elem)
		}
	}
	return res
}

func permuteUnique(nums []int) [][]int {
	stack := make([]int, 0, len(nums))
	usedIndex := make([]bool, len(nums))
	unuseNum := make(map[int]int)
	for _, num := range nums {
		unuseNum[num]++
	}
	res := make([][]int, 0)
	var inner func()
	inner = func() {
		if len(stack) == len(nums) {
			res = append(res, append([]int{}, stack...))
			return
		}
		//unuseNum := make(map[int]bool)
		//for i, num := range nums {
		//	if !usedIndex[i] {
		//		unuseNum[num] = true
		//	}
		//}
		for i, num := range nums {
			if !usedIndex[i] && unuseNum[num] > 0 {
				usedIndex[i] = true
				count := unuseNum[num]
				unuseNum[num] = 0

				stack = append(stack, num)
				inner()
				stack = stack[:len(stack)-1]

				usedIndex[i] = false
				unuseNum[num] = count - 1
			}

		}
	}
	inner()
	return res
}

func canCompareType() {
	var a [4]int
	var b [4]int
	a[1] = 1
	fmt.Printf("a[%T]==b[%T] is %t\n", a, b, a == b)
	var c struct{
		m int
	}
	var d struct{
		m int
	}
	fmt.Printf("c[%T]==d[%T] is %t\n", c, d, c == d)
	//var e func(a int) error
	//var f func(a int) error
	//fmt.Printf("e[%T]==f[%T] is %t", e, f, e == f)
	//同一个make创建的或者都为nil才相等
	var g chan bool
	var h chan bool
	fmt.Printf("g[%T]==h[%T] is %t\n", g, h, g == h)
	i := make(chan bool, 10)
	j := make(chan bool, 10)
	//i <- true
	//j <- false fmt.Printf("g[%T]==j[%T] is %t\n", g, j, g == j)
	fmt.Printf("i[%T]==j[%T] is %t\n", i, j, i == j)
	nums1 := 8
	//nums2 := 8
	ptr1 := &nums1
	ptr2 := &nums1
	fmt.Printf("ptr1[%T]==ptr2[%T] is %t\n", ptr1, ptr2, ptr1 == ptr2)
	var k interface{}
	var l interface{}
	fmt.Printf("k[%T]==l[%T] is %t\n", k, l, k ==l)
	k = "abk"
	l = "abc"
	fmt.Printf("k[%T]==l[%T] is %t\n", k, l, k ==l)
	l = "abk"
	fmt.Printf("k[%T]==l[%T] is %t\n", k, l, k ==l)
	l = 10
	fmt.Printf("k[%T]==l[%T] is %t\n", k, l, k ==l)
	q := map[chan bool]int{}
	fmt.Printf("%#v\n", q)
	o := map[interface{}]int{}
	fmt.Printf("%#v\n", o)
	p := map[*int]int{}
	fmt.Printf("%#v\n", p)
}
func main() {

	canCompareType()
	os.Exit(0)
	//fmt.Printf("hello, world\n")
	fmt.Println("permuteUnique: ", permuteUnique([]int{1, 1, 3}))
	fmt.Println("permuteUnique: ", permuteUnique([]int{1, 2, 3}))
	os.Exit(0)
	fmt.Println("permute", permute([]int{1, 2, 3}))
	fmt.Println("permute", permute([]int{1, 2, 3, 4}))
	os.Exit(0)
	fmt.Println("jump: ", jump([]int{2, 3, 1, 1, 4}))
	fmt.Println("jump: ", jump([]int{1, 1, 1, 1}))
	fmt.Println("jump: ", jump([]int{0}))
	os.Exit(0)
	slice1 := []int64{1, 2, 3, 4, 5}
	slice2 := slice1[2:]
	funcAppend(slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)
	os.Exit(0)
	list := &Node{
		Val: 1,
		Next: &Node{
			Val: 2,
			Next: &Node{
				Val: 3,
				Next: &Node{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	//reverse(list).Print()
	beginNode, _ := reverseV2(list)
	beginNode.Print()
	os.Exit(0)
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val:   8,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   9,
						Left:  nil,
						Right: nil,
					},
				},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println("inorderTraversal: ", inorderTraversal(root))
	os.Exit(0)
	fmt.Println("convertToBase7: ", convertToBase7(100))
	fmt.Println("convertToBase7: ", convertToBase7(-7))
	os.Exit(0)
	fmt.Println("triangleNumber: ", triangleNumber([]int{2, 2, 3, 4}))
	fmt.Println("triangleNumber: ", triangleNumber([]int{4, 2, 3, 4}))
	os.Exit(0)
	fmt.Println("canPlaceFlowers: ", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1))
	fmt.Println("canPlaceFlowers: ", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println("canPlaceFlowers: ", canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1))
	os.Exit(0)
	fmt.Println("findUnsortedSubarray: ", findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
	fmt.Println("findUnsortedSubarray: ", findUnsortedSubarray([]int{1, 2, 3, 4}))
	fmt.Println("findUnsortedSubarray: ", findUnsortedSubarray([]int{1}))
	fmt.Println("findUnsortedSubarray: ", findUnsortedSubarray([]int{2, 3, 4, 1}))
	fmt.Println("findUnsortedSubarray: ", findUnsortedSubarray([]int{-1, -1, -1, -1}))
	fmt.Println(findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1}))
	fmt.Println(findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
	fmt.Println(findMaximizedCapital(1, 0, []int{1, 2, 3}, []int{1, 1, 2}))
	os.Exit(0)
	fmt.Println(isBigger("3", "35"))
	fmt.Println(isBigger("3", "32"))
	fmt.Println(isBigger("111311", "1113"))
	bInt, _ := strconv.Atoi("b")
	aInt, _ := strconv.Atoi("a")
	fmt.Println(bInt - aInt)
	fmt.Println("removeDuplicateLetters: ", removeDuplicateLetters("cbacdcbc"))
	fmt.Println("removeDuplicateLetters: ", removeDuplicateLetters("bcabc"))
	fmt.Println("removeKdigits: ", removeKdigits("1432219", 3))
	fmt.Println("removeKdigits: ", removeKdigits("10200", 1))
	fmt.Println("removeKdigits: ", removeKdigits("10", 2))
	fmt.Println("removeKdigits: ", removeKdigits("11", 2))
	fmt.Println("removeKdigits: ", removeKdigits("9", 1))
	fmt.Println("maxList", maxList([]int{8, 9}, 2))
	fmt.Println("maxList", maxList([]int{8, 6, 9}, 2))
	fmt.Println("maxList", maxList([]int{3, 4, 6, 5}, 2))
	fmt.Println("maxList", maxList([]int{9, 1, 2, 5, 8, 3}, 3))
	fmt.Println("merge", merge([]int{6, 5}, []int{9, 8, 3}))
	fmt.Println("merge", merge([]int{6, 3, 5}, []int{6, 7, 3}))
	fmt.Println("merge", merge([]int{6, 3, 5}, []int{6, 3, 5}))
	fmt.Println("maxNumber", maxNumber([]int{3, 4, 6, 5}, []int{9, 1, 2, 5, 8, 3}, 5))
	fmt.Println("maxNumber", maxNumber([]int{6, 7}, []int{6, 0, 4}, 5))
	fmt.Println("maxNumber", maxNumber([]int{8, 6, 9}, []int{1, 7, 5}, 3))

	fmt.Println("increasingTriplet: ", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	fmt.Println("increasingTriplet: ", increasingTriplet([]int{1, 5, 0, 4, 1, 3}))
	fmt.Println("integerReplacement: ", integerReplacement(1))
	fmt.Println("integerReplacement: ", integerReplacement(2))
	fmt.Println("integerReplacement: ", integerReplacement(3))
	fmt.Println("integerReplacement: ", integerReplacement(4))
	fmt.Println("integerReplacement: ", integerReplacement(5))
	fmt.Println("integerReplacement: ", integerReplacement(6))
	fmt.Println("integerReplacement: ", integerReplacement(7))
	fmt.Println("integerReplacement: ", integerReplacement(8))
	fmt.Println("integerReplacement: ", integerReplacement(9))
	fmt.Println("integerReplacement: ", integerReplacement(10))
	fmt.Println("integerReplacement: ", integerReplacement(11))
	fmt.Println("integerReplacement: ", integerReplacement(12))
	fmt.Println("integerReplacement: ", integerReplacement(13))
	fmt.Println("integerReplacement: ", integerReplacement(14))
	fmt.Println("integerReplacement: ", integerReplacement(15))
	fmt.Println("integerReplacement: ", integerReplacement(16))
	fmt.Println("longestPalindromeCount: ", longestPalindromeCount("abccccdd"))
	fmt.Println("longestPalindromeCount: ", longestPalindromeCount("a"))
	fmt.Println("longestPalindromeCount: ", longestPalindromeCount("bb"))
	fmt.Println("splitArray: ", splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println("splitArray: ", splitArray([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println("splitArray: ", splitArray([]int{1, 4, 4}, 3))
	fmt.Println("longestPalindromeV2: ", longestPalindromeV2("babad"))
	fmt.Println("longestPalindromeV2: ", longestPalindromeV2("cbdd"))
	fmt.Println("eraseOverlapIntervals: ", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	fmt.Println("eraseOverlapIntervals: ", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	fmt.Println("eraseOverlapIntervals: ", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	fmt.Println("eraseOverlapIntervals: ", eraseOverlapIntervals([][]int{{-52, 31}, {-73, -26}, {82, 97}, {-65, -11}, {-62, -49}, {95, 99}, {58, 95}, {-31, 49}, {66, 98}, {-63, 2}, {30, 47}, {-40, -26}}))
	fmt.Println("findMinArrowShots: ", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	fmt.Println("findMinArrowShots: ", findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println("findMinArrowShots: ", findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println("findMinArrowShots: ", findMinArrowShots([][]int{{9, 12}, {1, 10}, {4, 11}, {8, 12}, {3, 9}, {6, 9}, {6, 7}}))
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
