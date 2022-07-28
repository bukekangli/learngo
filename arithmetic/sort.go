package arithmetic

import (
	"fmt"
	"sort"
)

func BubbleSort(l []int) {
	// 冒泡排序
	// 原地排序算法，时间复杂度O(n*n)
	length := len(l)
	if length <= 1 {
		return
	}
	for i := length - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if l[j] > l[j+1] {
				l[j], l[j+1] = l[j+1], l[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

func InsertSort(l []int) {
	// 插入排序
	// 原地排序算法、时间复杂度为O(n*n)
	length := len(l)
	for i := 1; i < length; i++ {
		value := l[i]
		j := i - 1
		for ; j >= 0; j-- {
			if l[j] > value {
				l[j+1] = l[j]
			} else {
				break
			}
		}
		l[j+1] = value
	}
}

func SelectSort(l []int) {
	// 选择排序
	// 原地排序算法，时间复杂度O(n*n)
	length := len(l)
	for i := 0; i < length; i++ {
		minIndex := i
		for j := i; j < length; j++ {
			if l[j] < l[minIndex] {
				minIndex = j
			}
		}
		l[i], l[minIndex] = l[minIndex], l[i]
	}
}

func MergeSort(arr []int) {
	// 归并排序
	// 递推公式：merge_sort(0, n) = merge_sort(0, i) + merge_sort(i+1, n)
	// 这个时间复杂度是O(n*lgn)，空间复杂度为O(n*n)
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, head int, tail int) {
	if tail <= head {
		return
	}
	middle := (head + tail) / 2
	mergeSort(arr, head, middle)
	mergeSort(arr, middle+1, tail)
	merge(arr, head, middle, tail)
}

func merge(arr []int, head int, mid int, tail int) {
	var mergeSlice []int
	leftIndex, rightIndex := head, mid+1
	// 将左右两个有序数组重新排序
	for i := leftIndex; i < rightIndex; i++ {
		// 左边有序slice使用完，只拼接右侧数据
		if leftIndex > mid {
			mergeSlice = append(mergeSlice, arr[rightIndex])
			rightIndex++
			// 右侧有序slice使用完，只拼接左侧slice
		} else if rightIndex > tail {
			mergeSlice = append(mergeSlice, arr[leftIndex])
			// 把2组有序slice拼到临时slice中
		} else if arr[leftIndex] < arr[rightIndex] {
			mergeSlice = append(mergeSlice, arr[leftIndex])
			leftIndex++
		} else {
			mergeSlice = append(mergeSlice, arr[rightIndex])
			rightIndex++
		}
	}
	for index, val := range mergeSlice {
		arr[head+index] = val
	}
}

type quickS struct {
	nums []int
}

func newQuickS(nums []int) quickS {
	return quickS{nums: nums}
}
func (s *quickS) _quickSort(p, r int) {
	if p < r {
		q := s.partition(p, r)
		s._quickSort(p, q-1)
		s._quickSort(q+1, r)
	}
}
func (s *quickS) quickSort() {
	s._quickSort(0, len(s.nums)-1)
}

func (s *quickS) partition(p, r int) int {
	x := s.nums[r]
	i := p - 1
	for j := p; j <= r-1; j++ {
		if s.nums[j] <= x {
			i = i + 1
			s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
		}
	}
	s.nums[i+1], s.nums[r] = s.nums[r], s.nums[i+1]
	return i + 1
}

func QuickSort(nums []int) {
	q := newQuickS(nums)
	q.quickSort()
}

type heap struct {
	nums     []int
	heapSize int
}

func newHeap(nums []int) heap {
	return heap{nums: nums}
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) left(i int) int {
	return 2*i + 1
}

func (h *heap) right(i int) int {
	return 2*i + 2
}

func (h *heap) maxHeapify(i int) {
	largest := i
	l := h.left(i)
	r := h.right(i)
	if l < h.heapSize && h.nums[l] > h.nums[i] {
		largest = l
	}
	if r < h.heapSize && h.nums[r] > h.nums[largest] {
		largest = r
	}
	if largest != i {
		h.nums[i], h.nums[largest] = h.nums[largest], h.nums[i]
		h.maxHeapify(largest)
	}
}

func (h *heap) buildMaxHeap() {
	h.heapSize = len(h.nums)
	for i := len(h.nums) / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
}

func (h *heap) heapSort() {
	h.buildMaxHeap()
	fmt.Println(h.nums)
	for i := len(h.nums) - 1; i >= 1; i-- {
		h.nums[0], h.nums[i] = h.nums[i], h.nums[0]
		h.heapSize--
		h.maxHeapify(0)
	}
}

func HeapSort(nums []int) {
	h := newHeap(nums)
	h.heapSort()
}

func Unique(s []int) uint {
	sort.Ints(s)
	firstIndex, secondIndex := 0, 1
	for firstIndex < secondIndex && secondIndex < len(s) {
		for s[firstIndex] == s[secondIndex] {
			secondIndex++
		}
		if s[firstIndex+1] < s[secondIndex] {
			firstIndex++
			s[firstIndex], s[secondIndex] = s[secondIndex], s[firstIndex]
		} else {
			secondIndex++
		}
	}
	return uint(firstIndex + 1)
}
