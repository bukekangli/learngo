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
