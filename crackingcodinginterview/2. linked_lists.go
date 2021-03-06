package crackingcodinginterview

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Node2Array(head *ListNode) []int {
	real := []int{}
	for head != nil {
		real = append(real, head.Val)
		head = head.Next
	}
	return real
}

func SprintNode(head *ListNode) string {
	real := []int{}
	for head != nil {
		real = append(real, head.Val)
		head = head.Next
	}
	return fmt.Sprint(real)
}

func PrintNode(head *ListNode) {
	real := []int{}
	for head != nil {
		real = append(real, head.Val)
		head = head.Next
	}
	fmt.Println(real)
}

/**
面试题 02.01. 移除重复节点
编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1:

 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
示例2:

 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
提示：

链表长度在[0, 20000]范围内。
链表元素在[0, 20000]范围内。
进阶：

如果不得使用临时缓冲区，该怎么解决？
*/
func removeDuplicateNodes(head *ListNode) *ListNode {
	// 解法一：使用buffer，时间复杂度O(n)，空间复杂度O(n)
	// maps := make(map[int]struct{})
	// var root, prev *ListNode
	// for head != nil {
	// 	if _, ok := maps[head.Val]; ok {
	// 		prev.Next = head.Next
	// 	} else {
	// 		if prev == nil {
	// 			root = head
	// 		}
	// 		prev = head
	// 		maps[head.Val] = struct{}{}
	// 	}
	// 	head = head.Next
	// }
	// return root

	// 解法二：不使用buffer，时间复杂度O(n^2)，空间复杂度O(1)（！！！）
	current := head // current为head前进的指针，每进一步，head加一个结点，同时去除后续值重复的结点
	for current != nil {
		runner := current // runner为current的指针，用于前移去除与current重复的结点
		for runner.Next != nil {
			if runner.Next.Val == current.Val { // 与current值相同的结点，直接跳过
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		current = current.Next // current前移
	}
	return head
}

/**
面试题 02.02. 返回倒数第 k 个节点
实现一种算法，找出单向链表中倒数第 k 个节点。返回该节点的值。

注意：本题相对原题稍作改动

示例：

输入： 1->2->3->4->5 和 k = 2
输出： 4
说明：

给定的 k 保证是有效的。
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func kthToLast(head *ListNode, k int) int {
	// 解法一：迭代，时间复杂度O(N)，空间复杂度O(1)
	// fast := head
	// i := 1
	// for ; i <= k && fast != nil; i++ {
	// 	fast = fast.Next
	// }

	// slow := head
	// for fast != nil {
	// 	slow = slow.Next
	// 	fast = fast.Next
	// }
	// return slow.Val

	// 解法二：递归，时间复杂度O(N)，空间复杂度O(N)
	index := 0
	var kthToLastSub func(root *ListNode, k int) *ListNode
	kthToLastSub = func(root *ListNode, k int) *ListNode {
		if root == nil {
			return nil
		}

		node := kthToLastSub(root.Next, k) // 后序遍历
		index += 1                         // 相当于从后往前计数，因此index=k时，表示找到了倒数第k个结点
		if index == k {
			return root
		}
		return node
	}
	return kthToLastSub(head, k).Val
}

/**
面试题 02.03. 删除中间节点
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。

假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。

例如，传入节点 c（位于单向链表 a->b->c->d->e->f 中），将其删除后，剩余链表为 a->b->d->e->f



示例：

输入：节点 5 （位于单向链表 4->5->1->9 中）
输出：不返回任何数据，从链表中删除传入的节点 5，使链表变为 4->1->9
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

/**
面试题 02.04. 分割链表
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你不需要 保留 每个分区中各节点的初始相对位置。



示例 1：


输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]
示例 2：

输入：head = [2,1], x = 2
输出：[1,2]


提示：

链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	// 解法一：造两个头结点，分别用于存储小于、大于等于的结点
	// big, small := &ListNode{}, &ListNode{}
	// bighead, smallhead := big, small
	// for head != nil {
	// 	next := head.Next
	// 	if head.Val < x {
	// 		small.Next = head
	// 		small = head
	// 	} else {
	// 		big.Next = head
	// 		big = head
	// 	}
	// 	head = next
	// }
	// big.Next = nil            // 避免循环，big的最后一个结点的Next置nil
	// small.Next = bighead.Next // small的最后一个结点的Next置为big开始结点
	// return smallhead.Next     // 返回small第一个结点

	// 解法二：原地（第一个结点会先成环）
	start, end := head, head
	for head != nil {
		next := head.Next
		if head.Val < x { // 小于，放在最前面
			head.Next = start
			start = head
		} else { // 大于，放在最后面
			end.Next = head
			end = head
		}
		head = next // next用来保证指针不断前移
	}
	end.Next = nil // 避免循环，最后一个结点的Next置为nil
	return start
}

/**
面试题 02.05. 链表求和
给定两个用链表表示的整数，每个节点包含一个数位。

这些数位是反向存放的，也就是个位排在链表首部。

编写函数对这两个整数求和，并用链表形式返回结果。



示例：

输入：(7 -> 1 -> 6) + (5 -> 9 -> 2)，即617 + 295
输出：2 -> 1 -> 9，即912
进阶：思考一下，假设这些数位是正向存放的，又该如何解决呢?

示例：

输入：(6 -> 1 -> 7) + (2 -> 9 -> 5)，即617 + 295
输出：9 -> 1 -> 2，即912
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, pre *ListNode
	mod := 0
	var dfsAddTwoNumbers func(d1, d2 *ListNode)
	dfsAddTwoNumbers = func(d1, d2 *ListNode) {
		if d1 == nil && d2 == nil {
			if mod > 0 {
				pre.Next = &ListNode{Val: mod}
			}
			return
		}

		sum := mod
		var d1next, d2next *ListNode
		if d1 != nil {
			sum += d1.Val
			d1next = d1.Next
		}
		if d2 != nil {
			sum += d2.Val
			d2next = d2.Next
		}
		mod = sum / 10

		tmp := &ListNode{Val: sum % 10}
		if pre == nil {
			head = tmp
			pre = tmp
		} else {
			pre.Next = tmp
			pre = tmp
		}
		dfsAddTwoNumbers(d1next, d2next)
	}
	dfsAddTwoNumbers(l1, l2)
	return head
}

/**
面试题 02.06. 回文链表
编写一个函数，检查输入的链表是否是回文的。



示例 1：

输入： 1->2
输出： false
示例 2：

输入： 1->2->2->1
输出： true


进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// 解法一：反转链表，再比对
	reverse := ReverseList(head)
	for head != nil {
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	return true

	// 解法二：快慢指针（借助stack比对前后两部分的值是否一致）
	// slow, fast := head, head
	// stack := []int{}
	// for fast != nil && fast.Next != nil {
	// 	stack = append(stack, slow.Val)
	// 	slow = slow.Next
	// 	fast = fast.Next.Next
	// }

	// if fast != nil { // 奇数个结点，跳过中间结点
	// 	slow = slow.Next
	// }

	// for slow != nil {
	// 	top := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]
	// 	if top != slow.Val {
	// 		return false
	// 	}
	// 	slow = slow.Next
	// }
	// return true
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
面试题 02.07. 链表相交
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

图示两个链表在节点 c1 开始相交：



题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。



示例 1：



输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
示例 2：



输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
示例 3：



输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。


提示：

listA 中节点数目为 m
listB 中节点数目为 n
0 <= m, n <= 3 * 104
1 <= Node.val <= 105
0 <= skipA <= m
0 <= skipB <= n
如果 listA 和 listB 没有交点，intersectVal 为 0
如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	if headA == headB {
		return headA
	}

	// 获取链表最后一个结点及长度
	var getTailAndSize func(head *ListNode) (*ListNode, int)
	getTailAndSize = func(head *ListNode) (*ListNode, int) {
		size := 1
		for head.Next != nil {
			size++
			head = head.Next
		}
		return head, size
	}

	tailA, sizeA := getTailAndSize(headA)
	tailB, sizeB := getTailAndSize(headB)
	if tailA != tailB {
		return nil
	}
	var short, long *ListNode
	if sizeA < sizeB {
		short, long = headA, headB
	} else {
		short, long = headB, headA
	}

	// 去除较长链表之前的结点，使两个链表等长
	diffsize := abs(sizeA - sizeB)
	for diffsize > 0 {
		long = long.Next
		diffsize--
	}

	for short != long {
		short = short.Next
		long = long.Next
	}

	return long
}

// 更简洁的方法
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}

		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

/**
面试题 02.08. 环路检测
给定一个链表，如果它是有环链表，实现一个算法返回环路的开头节点。若环不存在，请返回 null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。



示例 1：



输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
示例 2：



输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：



输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。


进阶：

你是否可以不用额外空间解决此题？
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
