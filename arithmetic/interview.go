package arithmetic

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// 给定一个数组奇数排前、偶数排后
func SortOddEvenNum(l []int) []int {
	length := len(l)
	if length <= 1 {
		return l
	}
	oddIndex, evenIndex := length-1, 0
	findOdd, findEven := false, false
	for {
		// 从前往后找偶数
		for i := evenIndex; i < oddIndex && !findEven; i++ {
			if l[i]%2 == 0 {
				findEven = true
				evenIndex = i
			}
		}
		// 从后往前找奇数
		for j := oddIndex; j > evenIndex && !findOdd; j-- {
			if l[j]%2 == 1 {
				findOdd = true
				oddIndex = j
			}
		}
		if !findOdd || !findEven {
			break
		}
		l[evenIndex], l[oddIndex] = l[oddIndex], l[evenIndex]
		findOdd, findEven = false, false
		if oddIndex-evenIndex == 1 {
			break
		}
	}
	return l
}

func alterPrint() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	go func() {
		i := 1
		for {
			select {
			case <-letter:
				fmt.Println(i)
				i++
				fmt.Println(i)
				i++
				number <- true
			}
		}
	}()
	wait.Add(1)
	go func() {
		c := 'A'
		for {
			select {
			case <-number:
				fmt.Println(c)
				c++
				fmt.Println(c)
				c++
				if c > 'Z' {
					wait.Done()
					return
				}
				letter <- true
			}
		}
	}()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2数之和
// 7->4->5 + 8->3->2 == 5->8->5
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2, 0)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	if l1 != nil {
		carry += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}
	dummy := &ListNode{
		Val:  carry % 10,
		Next: addTwoNumbers(l1, l2, carry/10),
	}
	return dummy
}

// MyPow 实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn ）
func MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	isNeg := n < 0
	if isNeg {
		n = -n
	}
	sum := MyPow(x, n/2)
	sum = sum * sum
	if n%2 == 1 {
		sum *= x
	}
	if isNeg {
		return 1 / sum
	}
	return sum
}

func MaxSubArray(nums []int) int {
	maxSum := nums[0]
	preSum := nums[0]
	for i := 1; i < len(nums); i++ {
		preSum = max(preSum+nums[i], nums[i])
		maxSum = max(preSum, maxSum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SpiralOrder 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
func SpiralOrder(matrix [][]int) []int {
	DIRECTION = RIGHT
	iMinIndex, jMinIndex = 0, 0
	iMaxIndex = len(matrix) - 1
	jMaxIndex = len(matrix[0]) - 1
	totalNum := len(matrix) * len(matrix[0])
	i, j := 0, 0
	res := make([]int, 0, totalNum)
	res = append(res, matrix[i][j])
	totalNum--
	for totalNum > 0 {
		turn(i, j)
		i, j = next(i, j)
		res = append(res, matrix[i][j])
		totalNum--
	}
	return res
}

const (
	RIGHT int = iota
	DOWN
	LEFT
	UP
)

var (
	DIRECTION            = RIGHT
	iMinIndex, iMaxIndex = 0, 0
	jMinIndex, jMaxIndex = 0, 0
)

func turn(i, j int) {
	switch DIRECTION {
	case RIGHT:
		if j >= jMaxIndex {
			DIRECTION = DOWN
			iMinIndex++
		}
	case DOWN:
		if i >= iMaxIndex {
			DIRECTION = LEFT
			jMaxIndex--
		}
	case LEFT:
		if j <= jMinIndex {
			DIRECTION = UP
			iMaxIndex--
		}
	case UP:
		if i <= iMinIndex {
			DIRECTION = RIGHT
			jMinIndex++
		}
	}
}

func next(i, j int) (int, int) {
	switch DIRECTION {
	case RIGHT:
		return i, j + 1
	case DOWN:
		return i + 1, j
	case LEFT:
		return i, j - 1
	//case UP:
	default:
		return i - 1, j
	}
}

func CanJump(nums []int) bool {
	maxIndex := 0
	for i, num := range nums {
		if maxIndex < i {
			return false
		}
		maxIndex = max(maxIndex, i+num)
		if maxIndex >= len(nums)-1 {
			return true
		}
	}
	return false
}

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	res := make([][]int, 0)
	var interval []int
	for i := 0; i < len(intervals); i++ {
		if len(interval) == 0 {
			interval = intervals[i]
			continue
		}
		if intervals[i][0] <= interval[1] {
			if intervals[i][1] > interval[1] {
				interval[1] = intervals[i][1]
			}
		} else {
			res = append(res, interval)
			interval = intervals[i]
		}
	}
	res = append(res, interval)
	return res
}

func Insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0)
	for _, inter := range intervals {
		if len(newInterval) == 0 {
			res = append(res, inter)
			continue
		}
		if inter[1] < newInterval[0] {
			res = append(res, inter)
			continue
		}
		if newInterval[1] < inter[0] {
			res = append(res, newInterval)
			res = append(res, inter)
			newInterval = []int{}
			continue
		}
		newInterval[0] = min(newInterval[0], inter[0])
		newInterval[1] = max(newInterval[1], inter[1])
	}
	if len(newInterval) != 0 {
		res = append(res, newInterval)
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func LengthOfLastWord(s string) int {
	Len := 0
	b := false
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			Len++
			b = true
			continue
		}
		if b {
			break
		}
	}
	return Len
}

func GenerateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	DIRECTION = RIGHT
	iMinIndex, jMinIndex = 0, 0
	iMaxIndex, jMaxIndex = n-1, n-1
	res := make([][]int, n)
	for i, _ := range res {
		res[i] = make([]int, n)
	}
	i, j := 0, 0
	start := 1
	res[i][j] = start
	for count := n*n - 1; count > 0; count-- {
		start++
		turn(i, j)
		i, j = next(i, j)
		res[i][j] = start
	}
	return res
}

func GetPermutation(n int, k int) string {
	res := make([]string, 0)
	ignoreMap := make(map[int]interface{})
	getValByIndex := func(n int, index int) string {
		res := 0
		step := 0
		if index > n {
			index %= n
		}
		for i := 1; i <= n; i++ {
			if _, ignore := ignoreMap[i]; !ignore {
				step++
			}
			if step == index {
				ignoreMap[i] = struct{}{}
				res = i
				break
			}
		}
		return strconv.Itoa(res)
	}
	jiecheng := make([]int, 0)
	jiecheng = append(jiecheng, 1)
	tmp := 1
	for i := 1; i <= 9; i++ {
		tmp *= i
		jiecheng = append(jiecheng, tmp)
	}
	for count := n; count > 0; count-- {
		jc := jiecheng[count-1]
		index := k / jc
		if k%jc != 0 {
			index++
		}
		k = k - (index-1)*jc
		res = append(res, getValByIndex(n, index))
	}
	return strings.Join(res, "")
}

func RotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	preNode := &ListNode{}
	node := head
	next := head.Next
	for {
		if node == nil {
			break
		}
		next = node.Next
		node.Next = preNode
		preNode = node
		node = next
	}
	var reverse func(node *ListNode) (*ListNode, *ListNode)
	reverse = func(node *ListNode) (*ListNode, *ListNode) {
		if node.Next != nil {
			start, end := reverse(node.Next)
			end.Next = node
			node.Next = nil
			return start, node
		} else {
			return node, node
		}
	}
	start, _ := reverse(head)
	return start
}
