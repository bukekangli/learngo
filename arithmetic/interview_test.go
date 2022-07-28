package arithmetic_test

import (
	"fmt"
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
	//print(arithmetic.RotateRight(initL(), 1))
	//t.Logf("-------")
	//print(arithmetic.RotateRight(initL(), 2))
	//t.Logf("-------")
	//print(arithmetic.RotateRight(initL(), 3))
	//print(arithmetic.ListNodeReverseV1(initL()))
	print(arithmetic.ListNodeReverseV2(initL()))
}

func TestUniquePaths(t *testing.T) {
	t.Logf("%v", arithmetic.UniquePaths(3, 7))
	t.Logf("%v", arithmetic.UniquePaths(3, 2))
	t.Logf("%v", arithmetic.UniquePaths(7, 3))
	t.Logf("%v", arithmetic.UniquePaths(3, 3))
}

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Logf("%v", arithmetic.UniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	t.Logf("%v", arithmetic.UniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}

func TestMinPathSum(t *testing.T) {
	t.Logf("%v", arithmetic.MinPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	t.Logf("%v", arithmetic.MinPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}

func TestPlusOne(t *testing.T) {
	t.Logf("%v", arithmetic.PlusOne([]int{1, 2, 3}))
	t.Logf("%v", arithmetic.PlusOne([]int{4, 3, 2, 1}))
	t.Logf("%v", arithmetic.PlusOne([]int{0}))
	t.Logf("%v", arithmetic.PlusOne([]int{9}))
}

func TestAddBinary(t *testing.T) {
	t.Logf("%v", arithmetic.AddBinary("11", "1"))
	t.Logf("%v", arithmetic.AddBinary("1010", "1011"))
	t.Logf("%v", arithmetic.AddBinary("0", "0"))
	t.Logf("%v", arithmetic.AddBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
}

func TestIsInterfaceIsNil(t *testing.T) {
	arithmetic.IsInterfaceIsNil()
}

func TestSolution(t *testing.T) {
	head := &arithmetic.Node{
		Data: 1,
		Next: &arithmetic.Node{
			Data: 2,
			Next: &arithmetic.Node{
				Data: 3,
				Next: &arithmetic.Node{
					Data: 4,
					Next: &arithmetic.Node{
						Data: 5,
					},
				},
			},
		},
	}
	ctx := arithmetic.Solution(head)
	print := func(head *arithmetic.Node) {
		for head != nil {
			fmt.Printf("%d -> ", head.Data)
			head = head.Next
		}
	}
	print(ctx.First)
	print(ctx.Second)
}

func TestSimplifyPath(t *testing.T) {
	t.Logf("%s", arithmetic.SimplifyPath("/home/"))
	t.Logf("%s", arithmetic.SimplifyPath("/../"))
	t.Logf("%s", arithmetic.SimplifyPath("/a/./b/../../c/"))
}

func TestInput(t *testing.T) {
	t.Logf("%v", arithmetic.Input([]int{2, 3, 1, 1}, 6))
}

func TestPrint(t *testing.T) {
	arithmetic.Print("1a2b")
}
