package crackingcodinginterview

import (
	"fmt"
	"sort"
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	if SprintNode(removeDuplicateNodes(root)) != "[1 2 3]" {
		t.Error(`TestRemoveDuplicateNodes failed`)
	}
}

func TestKthToLast(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	if kthToLast(root, 1) != 4 || kthToLast(root, 2) != 3 || kthToLast(root, 3) != 2 || kthToLast(root, 4) != 1 {
		t.Error(`TestKthToLast failed`)
	}
}

func TestDeleteNode(t *testing.T) {
	root := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	root.Next = second
	third := &ListNode{Val: 3}
	root.Next.Next = third
	root.Next.Next.Next = &ListNode{Val: 4}
	deleteNode(second)
	if SprintNode(root) != "[1 3 4]" {
		t.Error(`TestDeleteNode failed`)
	}

	deleteNode(root.Next)
	if SprintNode(root) != "[1 4]" {
		t.Error(`TestDeleteNode failed`, SprintNode(root))
	}
}

func TestPartition(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 4}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	// 保证升序输出
	ret := Node2Array(partition(root, 3))
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	if fmt.Sprint(ret) != "[1 2 2 3 4 5]" {
		t.Error(`TestPartition failed`)
	}

	root = &ListNode{Val: 2}
	root.Next = &ListNode{Val: 1}
	if SprintNode(partition(root, 2)) != "[1 2]" {
		t.Error(`TestPartition failed`)
	}
}

func TestReverseList(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 4}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	if SprintNode(ReverseList(root)) != "[2 5 2 3 4 1]" {
		t.Error(`TestReverseList failed`)
	}

	root = &ListNode{Val: 2}
	root.Next = &ListNode{Val: 1}
	if SprintNode(ReverseList(root)) != "[1 2]" {
		t.Error(`TestReverseList failed`)
	}
}

func TestIsPalindrome(t *testing.T) {
	root := &ListNode{Val: 1}
	if !isPalindrome(root) {
		t.Error(`TestIsPalindrome failed`)
	}
	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	if isPalindrome(root) {
		t.Error(`TestIsPalindrome failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 1}
	if !isPalindrome(root) {
		t.Error(`TestIsPalindrome failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 1}
	root.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next = &ListNode{Val: 1}
	if !isPalindrome(root) {
		t.Error(`TestIsPalindrome failed`)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 7}
	l1.Next = &ListNode{Val: 1}
	l1.Next.Next = &ListNode{Val: 6}
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 2}
	if SprintNode(addTwoNumbers(l1, l2)) != "[2 1 9]" {
		t.Error(`TestAddTwoNumbers failed`)
	}

	l1 = &ListNode{Val: 7}
	l1.Next = &ListNode{Val: 1}
	l1.Next.Next = &ListNode{Val: 6}
	l2 = &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 3}
	if SprintNode(addTwoNumbers(l1, l2)) != "[2 1 0 1]" {
		t.Error(`TestAddTwoNumbers failed`)
	}
}

func TestGetIntersectionNode(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 4}
	l3 := &ListNode{Val: 2}
	l3.Next = &ListNode{Val: 3}
	l1.Next = l3
	l2.Next = l3
	if getIntersectionNode(l1, l2) != l3 {
		t.Error(`TestGetIntersectionNode failed`)
	}

	l1 = &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l2 = &ListNode{Val: 1}
	l3.Next = &ListNode{Val: 2}
	if getIntersectionNode(l1, l2) != nil {
		t.Error(`TestGetIntersectionNode failed`)
	}
}

func TestDetectCycle(t *testing.T) {
	root := &ListNode{Val: 1}
	loops := &ListNode{Val: 2}
	root.Next = loops
	loops.Next = &ListNode{Val: 3}
	loops.Next.Next = &ListNode{Val: 4}
	loops.Next.Next.Next = loops
	if detectCycle(root) != loops {
		t.Error(`TestDetectCycle failed`)
	}

	root = &ListNode{Val: 1}
	loops = &ListNode{Val: 2}
	root.Next = loops
	loops.Next = &ListNode{Val: 3}
	loops.Next.Next = loops
	if detectCycle(root) != loops {
		t.Error(`TestDetectCycle failed`)
	}
}
