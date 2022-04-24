package arithmetic_test

import (
	"github.com/bukekangli/learngo/arithmetic"
	"testing"
)

func TestInsertSort(t *testing.T) {
	nums := []int{3, 2, 1}
	arithmetic.InsertSort(nums)
	t.Logf("nums: %v", nums)
}

func TestHeapSort(t *testing.T) {
	nums := []int{16, 14, 10, 8, 7, 9, 2, 4, 1}
	arithmetic.HeapSort(nums)
	t.Logf("nums: %v", nums)
}

func TestQuickSort(t *testing.T) {
	nums := []int{16, 14, 10, 8, 7, 9, 2, 4, 1}
	arithmetic.QuickSort(nums)
	t.Logf("nums: %v", nums)
}
