package nowcoder

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	operators := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{2, 1},
		{1, 4, 4},
		{2, 2},
		{1, 3, 5},
		{2, 3},
	}
	result := LRU(operators, 3)
	target := "[1 -1 5]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}

	operators = [][]int{{1, 1, 1}, {1, 2, 2}, {2, 1}, {1, 3, 3}, {2, 2}, {1, 4, 4}, {2, 1}, {2, 3}, {2, 4}}
	result = LRU(operators, 2)
	target = "[1 -1 -1 3 4]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}
}

func TestThreeOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	target := "[[1 2 3] [2 1 3] [2 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}

	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}

	right.Right = &TreeNode{Val: 6}
	target = "[[1 2 4 5 7 8 3 6] [4 2 7 5 8 1 3 6] [4 7 8 5 2 6 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right
	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}
	right.Right = &TreeNode{Val: 6}

	target := "[[1] [2 3] [4 5 6] [7 8]]"
	if fmt.Sprint(levelOrder(root)) != target {
		t.Error(`TestLevelOrder failed`)
	}
}

func TestGetLeastNumbersQuickSort(t *testing.T) {
	least4 := GetLeastNumbers_QuickSort([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersQuickSort failed`)
	}

	least3 := GetLeastNumbers_QuickSort([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersQuickSort failed`)
	}
}

func TestGetLeastNumbersSelfSort(t *testing.T) {
	least4 := GetLeastNumbers_SelfSort([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersSelfSort failed`)
	}

	least3 := GetLeastNumbers_SelfSort([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersSelfSort failed`)
	}
}

func TestGetLeastNumbersHeap(t *testing.T) {
	least4 := GetLeastNumbers_Heap([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[4 3 1 2]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersHeap failed`, fmt.Sprint(least4))
	}

	least3 := GetLeastNumbers_Heap([]int{0, 1, 2, 1, 2}, 3)
	target = "[1 0 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersHeap failed`, fmt.Sprint(least3))
	}
}

func TestGetLeastNumbersQuickSearchp(t *testing.T) {
	least4 := GetLeastNumbers_QuickSearch([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	target := "[1 2 3 4]"
	if fmt.Sprint(least4) != target {
		t.Error(`TestGetLeastNumbersQuickSearchp failed`, fmt.Sprint(least4))
	}

	least3 := GetLeastNumbers_QuickSearch([]int{0, 1, 2, 1, 2}, 3)
	target = "[0 1 1]"
	if fmt.Sprint(least3) != target {
		t.Error(`TestGetLeastNumbersQuickSearchp failed`, fmt.Sprint(least3))
	}
}
