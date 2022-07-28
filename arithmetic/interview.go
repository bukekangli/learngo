package arithmetic

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
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
	iter := head
	nodeLen := 1
	for iter.Next != nil {
		iter = iter.Next
		nodeLen++
	}
	iter.Next = head
	preNode := iter
	iter = iter.Next
	for k = nodeLen - k%nodeLen; k > 0; k-- {
		preNode = iter
		iter = iter.Next
	}
	preNode.Next = nil
	return iter
}

// ListNodeReverseV1 插入法实现的链表反转
func ListNodeReverseV1(head *ListNode) *ListNode {
	node := &ListNode{}
	node.Next = head
	head = head.Next
	node.Next.Next = nil

	for head != nil {
		tmp := head.Next
		head.Next = node.Next
		node.Next = head
		head = tmp
	}
	return node.Next
}

// ListNodeReverseV2 递归法实现的链表反转
func ListNodeReverseV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ListNodeReverseV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func UniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	res := make([][]int, m)
	for i, _ := range res {
		res[i] = make([]int, n)
		res[i][0] = 1
	}
	for j := 0; j < n; j++ {
		res[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
		}
	}
	return res[m-1][n-1]
}

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 {
		for i := 0; i < len(obstacleGrid[0]); i++ {
			if obstacleGrid[0][i] == 1 {
				return 0
			}
		}
		return 1
	}
	if len(obstacleGrid[0]) == 1 {
		for i := 0; i < len(obstacleGrid); i++ {
			if obstacleGrid[i][0] == 1 {
				return 0
			}
		}
		return 1
	}
	if len(obstacleGrid) == 1 || len(obstacleGrid[0]) == 1 {
		return 0
	}
	for i, _ := range obstacleGrid {
		if obstacleGrid[i][0] == 1 {
			break
		}
		obstacleGrid[i][0] = -1
	}
	for j := 0; j < len(obstacleGrid[0]); j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		obstacleGrid[0][j] = -1
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			val1 := obstacleGrid[i-1][j]
			if val1 == 1 {
				val1 = 0
			}
			val2 := obstacleGrid[i][j-1]
			if val2 == 1 {
				val2 = 0
			}
			if obstacleGrid[i][j] == 1 {
				continue
			}
			obstacleGrid[i][j] = val1 + val2
		}
	}
	val := obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
	if val == 1 {
		return 0
	}
	return -val
}

func MinPathSum(grid [][]int) int {
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			grid[i][j] = grid[i][j] + min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func PlusOne(digits []int) []int {
	flag := 0
	for i := len(digits) - 1; i >= 0; i-- {
		var val int
		if i == len(digits)-1 {
			val = digits[i] + 1
		} else {
			val = digits[i] + flag
			flag = 0
		}
		if val >= 10 {
			val -= 10
			flag = 1
		}
		digits[i] = val
	}
	if flag != 0 {
		tmp := make([]int, 1)
		tmp[0] = 1
		digits = append(tmp, digits...)
	}
	return digits
}

func AddBinary(a string, b string) string {
	lenA, lenB := len(a), len(b)
	longest := lenA
	if lenB > lenA {
		longest = lenB
	}
	var flag byte
	r := ""
	res := ""
	for index := 0; index < longest; index++ {
		var val1, val2 byte
		if index < lenA {
			val1 = a[lenA-1-index]
		}
		if index < lenB {
			val2 = b[lenB-1-index]
		}
		r, flag = sum(val1, val2, flag)
		res = r + res
	}
	if flag == '1' {
		res = "1" + res
	}
	return res
}

func sum(a, b, c byte) (string, byte) {
	count := 0
	if a == '1' {
		count++
	}
	if b == '1' {
		count++
	}
	if c == '1' {
		count++
	}
	switch count {
	case 3:
		return "1", '1'
	case 2:
		return "0", '1'
	case 1:
		return "1", '0'
	default:
		return "0", '0'
	}
}

func IsInterfaceIsNil() {
	var writer io.Writer
	fmt.Printf("writer is nil => %t\n", writer == nil)
	var bufWriter *bufio.Writer
	fmt.Printf("bufWriter is nil => %t\n", bufWriter == nil)
	bufWriter2 := func() io.Writer {
		var w *bufio.Writer
		fmt.Printf("w is nil => %t\n", w == nil)
		if w == nil {
			return nil
		}
		return w
	}()
	// 因为bufWriter2会比较类型和值是不是都为nil，这时它已经有类型了
	fmt.Printf("bufWriter is nil => %t\n", bufWriter2 == nil)
	if bufWriter2 != nil {
		bufWriter2.Write([]byte("golang"))
	} else {
		fmt.Println("bufWriter is nil")
	}
	bufWriter3 := func() io.Writer {
		var w *bufio.Writer
		fmt.Printf("w is nil => %t\n", w == nil)
		return w
	}()
	fmt.Printf("bufWriter is nil => %t\n", bufWriter3 == nil)
	fmt.Printf("IsNil => bufWriter is nil => %t\n", reflect.ValueOf(bufWriter3).IsNil())

	fmt.Printf("IsNil => bufWriter is nil => %T\n", bufWriter)
	fmt.Printf("IsNil => bufWriter is nil => %s\n", reflect.TypeOf(bufWriter).Kind())
}

type Node struct {
	Data int
	Next *Node
}

type CTX struct {
	First  *Node
	Second *Node
}

func Solution(head *Node) *CTX {
	ctx := &CTX{}
	first, second := &Node{}, &Node{}
	var preFirst, preSecond *Node
	flag := 0
	ctx.First, ctx.Second = first, second
	for head != nil {
		if flag == 0 {
			flag = 1
			first.Data = head.Data
			first.Next = &Node{}
			preFirst = first
			first = first.Next
		} else {
			flag = 0
			second.Data = head.Data
			second.Next = &Node{}
			preSecond = second
			second = second.Next
		}
		head = head.Next
	}
	preFirst.Next = nil
	preSecond.Next = nil
	return ctx
}

func SimplifyPath(path string) string {
	stack := make([]string, 0)
	for _, part := range strings.Split(path, "/") {
		switch part {
		case ".", "":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	return "/" + strings.Join(stack, "/")
}

// 2 3 1 1    6    out: 231 231
func Input(list []int, target int) [][]int {
	sort.Ints(list)
	res := make([][]int, 0)
	for i := 0; i < len(list); i++ {
		head, tail := i+1, len(list)-1
		for head < tail {
			_sum := list[head] + list[tail] + list[i]
			if _sum == target {
				res = append(res, []int{list[i], list[head], list[tail]})
				head++
			} else if _sum > target {
				tail--
			} else {
				head++
			}
		}
	}
	return res
}

func Print(s string) {
	var wg sync.WaitGroup
	c1 := make(chan struct{}, 1)
	c2 := make(chan struct{}, 1)
	buff := make(chan string, 1)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			<-c1
			if val, ok := <-buff; ok {
				fmt.Println(val)
				c2 <- struct{}{}
			} else {
				break
			}
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		for {
			<-c2
			if val, ok := <-buff; ok {
				fmt.Println(val)
				c1 <- struct{}{}
			} else {
				break
			}
		}
	}()
	c1 <- struct{}{}
	for _, c := range s {
		buff <- string(c)
	}
	wg.Wait()
}
