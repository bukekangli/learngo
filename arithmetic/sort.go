package arithmetic

// 冒泡排序
func BubbleSort(l []int) {
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

// 插入排序
func InsertSort(l []int) {
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

// 选择排序
func SelectSort(l []int) {
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

// 归并排序
// 递推公式：merge_sort(0, n) = merge_sort(0, i) + merge_sort(i+1, n)
// 我这个版本不是原地排序还比较复杂

func MyMergeSort(l []int) []int {
	var mergeList []int
	length := len(l)
	if length <= 1 {
		return l
	}
	i := len(l) / 2
	l1 := MyMergeSort(l[:i])
	l2 := MyMergeSort(l[i:])
	l1Length, l2Length := len(l1), len(l2)
	l1Index, l2Index := 0, 0
	for {
		if l1[l1Index] < l2[l2Index] {
			mergeList = append(mergeList, l1[l1Index])
			l1Index++
		} else {
			mergeList = append(mergeList, l2[l2Index])
			l2Index++
		}
		if l1Index >= l1Length {
			for _, val := range l2[l2Index:] {
				mergeList = append(mergeList, val)
			}
			break
		} else if l2Index >= l2Length {
			for _, val := range l1[l1Index:] {
				mergeList = append(mergeList, val)
			}
			break
		}
	}
	return mergeList
}

func MergeSort(arr []int) {
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
