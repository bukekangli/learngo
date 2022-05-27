package arithmetic_test

import (
	"github.com/bukekangli/learngo/arithmetic"
	"testing"
)

func TestMyPow(t *testing.T) {
	t.Logf("2*10=%f", arithmetic.MyPow(2, 10))
	t.Logf("2*-2=%f", arithmetic.MyPow(2, -2))
}

func TestMaxSubArray(t *testing.T) {
	t.Logf("%d", arithmetic.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	t.Logf("%d", arithmetic.MaxSubArray([]int{1}))
	t.Logf("%d", arithmetic.MaxSubArray([]int{5, 4, -1, 7, 8}))
}

func TestSpiralOrder(t *testing.T) {
	t.Logf("%v", arithmetic.SpiralOrder([][]int{{6, 9, 7}}))
	t.Logf("%v", arithmetic.SpiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	t.Logf("%v", arithmetic.SpiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

func TestCanJump(t *testing.T) {
	t.Logf("%t", arithmetic.CanJump([]int{2, 3, 1, 1, 4}))
	t.Logf("%t", arithmetic.CanJump([]int{3, 2, 1, 0, 4}))
}

func TestMerge(t *testing.T) {
	t.Logf("%v", arithmetic.Merge([][]int{{2, 6}, {1, 3}, {15, 18}, {8, 10}}))
	t.Logf("%v", arithmetic.Merge([][]int{{1, 4}, {4, 5}}))
}

func TestInsert(t *testing.T) {
	t.Logf("%v", arithmetic.Insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	t.Logf("%v", arithmetic.Insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{5, 7}))
	t.Logf("%v", arithmetic.Insert([][]int{}, []int{5, 7}))
	t.Logf("%v", arithmetic.Insert([][]int{{1, 5}}, []int{2, 3}))
	t.Logf("%v", arithmetic.Insert([][]int{{1, 5}}, []int{2, 7}))
}

func TestLengthOfLastWord(t *testing.T) {
	t.Logf("%v", arithmetic.LengthOfLastWord("Hello World"))
	t.Logf("%v", arithmetic.LengthOfLastWord("   fly me   to   the moon  "))
	t.Logf("%v", arithmetic.LengthOfLastWord("luffy is still joyboy"))
}

func TestGenerateMatrix(t *testing.T) {
	t.Logf("%v", arithmetic.GenerateMatrix(3))
	t.Logf("%v", arithmetic.GenerateMatrix(1))
}

func TestGetPermutation(t *testing.T) {
	t.Logf("%v", arithmetic.GetPermutation(3, 3))
	t.Logf("%v", arithmetic.GetPermutation(4, 9))
	t.Logf("%v", arithmetic.GetPermutation(3, 1))
}

func TestRotateRight(t *testing.T) {
	initL := func() *arithmetic.ListNode {

		return &arithmetic.ListNode{
			Val: 1,
			Next: &arithmetic.ListNode{
				Val: 2,
				Next: &arithmetic.ListNode{
					Val: 3,
					Next: &arithmetic.ListNode{
						Val: 4,
						Next: &arithmetic.ListNode{
							Val: 5,
						},
					},
				},
			},
		}
	}
	print := func(res *arithmetic.ListNode) {
		for {
			if res == nil {
				break
			}
			t.Logf("%d", res.Val)
			res = res.Next
		}
	}
	print(arithmetic.RotateRight(initL(), 1))
	t.Logf("-------")
	print(arithmetic.RotateRight(initL(), 2))
	t.Logf("-------")
	print(arithmetic.RotateRight(initL(), 3))
}
