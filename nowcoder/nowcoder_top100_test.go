package nowcoder

import (
	"fmt"
	"testing"
)

func TestIsPail(t *testing.T) {
	head := &ListNode{Val: 1}
	if !isPail(head) || !isPail(nil) {
		t.Error(`TestIsPail failed`)
	}

	head.Next = &ListNode{Val: 2}
	if isPail(head) {
		t.Error(`TestIsPail failed`)
	}

	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	if !isPail(head) {
		t.Error(`TestIsPail failed`)
	}
}

func TestSumNumbers(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	if sumNumbers(tree) != 12 {
		t.Error(`TestSumNumbers failed`)
	}
	tree.Right = &TreeNode{Val: 3}
	if sumNumbers(tree) != 25 {
		t.Error(`TestSumNumbers failed`)
	}
}

func TestPathSum(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	if fmt.Sprint(pathSum(tree, 1)) != "[]" {
		t.Error(`TestSumNumbers failed`)
	}
	if fmt.Sprint(pathSum(tree, 3)) != "[[1 2]]" {
		t.Error(`TestSumNumbers failed`)
	}
	tree.Right = &TreeNode{Val: 1}
	tree.Right.Left = &TreeNode{Val: 1}
	if fmt.Sprint(pathSum(tree, 3)) != "[[1 2] [1 1 1]]" {
		t.Error(`TestSumNumbers failed`)
	}
}

func TestReverseBetween(t *testing.T) {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	if SprintNode(reverseBetween(list, 2, 3)) != "[1 3 2]" {
		t.Error(`TestReverseBetween failed`)
	}

	list = &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	if SprintNode(reverseBetween(list, 2, 4)) != "[1 4 3 2 5]" {
		t.Error(`TestReverseBetween failed`)
	}
}

func TestReverseInt(t *testing.T) {
	if reverseInt(12) != 21 || reverseInt(-123) != -321 || reverseInt(10) != 1 || reverseInt(1147483649) != 0 {
		t.Error(`TestReverseInt failed`, reverseInt(-123))
	}
}

func TestUniquePaths(t *testing.T) {
	if uniquePaths(1, 1) != 1 || uniquePaths(2, 1) != 1 || uniquePaths(2, 2) != 2 || uniquePaths(4, 4) != 20 {
		t.Error(`TestUniquePaths failed`)
	}
}

func TestMergeInterval(t *testing.T) {
	// [[10,30],[20,60],[80,100],[150,180]]
	intervals := []*Interval{
		{20, 60},
		{10, 30},
		{80, 100},
		{150, 180},
	}

	merged := mergeInterval(intervals)
	if SPrintInterval(merged) != "[[10 60] [80 100] [150 180]]" {
		t.Error(`TestMergeInterval failed`)
	}

	// [[0,10],[10,20]]
	intervals = []*Interval{
		{0, 10},
		{10, 20},
	}
	merged = mergeInterval(intervals)
	if SPrintInterval(merged) != "[[0 20]]" {
		t.Error(`TestMergeInterval failed`)
	}
}

func TestFindMedianinTwoSortedAray(t *testing.T) {
	if findMedianinTwoSortedAray([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}) != 3 {
		t.Error(`TestFindMedianinTwoSortedAray failed`)
	}

	if findMedianinTwoSortedAray([]int{0, 1, 2}, []int{3, 4, 5}) != 2 {
		t.Error(`TestFindMedianinTwoSortedAray failed`)
	}
}

func TestGetMedian(t *testing.T) {
	// [5,2,3,4,1,6,7,0,8]
	Insert(5)
	fmt.Println(GetMedian())
	Insert(2)
	fmt.Println(GetMedian())
	Insert(3)
	fmt.Println(GetMedian())
	Insert(4)
	fmt.Println(GetMedian())
	Insert(1)
	fmt.Println(GetMedian())
	Insert(6)
	fmt.Println(GetMedian())
	Insert(7)
	fmt.Println(GetMedian())
	Insert(0)
	fmt.Println(GetMedian())
	Insert(8)
	fmt.Println(GetMedian())
}

func TestFindElement(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	if fmt.Sprint(findElement(mat, 2, 3, 6)) != "[1 2]" {
		t.Error(`TestFindElement failed`)
	}
	mat = [][]int{
		{1, 2, 3},
	}
	if fmt.Sprint(findElement(mat, 1, 3, 2)) != "[0 1]" {
		t.Error(`TestFindElement failed`)
	}
	mat = [][]int{
		{1, 4, 8},
		{2, 5, 9},
	}
	if fmt.Sprint(findElement(mat, 2, 3, 5)) != "[1 1]" {
		t.Error(`TestFindElement failed`)
	}
}

func TestLCSPlus1(t *testing.T) {
	if LCSPlus1("1A2C3D4B56", "B1D23A456A") != "123456" || LCSPlus1("abc", "def") != "-1" || LCSPlus1("abc", "abc") != "abc" || LCSPlus1("abc", "") != "-1" {
		t.Error(`TestLCSPlus1 failed`)
	}
}

func TestJudgeIt(t *testing.T) {
	if fmt.Sprint(judgeIt(nil)) != "[true true]" {
		t.Error("TestJudgeIt failed")
	}

	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	if fmt.Sprint(judgeIt(root)) != "[true true]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	if fmt.Sprint(judgeIt(root)) != "[true false]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	if fmt.Sprint(judgeIt(root)) != "[false true]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 8}
	if fmt.Sprint(judgeIt(root)) != "[true false]" {
		t.Error("TestJudgeIt failed", fmt.Sprint(judgeIt(root)))
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	if fmt.Sprint(judgeIt(root)) != "[false false]" {
		t.Error("TestJudgeIt failed", fmt.Sprint(judgeIt(root)))
	}
}

func TestDeleteDuplicates(t *testing.T) {
	root := &ListNode{Val: 1}
	if SprintNode(deleteDuplicates(root)) != "[1]" {
		t.Error(`TestDeleteDuplicates failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicates(root)) != "[2]" {
		t.Error(`TestDeleteDuplicates failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicates(root)) != "[]" {
		t.Error(`TestDeleteDuplicates failed`)
	}
}

func TestDeleteDuplicatesEasy(t *testing.T) {
	root := &ListNode{Val: 1}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1 2]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1 2]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}
}

func TestMinNumberDisappeared(t *testing.T) {
	if minNumberDisappeared([]int{1, 0, 2}) != 3 {
		t.Error(`TestMinNumberDisappeared failed`)
	}

	if minNumberDisappeared([]int{-2, 3, 4, 1, 5}) != 2 {
		t.Error(`TestMinNumberDisappeared failed`)
	}

	if minNumberDisappeared([]int{4, 5, 6, 8, 9}) != 1 {
		t.Error(`TestMinNumberDisappeared failed`)
	}
}

func TestAtoi(t *testing.T) {
	if atoi("  010") != 10 || atoi("+12") != 12 || atoi("-12ab") != -12 || atoi("123") != 123 || atoi("123e123") != 123 || atoi("e123") != 0 {
		t.Error(`TestAtoi failed`)
	}
}

func TestSolveLostNumber(t *testing.T) {
	if solveLostNumber([]int{0, 1, 2, 3, 4, 5, 7}) != 6 {
		t.Error(`TestSolveLostNumber failed`)
	}
	if solveLostNumber([]int{0, 2, 3}) != 1 {
		t.Error(`TestSolveLostNumber failed`)
	}
	if solveLostNumber([]int{0, 1, 2, 3, 4}) != 5 {
		t.Error(`TestSolveLostNumber failed`)
	}
}

func TestOddEvenList(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	if SprintNode(oddEvenList(root)) != "[1 3 5 2 4 6]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 4}
	root.Next.Next = &ListNode{Val: 6}
	root.Next.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next.Next = &ListNode{Val: 7}
	if SprintNode(oddEvenList(root)) != "[1 6 7 4 3]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}
}

func TestMaxPathSum(t *testing.T) {
	root := &TreeNode{Val: -2}
	root.Left = &TreeNode{Val: 1}
	if maxPathSum(root) != 1 {
		t.Error(`TestMaxPathSum failed`)
	}

	root = &TreeNode{Val: -2}
	root.Right = &TreeNode{Val: -3}
	if maxPathSum(root) != -2 {
		t.Error(`TestMaxPathSum failed`)
	}

	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	if maxPathSum(root) != 3 {
		t.Error(`TestMaxPathSum failed`)
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	if maxPathSum(root) != 18 {
		t.Error(`TestMaxPathSum failed`)
	}
}

func TestIsSymmetric(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	if isSymmetric(root) != true {
		t.Error("TestIsSymmetric failed")
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}
	if isSymmetric(root) != false {
		t.Error("TestIsSymmetric failed")
	}
}

func TestIsSymmetricRecursion(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	if isSymmetricRecursion(root) != true {
		t.Error("TestIsSymmetricRecursion failed")
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}
	if isSymmetricRecursion(root) != false {
		t.Error("TestIsSymmetricRecursion failed")
	}
}

func TestGenerateParenthesis(t *testing.T) {
	if fmt.Sprint(generateParenthesis(1)) != "[()]" {
		t.Error(`TestGenerateParenthesis failed`)
	}

	if fmt.Sprint(generateParenthesis(2)) != "[(()) ()()]" {
		t.Error(`TestGenerateParenthesis failed`)
	}

	if fmt.Sprint(generateParenthesis(3)) != "[((())) (()()) (())() ()(()) ()()()]" {
		t.Error(`TestGenerateParenthesis failed`)
	}
}

func TestRotateMatrix(t *testing.T) {
	if fmt.Sprint(rotateMatrix([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}, 3)) != "[[7 4 1] [8 5 2] [9 6 3]]" {
		t.Error(`TestRotateMatrix failed`)
	}

	if fmt.Sprint(rotateMatrix([][]int{
		{1, 2},
		{3, 4},
	}, 2)) != "[[3 1] [4 2]]" {
		t.Error(`TestRotateMatrix failed`)
	}

	if fmt.Sprint(rotateMatrix([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}, 4)) != "[[3 1] [4 2]]" {
		t.Error(`TestRotateMatrix failed`)
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	if fmt.Sprint(restoreIpAddresses("25525522135")) != "[255.255.22.135 255.255.221.35]" {
		t.Error(`TestRestoreIpAddresses failed`)
	}

	if fmt.Sprint(restoreIpAddresses("1111")) != "[1.1.1.1]" {
		t.Error(`TestRestoreIpAddresses failed`)
	}

	if fmt.Sprint(restoreIpAddresses("000256")) != "[]" {
		t.Error(`TestRestoreIpAddresses failed`)
	}
}

func TestKnapsack(t *testing.T) {
	if knapsack(10, 2, [][]int{{1, 3}, {10, 4}}) != 4 {
		t.Error(`TestKnapsack failed`)
	}

	if knapsack(38, 50, [][]int{{43, 50}, {50, 45}, {31, 10}, {3, 32}, {2, 9}, {47, 15}, {27, 46}, {24, 45}, {36, 48}, {38, 29}, {25, 23}, {36, 48}, {10, 7}, {1, 5}, {22, 3}, {12, 13}, {35, 9}, {33, 12}, {3, 15}, {50, 30}, {18, 48}, {20, 19}, {34, 24}, {21, 10}, {25, 9}, {10, 21}, {27, 35}, {12, 27}, {35, 39}, {7, 45}, {25, 46}, {18, 23}, {33, 28}, {22, 50}, {17, 48}, {3, 44}, {34, 13}, {41, 2}, {34, 6}, {47, 46}, {48, 30}, {31, 7}, {34, 32}, {40, 50}, {31, 39}, {10, 41}, {22, 36}, {14, 48}, {46, 17}, {8, 33}}) != 224 {
		t.Error(`TestKnapsack failed`)
	}
}

func TestMinEditCost(t *testing.T) {
	if minEditCost("abc", "adc", 5, 3, 2) != 2 {
		t.Error(`TestMinEditCost failed`)
	}
	if minEditCost("abc", "adc", 5, 3, 100) != 8 {
		t.Error(`TestMinEditCost failed`)
	}
}

func TestLFU(t *testing.T) {
	if fmt.Sprint(LFU([][]int{{1, 1, 1}, {1, 2, 2}, {1, 3, 2}, {1, 2, 4}, {1, 3, 5}, {2, 2}, {1, 4, 4}, {2, 1}}, 3)) != "[4 -1]" {
		t.Error(`TestLFU failed`)
	}
}

func TestRemoveKey(t *testing.T) {
	data := []int{1, 2, 3}
	if fmt.Sprint(removeKey(data, 1)) != "[2 3]" {
		t.Error(`TestRemoveKey failed`)
	}
}

func TestReorderList(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}

	if SprintNode(reorderList(root)) != "[1 3 2]" {
		t.Error(`TestReorderList failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	if SprintNode(reorderList(root)) != "[1 4 2 3]" {
		t.Error(`TestReorderList failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	if SprintNode(reorderList(root)) != "[1 5 2 4 3]" {
		t.Error(`TestReorderList failed`)
	}
}

func TestPermuteUnique(t *testing.T) {
	data := []int{1, 1, 2}
	if fmt.Sprint(permuteUnique(data)) != "[[1 1 2] [1 2 1] [2 1 1]]" {
		t.Error(`TestPermuteUnique failed`)
	}

	data = []int{0, -1}
	if fmt.Sprint(permuteUnique(data)) != "[[-1 0] [0 -1]]" {
		t.Error(`TestPermuteUnique failed`)
	}
}

func TestMaxInWindows(t *testing.T) {
	if fmt.Sprint(maxInWindows([]int{2, 3, 4, 2, 6, 2, 5, 1}, 3)) != "[4 4 6 6 6 5]" {
		t.Error(`TestMaxInWindows failed`)
	}

	if fmt.Sprint(maxInWindows([]int{9, 10, 9, -7, -3, 8, 2, -6}, 5)) != "[10 10 9 8]" {
		t.Error(`TestMaxInWindows failed`)
	}

	if fmt.Sprint(maxInWindows([]int{1, 2, 3, 4}, 5)) != "[]" {
		t.Error(`TestMaxInWindows failed`)
	}
}

func TestHasPathSum(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	if hasPathSum(root, 0) != false || hasPathSum(root, 3) != true {
		t.Error(`TestHasPathSum failed`)
	}
}

func TestCombinationSum2(t *testing.T) {
	if fmt.Sprint(combinationSum2([]int{2, 1}, 0)) != "[]" {
		t.Error(`TestCombinationSum2 failed`)
	}
	if fmt.Sprint(combinationSum2([]int{100, 10, 20, 70, 60, 10, 50}, 80)) != "[[10 10 60] [10 20 50] [10 70] [20 60]]" {
		t.Error(`TestCombinationSum2 failed`)
	}

	if fmt.Sprint(combinationSum2([]int{22, 49, 5, 24, 26}, 77)) != "[[5 22 24 26]]" {
		t.Error(`TestCombinationSum2 failed`)
	}
}

func TestLongestValidParentheses(t *testing.T) {
	if longestValidParentheses("(())") != 4 {
		t.Error(`TestLongestValidParentheses failed`)
	}

	if longestValidParentheses("(()") != 2 {
		t.Error(`TestLongestValidParentheses failed`)
	}

	if longestValidParentheses(")()(()()((((((())(") != 4 {
		t.Error(`TestLongestValidParentheses failed`)
	}

	if longestValidParentheses(")()())") != 4 {
		t.Error(`TestLongestValidParentheses failed`)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	if longestCommonPrefix([]string{}) != "" {
		t.Error(`TestLongestCommonPrefix failed`)
	}
	if longestCommonPrefix([]string{"abc"}) != "abc" {
		t.Error(`TestLongestCommonPrefix failed`)
	}
	if longestCommonPrefix([]string{"abca", "abc", "abca", "abc", "abcc"}) != "abc" {
		t.Error(`TestLongestCommonPrefix failed`)
	}
}

func TestIsPalindrome(t *testing.T) {
	if isPalindrome(-1) != false || isPalindrome(121) != true || isPalindrome(122) != false {
		t.Error(`TestIsPalindrome failed`, isPalindrome(-1))
	}
}

func TestBinarySearchWithDuplicate(t *testing.T) {
	if binarySearchWithDuplicate([]int{1, 2, 4, 4, 5}, 3) != -1 || binarySearchWithDuplicate([]int{1, 1, 1, 1, 1}, 1) != 0 {
		t.Error(`TestBinarySearchWithDuplicate failed`)
	}
}

func TestSolvePieces(t *testing.T) {
	if solvePieces(10, 2) != 4 || solvePieces(105, 2) != 14 {
		t.Error(`TestSolvePieces failed`)
	}
}

func TestSerialize(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	if Serialize(root) != "1,2,3,#,#,6,7,#,#,#,#" {
		t.Error(`TestSerialize failed`)
	}
}

func TestDeserialize(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	PrintTreeNode(Deserialize(Serialize(root)))
}

func TestKthNode(t *testing.T) {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	if KthNode(root, 3).Val != 3 || KthNode(root, 4).Val != 6 || KthNode(root, 6) != nil {
		t.Error(`TestKthNode failed`)
	}
}

func TestTran(t *testing.T) {
	if trans("This is a sample", 16) != "SAMPLE A IS tHIS" {
		t.Error(`TestTran failed`)
	}

	if trans("iOS", 3) != "Ios" || trans("nowcoder", 8) != "NOWCODER" {
		t.Error(`TestTran failed`)
	}
}

func TestMLS(t *testing.T) {
	if MLS([]int{7, 8, 2, 6, 1, 1, 7, 2, 4, 3, 4, 7, 5, 6, 8, 2}) != 8 {
		t.Error(`TestMLS failed`)
	}
	if MLS([]int{100, 4, 200, 1, 3, 2}) != 4 || MLS([]int{1, 1, 1}) != 1 {
		t.Error(`TestMLS failed`)
	}
}

func TestYsf(t *testing.T) {
	if ysf(5, 2) != 3 || ysf(1, 1) != 1 {
		t.Error(`TestYsf failed`)
	}
}

func TestSolveMaxValue(t *testing.T) {
	if solveMaxValue([]int{2, 20, 23, 4, 8}) != "8423220" || solveMaxValue([]int{30, 1}) != "301" {
		t.Error(`TestSolveMaxValue failed`)
	}
}

func TestValidIP(t *testing.T) {
	fmt.Println(validIP("172.16.254.1"))
	fmt.Println(validIP("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	fmt.Println(validIP("256.256.256.256"))
}

func TestFindKthToTail(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	if SprintNode(FindKthToTail(head, 3)) != "[3 4 5]" {
		t.Error(`TestFindKthToTail failed`)
	}
}

func TestFindValue(t *testing.T) {
	root := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}

	if Find(7, root) != true || Find(3, root) != false {
		t.Error(`TestFind failed`)
	}
}

func TestSolveMultiple(t *testing.T) {
	if solveMultiple("99", "11") != "1089" || solveMultiple("999", "111") != "110889" || solveMultiple("733064366", "459309139") != "336703162779040874" {
		t.Error(`TestSolveMultiple failed`)
	}
}

func TestSubsets(t *testing.T) {
	if fmt.Sprint(subsets([]int{1, 2, 3})) != "[[] [1] [2] [3] [1 2] [1 3] [2 3] [1 2 3]]" {
		t.Error(`TestSubsets failed`)
	}

	if fmt.Sprint(subsets([]int{1, 1, 2})) != "[[] [1] [2] [1 1] [1 2] [1 1 2]]" {
		t.Error(`TestSubsets failed`)
	}
}

func TestDfsPermute(t *testing.T) {
	if fmt.Sprint(permute([]int{2, 1, 3})) != "[[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]" {
		t.Error(`TestDfsPermute failed`)
	}
}

func TestFindPeakElement(t *testing.T) {
	if findPeakElement([]int{1, 2, 3, 4, 5}) != 4 {
		t.Error(`TestFindPeakElement failed`)
	}
}

func TestMinWindow(t *testing.T) {
	if minWindow("XDOYEZODEYXNZ", "XYZ") != "YXNZ" || minWindow("ab", "b") != "b" || minWindow("aa", "aaa") != "" {
		t.Error(`TestMinWindow failed`)
	}
	// fmt.Println(minWindow("XDOYEZODEYXNZ", "XYZ"))
	// fmt.Println(minWindow("ab", "a"))
	// fmt.Println(minWindow("aaaa", "aaaa"))
}

func TestMinMoney(t *testing.T) {
	if minMoney([]int{2, 3, 5}, 20) != 4 {
		t.Error(`TestMinMoney failed`)
	}
}
