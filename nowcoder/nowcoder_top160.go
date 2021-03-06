package nowcoder

import (
	"math"
)

/**
NC118 数组中的逆序对
 算法知识视频讲解
中等  通过率：16.63%  时间限制：3秒  空间限制：64M
知识点
数组
题目
题解(70)
讨论(869)
排行
描述
在数组中的两个数字，如果前面一个数字大于后面的数字，则这两个数字组成一个逆序对。输入一个数组,求出这个数组中的逆序对的总数P。并将P对1000000007取模的结果输出。 即输出P mod 1000000007

数据范围：  对于 50\%50% 的数据, size\leq 10^4size≤10
4

对于 100\%100% 的数据, size\leq 10^5size≤10
5

数组中所有数字的值满足 0 \le val \le 10000000≤val≤1000000

要求：空间复杂度 O(n)O(n)，时间复杂度 O(nlogn)O(nlogn)
输入描述：
题目保证输入的数组中没有的相同的数字
示例1
输入：
[1,2,3,4,5,6,7,0]
复制
返回值：
7
复制
示例2
输入：
[1,2,3]
复制
返回值：
0
*/
func InversePairs(data []int) int {
	// write code here
	if len(data) < 2 {
		return 0
	}

	count := 0

	// 归并
	merges := func(arr []int, left, mid, right int) {
		tmp := make([]int, right-left+1)
		c, s, l, r := 0, left, left, mid+1
		for l <= mid && r <= right {
			if arr[l] <= arr[r] { // 无逆序
				tmp[c] = arr[l]
				c++
				l++
			} else {
				tmp[c] = arr[r]
				count += mid + 1 - l // mid+1-l是arr[r]的逆序度
				count %= 1000000007
				c++
				r++
			}
		}

		for l <= mid {
			tmp[c] = arr[l]
			l++
			c++
		}

		for r <= right {
			tmp[c] = arr[r]
			r++
			c++
		}

		for _, num := range tmp {
			arr[s] = num
			s++
		}
	}

	var mergeSort func([]int, int, int)
	mergeSort = func(arr []int, left, right int) {
		mid := (left + right) >> 1
		if left < right {
			mergeSort(arr, left, mid)
			mergeSort(arr, mid+1, right)
			merges(arr, left, mid, right)
		}
	}

	mergeSort(data, 0, len(data)-1)
	return count
}

/**
NC120 二进制中1的个数
 算法知识视频讲解
中等  通过率：35.39%  时间限制：1秒  空间限制：64M
知识点
数学
题目
题解(125)
讨论(1k)
排行
描述
输入一个整数 n ，输出该数32位二进制表示中1的个数。其中负数用补码表示。

数据范围：- 2^{31} <= n <= 2^{31}-1−2
31
 <=n<=2
31
 −1
即范围为:-2147483648<= n <= 2147483647−2147483648<=n<=2147483647
示例1
输入：
10
复制
返回值：
2
复制
说明：
十进制中10的32位二进制表示为0000 0000 0000 0000 0000 0000 0000 1010，其中有两个1。
示例2
输入：
-1
复制
返回值：
32
复制
说明：
负数使用补码表示 ，-1的32位二进制表示为1111 1111 1111 1111 1111 1111 1111 1111，其中32个1
*/
func NumberOf1(n int) int {
	// write code here
	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	}

	num := 0
	bits := make([]int, 32)
	for i := 31; i >= 0; i-- {
		if n < 1<<i {
			continue
		} else {
			num++
			n -= 1 << i
			bits[i] = 1
		}
	}

	if isNegative {
		num = 0
		incr := true
		for i := 0; i <= 31; i++ {
			tmpInt := (bits[i] + 1) % 2
			if incr {
				tmpInt += 1
			}
			if tmpInt > 1 {
				tmpInt = 0
				incr = true
			} else {
				if tmpInt == 1 {
					num++
				}
				incr = false
			}
			bits[i] = tmpInt
		}
	}

	return num
}

/**
NC108 最大正方形
 算法知识视频讲解
中等  通过率：44.92%  时间限制：1秒  空间限制：64M
知识点
动态规划
题目
题解(22)
讨论(46)
排行
描述
给定一个由'0'和'1'组成的2维矩阵，返回该矩阵中最大的由'1'组成的正方形的面积，输入的矩阵是字符形式而非数字形式。

数据范围：矩阵的长宽满足 0 \le n \le 200≤n≤20,矩阵中的元素属于 {'1','0'}
进阶：空间复杂度 O(n^2)O(n
2
 ) ， 时间复杂度 O(n^2)O(n
2
 )
示例1
输入：
[[1,0,1,0,0],[1,0,1,1,1],[1,1,1,1,1],[1,0,0,1,0]]
复制
返回值：
4
复制
示例2
输入：
[[1,0,0],[0,0,0],[0,0,0]]
复制
返回值：
1
*/
func solveMaxSquare(matrix [][]byte) int {
	// write code here
	if len(matrix) == 0 {
		return 0
	}

	n := len(matrix)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				if max < 1 {
					max = 1
				}
				continue
			}
			temp := minInts(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			if temp > max {
				max = temp
			}
			dp[i][j] = temp
		}
	}
	// fmt.Println(dp)
	return max * max
}

func minInts(a, b, c int) int {
	min := a
	if b < min {
		min = b
	}

	if c < min {
		min = c
	}
	return min
}

/**
NC110 旋转数组
 算法知识视频讲解
入门  通过率：46.60%  时间限制：1秒  空间限制：64M
知识点
数组
题目
题解(25)
讨论(58)
排行
描述
一个数组A中存有 n 个整数，在不允许使用另外数组的前提下，将每个整数循环向右移 M（ M >=0）个位置，即将A中的数据由（A0 A1 ……AN-1 ）变换为（AN-M …… AN-1 A0 A1 ……AN-M-1 ）（最后 M 个数循环移至最前面的 M 个位置）。如果需要考虑程序移动数据的次数尽量少，要如何设计移动的方法？

数据范围：0 < n \le 100<n≤10，0 \le m \le 10000≤m≤1000
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
6,2,[1,2,3,4,5,6]
复制
返回值：
[5,6,1,2,3,4]
复制
示例2
输入：
4,0,[1,2,3,4]
复制
返回值：
[1,2,3,4]
复制
备注：
(1<=N<=100,M>=0)
*/
func solveRotateArray(n int, m int, a []int) []int {
	// write code here
	if n < 2 || m == 0 {
		return a
	}

	leftRotateByOne := func(arr []int) []int {
		temp := arr[n-1]
		for i := n - 1; i > 0; i-- {
			arr[i] = arr[i-1]
		}
		arr[0] = temp
		return arr
	}

	for i := 0; i < m; i++ {
		a = leftRotateByOne(a)
	}
	return a
}

/**
NC122 正则表达式匹配
 算法知识视频讲解
中等  通过率：30.62%  时间限制：1秒  空间限制：256M
知识点
字符串
题目
题解(32)
讨论(42)
排行
描述
请实现一个函数用来匹配包括'.'和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（包含0次）。 在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配

数据范围：字符串长度 0 \le |str| \le 10000≤∣str∣≤1000。
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

提示：如果输入空字符串，则返回True
示例1
输入：
"aaa","a*a"
复制
返回值：
true
复制
示例2
输入：
"",".*"
复制
返回值：
true
*/
// func match( str string ,  pattern string ) bool {
//     // write code here
// }

/**
NC131 数据流中的中位数
 算法知识视频讲解
中等  通过率：27.35%  时间限制：1秒  空间限制：64M
知识点
排序
堆
题目
题解(69)
讨论(664)
排行
描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
示例1
输入：
[5,2,3,4,1,6,7,0,8]
复制
返回值：
"5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00 "
复制
说明：
数据流里面不断吐出的是5,2,3...,则得到的平均数分别为5,(5+2)/2,3...
*/
// func Insert(num int){

// }

// func GetMedian() float64{

// }

/**
NC39 N皇后问题
 算法知识视频讲解
较难  通过率：41.69%  时间限制：2秒  空间限制：256M
知识点
回溯
题目
题解(14)
讨论(42)
排行
描述
N 皇后问题是指在 n * n 的棋盘上要摆 n 个皇后，
要求：任何两个皇后不同行，不同列也不再同一条斜线上，
求给一个整数 n ，返回 n 皇后的摆法数。

数据范围: 1 \le n \le 91≤n≤9
要求：空间复杂度 O(1)O(1) ，时间复杂度 O(n!)O(n!)
示例1
输入：
1
复制
返回值：
1
复制
示例2
输入：
8
复制
返回值：
92
*/
// func Nqueen( n int ) int {
//     // write code here
// }

/**
NC124 字典树的实现
 算法知识视频讲解
中等  通过率：39.83%  时间限制：2秒  空间限制：256M
知识点
高级结构
字符串
题目
题解(14)
讨论(31)
排行
描述
字典树又称为前缀树或者Trie树，是处理字符串常用的数据结构。

假设组成所有单词的字符仅是‘a’～‘z’，请实现字典树的结构，并包含以下四个主要的功能。

1. void insert(String word)：添加word，可重复添加；
2. void delete(String word)：删除word，如果word添加过多次，仅删除一次；
3. boolean search(String word)：查询word是否在字典树中出现过(完整的出现过，前缀式不算)；
4. int prefixNumber(String pre)：返回以字符串pre作为前缀的单词数量。

现在给定一个m，表示有m次操作，每次操作都为以上四种操作之一。每次操作会给定一个整数op和一个字符串word，op代表一个操作码，如果op为1，则代表添加word，op为2则代表删除word，op为3则代表查询word是否在字典树中，op为4代表返回以word为前缀的单词数量（数据保证不会删除不存在的word）。

对于每次操作，如果op为3时，如果word在字典树中，请输出“YES”，否则输出“NO”；如果op为4时，请输出返回以word为前缀的单词数量，其它情况不输出。
数据范围：操作数满足 0\le m \le 10^50≤m≤10
5
 ，字符串长度都满足 0 \le n \le 200≤n≤20
进阶：所有操作的时间复杂度都满足 O(n)O(n)
示例1
输入：
[["1","qwer"],["1","qwe"],["3","qwer"],["4","q"],["2","qwer"],["3","qwer"],["4","q"]]
复制
返回值：
["YES","2","NO","1"]
复制
备注：
m \leq 10^5m≤10
5

|word| \leq 20∣word∣≤20
*/
// func trieU( operators [][]string ) []string {
//     // write code here
// }

/**
NC80 把二叉树打印成多行
 算法知识视频讲解
中等  通过率：32.92%  时间限制：1秒  空间限制：64M
知识点
树
bfs
题目
题解(79)
讨论(837)
排行
描述
给定一个节点数为 n 二叉树，要求从上到下按层打印二叉树的 val 值，同一层结点从左至右输出。每一层输出一行。
例如：
给定的二叉树是{1,2,3,#,#,4,5}

该二叉树多行打印层序遍历的结果是
[
[1],
[2,3],
[4,5]
]

数据范围：二叉树的节点数 0 \le n \le 10000≤n≤1000，0 \le val \le 10000≤val≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
输入描述：
给定一个二叉树的根节点
示例1
输入：
{1,2,3,#,#,4,5}
复制
返回值：
[[1],[2,3],[4,5]]
复制
示例2
输入：
{8,6,10,5,7,9,11}
复制
返回值：
[[8],[6,10],[5,7,9,11]]
复制
示例3
输入：
{1,2,3,4,5}
复制
返回值：
[[1],[2,3],[4,5]]
复制
示例4
输入：
{}
复制
返回值：
[]
*/
func PrintTree(pRoot *TreeNode) [][]int {
	// write code here
	if pRoot == nil {
		return [][]int{}
	}
	var ret [][]int
	var tmp []int
	var queue []*TreeNode
	queue = append(queue, pRoot)
	for len(queue) > 0 {
		lens := len(queue)
		tmp = []int{}
		for i := 0; i < lens; i++ {
			tmpNode := queue[i]
			tmp = append(tmp, tmpNode.Val)
			if tmpNode.Left != nil {
				queue = append(queue, tmpNode.Left)
			}

			if tmpNode.Right != nil {
				queue = append(queue, tmpNode.Right)
			}
		}
		ret = append(ret, tmp)
		queue = queue[lens:]
	}
	return ret
}

/**
NC83 子数组最大乘积
 算法知识视频讲解
中等  通过率：35.27%  时间限制：1秒  空间限制：64M
知识点
数组
动态规划
题目
题解(26)
讨论(49)
排行
描述
给定一个double类型的数组arr，其中的元素可正可负可0，返回连续子数组累乘的最大乘积。

数据范围：数组大小满足 0 \le n \le 100≤n≤10，数组中元素满足 |val| \le 10∣val∣≤10
进阶：空间复杂度 O(1)O(1) ，时间复杂度 O(n)O(n)
示例1
输入：
[-2.5,4,0,3,0.5,8,-1]
复制
返回值：
12.00000
复制
示例2
输入：
[1.0,0.0,0.0]
复制
返回值：
1.00000
*/
// func maxProduct( arr []float64 ) float64 {
//     // write code here
// }

/**
NC116 把数字翻译成字符串
 算法知识视频讲解
中等  通过率：27.16%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(21)
讨论(26)
排行
描述
有一种将字母编码成数字的方式：'a'->1, 'b->2', ... , 'z->26'。

现在给一串数字，返回有多少种可能的译码结果

数据范围：字符串长度满足 0 < n \le 900<n≤90
进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
"12"
复制
返回值：
2
复制
说明：
2种可能的译码结果（”ab” 或”l”）
示例2
输入：
"31717126241541717"
复制
返回值：
192
复制
说明：
192种可能的译码结果
*/
// func solve( nums string ) int {
//     // write code here
// }

/**
NC134 股票(无限次交易)
 算法知识视频讲解
简单  通过率：58.69%  时间限制：1秒  空间限制：256M
知识点
贪心
题目
题解(23)
讨论(42)
排行
描述
假定你知道 n 天内的某只股票每一天价格的变动。
你最多可以同时持有一只股票。但你可以无限次的交易（买进和卖出均无手续费）。
请设计一个函数，计算你所能获得的最大收益。

输入一个数组，数组长度为 n ，数组中每一个元素为 arri 代表当天股票的价格，

数据范围： 0 \le n \le 2 \times 10^50≤n≤2×10
5
  ， 1 \le arr_i \le 10^51≤arr
i
​
 ≤10
5

进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
[5,4,3,2,1]
复制
返回值：
0
复制
说明：
由于每天股票都在跌，因此不进行任何交易最优。最大收益为0。
示例2
输入：
[1,2,3,4,5]
复制
返回值：
4
复制
说明：
第一天买进，最后一天卖出最优。中间的当天买进当天卖出不影响最终结果。最大收益为4。
备注：
总天数不大于200000。保证股票每一天的价格在[1,100]范围内。
*/
func maxProfit134(prices []int) int {
	// write code here
	days := len(prices)
	if days <= 1 {
		return 0
	}

	max := 0
	for i := 1; i < days; i++ {
		if prices[i] > prices[i-1] {
			max += prices[i] - prices[i-1]
		}
	}
	return max
}

/**
NC44 通配符匹配
 算法知识视频讲解
较难  通过率：31.02%  时间限制：1秒  空间限制：64M
知识点
贪心
字符串
动态规划
回溯
题目
题解(14)
讨论(49)
排行
描述
请实现支持'?'and'*'.的通配符模式匹配
'?' 可以匹配任何单个字符。
'*' 可以匹配任何字符序列（包括空序列）。
返回两个字符串是否匹配
函数声明为：
bool isMatch(const char *s, const char *p)
下面给出一些样例：
isMatch("aa","a") → false
isMatch("aa","aa") → true
isMatch("aaa","aa") → false
isMatch("aa", "*") → true
isMatch("aa", "a*") → true
isMatch("ab", "?*") → true
isMatch("aab", "d*a*b") → false
数据范围：字符串长度满足 0 \le n \le 10000≤n≤1000
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
"ab","?*"
复制
返回值：
true
复制
示例2
输入：
"ab","*"
复制
返回值：
true
*/
// func isMatch( s string ,  p string ) bool {
//     // write code here
// }

/**
NC72 二叉树的镜像
 算法知识视频讲解
中等  通过率：64.69%  时间限制：1秒  空间限制：256M
知识点
树
题目
题解(78)
讨论(99)
排行
描述
操作给定的二叉树，将其变换为源二叉树的镜像。
数据范围：二叉树的节点数 0 \le n \le 10000≤n≤1000 ， 二叉树每个节点的值 0\le val \le 10000≤val≤1000
要求： 空间复杂度 O(n)O(n) 。本题也有原地操作，即空间复杂度 O(1)O(1) 的解法，时间复杂度 O(n)O(n)

比如：
                                                                    源二叉树

                                                                    镜像二叉树

示例1
输入：
{8,6,10,5,7,9,11}
复制
返回值：
{8,10,6,11,9,7,5}
复制
说明：
如题面所示
示例2
输入：
{}
复制
返回值：
{}
*/
// func Mirror( pRoot *TreeNode ) *TreeNode {
//     // write code here
// }

/**
NC77 调整数组顺序使奇数位于偶数前面(一)
 算法知识视频讲解
简单  通过率：51.42%  时间限制：1秒  空间限制：256M
知识点
数组
题目
题解(82)
讨论(114)
排行
描述
输入一个长度为 n 整数数组，数组里面不含有相同的元素，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

数据范围：0 \le n \le 50000≤n≤5000，数组中每个数的值 0 \le val \le 100000≤val≤10000
要求：时间复杂度 O(n)O(n)，空间复杂度 O(n)O(n)
进阶：时间复杂度 O(n^2)O(n
2
 )，空间复杂度 O(1)O(1)
示例1
输入：
[1,2,3,4]
复制
返回值：
[1,3,2,4]
复制
示例2
输入：
[2,4,6,5,7]
复制
返回值：
[5,7,2,4,6]
复制
示例3
输入：
[1,3,5,6,7]
复制
返回值：
[1,3,5,7,6]
*/
func reOrderArray(array []int) []int {
	// write code here
	if len(array) <= 1 {
		return array
	}

	evenS := -1
	ret := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		if array[i]%2 == 0 { // 偶数
			ret[i] = array[i]
			if evenS < 0 {
				evenS = i
			}
		} else { // 奇数
			if evenS < 0 {
				ret[i] = array[i]
				continue
			}
			for j := i; j > evenS; j-- {
				ret[j] = ret[j-1]
			}
			ret[evenS] = array[i]
			evenS++
		}
	}
	return ret
}

/**
NC98 判断t1树中是否有与t2树完全相同的子树
 算法知识视频讲解
简单  通过率：40.17%  时间限制：2秒  空间限制：256M
知识点
树
题目
题解(16)
讨论(30)
排行
描述
给定彼此独立的两棵二叉树，树上的节点值两两不同，判断 t1 树是否有与 t2 树完全相同的子树。

子树指一棵树的某个节点的全部后继节点

数据范围：树的节点数满足 0 < n \le 5000000<n≤500000，树上每个节点的值一定在32位整型范围内
进阶：空间复杂度: O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,3,4,5,6,7,#,8,9},{2,4,5,#,8,9}
复制
返回值：
true
复制
备注：
1 \leq n \leq 5000001≤n≤500000
*/
func isContains(root1 *TreeNode, root2 *TreeNode) bool {
	// write code here
	if root1 == nil && root2 != nil {
		return false
	}
	if root2 == nil {
		return true
	}

	ret := false

	var isSame func(t1, t2 *TreeNode) bool
	isSame = func(t1, t2 *TreeNode) bool {
		if t1 == nil && t2 == nil {
			return true
		}

		if t1 == nil && t2 != nil || t1 != nil && t2 == nil {
			return false
		}

		// 根结点一致
		if t1.Val != t2.Val {
			return false
		}

		// 子树也一致
		if isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right) {
			return true
		}
		return false
	}

	var dfsTree func(t1 *TreeNode)
	dfsTree = func(t1 *TreeNode) {
		if ret { // 已存在直接退出
			return
		}
		if t1.Val == root2.Val && isSame(t1, root2) {
			ret = true
			return
		}

		if t1.Left != nil {
			dfsTree(t1.Left)
		}

		if t1.Right != nil {
			dfsTree(t1.Right)
		}
	}

	// 遍历root1
	dfsTree(root1)
	return ret
}

/**
NC117 合并二叉树
 算法知识视频讲解
简单  通过率：69.81%  时间限制：1秒  空间限制：256M
知识点
树
题目
题解(19)
讨论(35)
排行
描述
已知两颗二叉树，将它们合并成一颗二叉树。合并规则是：都存在的结点，就将结点值加起来，否则空的位置就由另一个树的结点来代替。例如：
两颗二叉树是:
                                                                    Tree 1


                                                                        Tree 2

                                                                    合并后的树为

数据范围：树上节点数量满足 0 \le n \le 5000≤n≤500，树上节点的值一定在32位整型范围内。
进阶：空间复杂度 O(1)O(1) ，时间复杂度 O(n)O(n)
示例1
输入：
{1,3,2,5},{2,1,3,#,4,#,7}
复制
返回值：
{3,4,5,5,4,#,7}
复制
说明：
如题面图
示例2
输入：
{1},{}
复制
返回值：
{1}
*/
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	// write code here
	var mergeTreeSub func(r1, r2 *TreeNode) *TreeNode
	mergeTreeSub = func(r1, r2 *TreeNode) *TreeNode {
		if r1 == nil {
			return r2
		}

		if r2 == nil {
			return r1
		}

		head := &TreeNode{Val: r1.Val + r2.Val}
		head.Left = mergeTreeSub(r1.Left, r2.Left)
		head.Right = mergeTreeSub(r1.Right, r2.Right)
		return head
	}
	return mergeTreeSub(t1, t2)
}

/**
NC135 股票交易的最大收益（二）
 算法知识视频讲解
中等  通过率：46.44%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(19)
讨论(24)
排行
描述
假定你知道某只股票每一天价格的变动。
你最多可以同时持有一只股票。但你最多只能进行两次交易（一次买进和一次卖出记为一次交易。买进和卖出均无手续费）。
请设计一个函数，计算你所能获得的最大收益。

数据范围：0 \le n \le 2000000≤n≤200000，股票的价格满足 1 \le val\le 1000001≤val≤100000
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
[8,9,3,5,1,3]
复制
返回值：
4
复制
说明：
第三天买进，第四天卖出，第五天买进，第六天卖出。总收益为4。
示例2
输入：
[9,8,4,1]
复制
返回值：
0
复制
备注：
总天数不大于200000。保证股票每一天的价格在[1,100]范围内。
*/
// func maxProfit( prices []int ) int {
//     // write code here
// }

/**
NC138 矩阵最长递增路径
 算法知识视频讲解
中等  通过率：48.35%  时间限制：5秒  空间限制：256M
知识点
dfs
题目
题解(14)
讨论(18)
排行
描述
给定一个 n 行 m 列矩阵 matrix ，矩阵内所有数均为非负整数。 你需要在矩阵中找到一条最长路径，使这条路径上的元素是递增的。并输出这条最长路径的长度。
这个路径必须满足以下条件：

1. 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外。
2. 你不能走重复的单元格。即每个格子最多只能走一次。

数据范围：1 \le n,m \le 10001≤n,m≤1000，0 \le matrix[i][j] \le 10000≤matrix[i][j]≤1000
进阶：空间复杂度 O(nm)O(nm) ，时间复杂度 O(nm)O(nm)
示例1
输入：
[[1,2,3],[4,5,6],[7,8,9]]
复制
返回值：
5
复制
说明：
1->2->3->6->9即可。当然这种递增路径不是唯一的。
示例2
输入：
[[1,2],[4,3]]
复制
返回值：
4
复制
说明：
 1->2->3->4

备注：
矩阵的长和宽均不大于1000，矩阵内每个数不大于1000
*/
// func solve( matrix [][]int ) int {
//     // write code here
// }

/**
NC143 矩阵乘法
 算法知识视频讲解
中等  通过率：58.38%  时间限制：1秒  空间限制：256M
知识点
数组
模拟
题目
题解(10)
讨论(16)
排行
描述
给定两个 n*n 的矩阵 A 和 B ，求 A*B 。

数据范围：1 \le n \le 1001≤n≤100，-100 \le Matrix_{i,j}\le 100−100≤Matrix
i,j
​
 ≤100

要求：空间复杂度 O(n^2)O(n
2
 ) ， 时间复杂度 O(n^3 )O(n
3
 )
进阶：本题也有空间复杂度 O(n^2)O(n
2
 )，时间复杂度 O(n^{log7})O(n
log7
 )的解法
PS：更优时间复杂度的算法这里并不考察
示例1
输入：
[[1,2],[3,2]],[[3,4],[2,1]]
复制
返回值：
[[7,6],[13,14]]
复制
示例2
输入：
[[1]],[[1]]
复制
返回值：
[[1]]
*/
// func solve( a [][]int ,  b [][]int ) [][]int {
//     // write code here
// }

/**
NC64 二叉搜索树与双向链表
 算法知识视频讲解
中等  通过率：29.71%  时间限制：1秒  空间限制：64M
知识点
分治
题目
题解(120)
讨论(1k)
排行
描述
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示


数据范围：输入二叉树的节点数 0 \le n \le 10000≤n≤1000，二叉树中每个节点的值 0\le val \le 10000≤val≤1000
要求：空间复杂度O(1)O(1)（即在原树上操作），时间复杂度 O(n)O(n)

注意:
1.要求不能创建任何新的结点，只能调整树中结点指针的指向。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继
2.返回链表中的第一个节点的指针
3.函数返回的TreeNode，有左右指针，其实可以看成一个双向链表的数据结构
4.你不用输出双向链表，程序会根据你的返回值自动打印输出
输入描述：
二叉树的根节点
返回值描述：
双向链表的其中一个头节点。
示例1
输入：
{10,6,14,4,8,12,16}
复制
返回值：
From left to right are:4,6,8,10,12,14,16;From right to left are:16,14,12,10,8,6,4;
复制
说明：
输入题面图中二叉树，输出的时候将双向链表的头节点返回即可。
示例2
输入：
{5,4,#,3,#,2,#,1}
复制
返回值：
From left to right are:1,2,3,4,5;From right to left are:5,4,3,2,1;
复制
说明：
                    5
                  /
                4
              /
            3
          /
        2
      /
    1
树的形状如上图
*/
// func Convert( pRootOfTree *TreeNode ) *TreeNode {
//     // write code here
// }

/**
NC129 阶乘末尾0的数量
 算法知识视频讲解
中等  通过率：45.70%  时间限制：2秒  空间限制：256M
知识点
数学
题目
题解(17)
讨论(17)
排行
描述
给定一个非负整数 n ，返回 n! 结果的末尾为 0 的数量。

n! 是指自然数 n! 的阶乘,即 : N!=1 \times 2\times3…(N-2)\times(N-1) \times NN!=1×2×3…(N−2)×(N−1)×N。
特殊的,  0 的阶乘是 1 。

数据范围： 0 < n \le 10^{14}0<n≤10
14

进阶：空间复杂度 O(1)O(1)，时间复杂度 O(logn)O(logn)
复杂度要求：
不大于 O(\log n)O(logn)
示例1
输入：
3
复制
返回值：
0
复制
说明：
3!=6
示例2
输入：
5
复制
返回值：
1
复制
说明：
5!=120
示例3
输入：
1000000000
复制
返回值：
249999998
*/
// func thenumberof0( n int64 ) int64 {
//     // write code here
// }

/**
NC58 找到搜索二叉树中两个错误的节点
 算法知识视频讲解
简单  通过率：39.17%  时间限制：2秒  空间限制：256M
知识点
树
dfs
题目
题解(11)
讨论(27)
排行
描述
一棵二叉树原本是搜索二叉树，但是其中有两个节点调换了位置，使得这棵二叉树不再是搜索二叉树，请按升序输出这两个错误节点的值。(每个节点的值各不相同)
搜索二叉树：满足每个节点的左子节点小于当前节点，右子节点大于当前节点。
样例1图

样例2图

数据范围：3 \le n \le 1000003≤n≤100000,节点上的值满足 1 \le val \le n1≤val≤n ，保证每个value各不相同
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,3}
复制
返回值：
[1,2]
复制
说明：
如题面图
示例2
输入：
{4,2,5,3,1}
复制
返回值：
[1,3]
*/
func findError(root *TreeNode) []int {
	// write code here
	arr := order(root)
	if len(arr) < 2 {
		return []int{}
	}
	min, max := math.MinInt64, math.MinInt64
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			if max == math.MinInt64 {
				max = i
			} else {
				min = i + 1
			}
			continue
		}
	}
	if max == math.MinInt64 {
		return []int{}
	}

	if min == math.MinInt64 {
		min = max + 1
	}

	return []int{arr[min], arr[max]}
}

/**
NC142 最长重复子串
 算法知识视频讲解
中等  通过率：34.96%  时间限制：5秒  空间限制：256M
知识点
字符串
题目
题解(18)
讨论(27)
排行
描述
定义重复字符串是由两个相同的字符串首尾拼接而成，例如 abcabcabcabc 便是长度为6的一个重复字符串，而 abcbaabcba 则不存在重复字符串。

给定一个字符串，请返回其最长重复子串的长度。

若不存在任何重复字符子串，则返回 0 。

本题中子串的定义是字符串中一段连续的区间。

数据范围：字符串长度不大于 1 \times 10^31×10
3
 ，保证字符串一定由小写字母构成。
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n^2)O(n
2
 )
示例1
输入：
"ababc"
复制
返回值：
4
复制
说明：
abab为最长的重复字符子串，长度为4
示例2
输入：
"abcab"
复制
返回值：
0
复制
说明：
该字符串没有重复字符子串
备注：
字符串长度不超过10000，且仅由小写字母组成
*/
// func solve( a string ) int {
//     // write code here
// }

/**
NC144 不相邻最大子序列和
 算法知识视频讲解
中等  通过率：44.97%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(22)
讨论(32)
排行
描述
给你一个数组，其长度为 n  ，在其中选出一个子序列，子序列中任意两个数不能有相邻的下标（子序列可以为空）

本题中子序列指在数组中任意挑选若干个数组成的数组。

数据范围：1 \le n \le 10^51≤n≤10
5
 ，数组中所有数的值满足 |val| \le 2147483647∣val∣≤2147483647
进阶：空间复杂度 O(1)O(1) ， 时间复杂度 O(n)O(n)
示例1
输入：
3,[1,2,3]
复制
返回值：
4
复制
说明：
有[],[1],[2],[3],[1,3] 4种选取方式其中[1,3]选取最优，答案为4
示例2
输入：
4,[4,2,3,5]
复制
返回值：
9
复制
说明：
其中[4,5]的选取方案是在满足不同时选取相邻位置的数的情况下是最优的答案
示例3
输入：
1,[-1]
复制
返回值：
0
复制
说明：
选择子序列为空最优
示例4
输入：
5,[3,2,3,4,5]
复制
返回值：
11
复制
说明：
其中选择[3,3,5]的方案是最优的答案
*/
// func subsequence( n int ,  array []int ) int64 {
//     // write code here
// }

/**
NC11 将升序数组转化为平衡二叉搜索树
 算法知识视频讲解
简单  通过率：36.24%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(19)
讨论(84)
排行
描述
给定一个升序排序的数组，将其转化为平衡二叉搜索树（BST）.

平衡二叉搜索树指树上每个节点 node 都满足左子树中所有节点的的值都小于 node 的值，右子树中所有节点的值都大于 node 的值，并且左右子树的节点数量之差不大于1

数据范围：0 \le n \le 100000≤n≤10000，数组中每个值满足 |val| \le 5000∣val∣≤5000
进阶：空间复杂度 O(n)O(n) ，时间复杂度 O(n)O(n)
示例1
输入：
[-1,0,1,2]
复制
返回值：
{1,0,2,-1}
复制
示例2
输入：
[]
复制
返回值：
{}
*/
func sortedArrayToBST(num []int) *TreeNode {
	// write code here
	var head *TreeNode
	lens := len(num)
	if lens == 0 {
		return head
	}

	var dfsBST func(start, end int) *TreeNode
	dfsBST = func(start, end int) *TreeNode {
		if start > end {
			return nil
		}

		if start == end {
			return &TreeNode{Val: num[start]}
		}

		mid := (start+end)>>1 + 1 // 左测结点多于右侧
		root := &TreeNode{Val: num[mid]}
		root.Left = dfsBST(start, mid-1)
		root.Right = dfsBST(mid+1, end)
		return root
	}
	head = dfsBST(0, lens-1)
	return head
}

/**
描述
请编写一个程序，给数独中的剩余的空格填写上数字
空格用字符'.'表示
假设给定的数独只有唯一的解法

这盘数独的解法是：

红色表示填上的解
示例1
输入：
[[.,.,9,7,4,8,.,.,.],[7,.,.,.,.,.,.,.,.],[.,2,.,1,.,9,.,.,.],[.,.,7,.,.,.,2,4,.],[.,6,4,.,1,.,5,9,.],[.,9,8,.,.,.,3,.,.],[.,.,.,8,.,3,.,2,.],[.,.,.,.,.,.,.,.,6],[.,.,.,2,7,5,9,.,.]]
复制
返回值：
[[5,1,9,7,4,8,6,3,2],[7,8,3,6,5,2,4,1,9],[4,2,6,1,3,9,8,7,5],[3,5,7,9,8,6,2,4,1],[2,6,4,3,1,7,5,9,8],[1,9,8,5,2,4,3,6,7],[9,7,5,8,6,3,1,2,4],[8,3,2,4,9,1,7,5,6],[6,4,1,2,7,5,9,8,3]]
*/
// func solveSudoku( board [][]byte ) {
//     // write code here
// }

/**
NC63 扑克牌顺子
 算法知识视频讲解
简单  通过率：27.46%  时间限制：1秒  空间限制：64M
知识点
模拟
题目
题解(111)
讨论(1k)
排行
描述
现在有2副扑克牌，从扑克牌中随机五张扑克牌，我们需要来判断一下是不是顺子。
有如下规则：
1. A为1，J为11，Q为12，K为13，A不能视为14
2. 大、小王为 0，0可以看作任意牌
3. 如果给出的五张牌能组成顺子（即这五张牌是连续的）就输出true，否则就输出false。
4.数据保证每组5个数字，每组最多含有4个零，数组的数取值为 [0, 13]

要求：空间复杂度 O(1)O(1)，时间复杂度 O(nlogn)O(nlogn)，本题也有时间复杂度 O(n)O(n) 的解法
输入描述：
输入五张扑克牌的值
返回值描述：
五张扑克牌能否组成顺子。
示例1
输入：
[6,0,2,0,4]
复制
返回值：
true
复制
说明：
中间的两个0一个看作3，一个看作5 。即：[6,3,2,5,4]
这样这五张牌在[2,6]区间连续，输出true
示例2
输入：
[0,3,2,6,4]
复制
返回值：
true
复制
示例3
输入：
[1,0,0,1,0]
复制
返回值：
false
复制
示例4
输入：
[13,12,11,0,1]
复制
返回值：
false
*/
func IsContinuous(numbers []int) bool {
	// write code here
	maps := make(map[int]struct{})
	min, max, zeroNum := 15, 0, 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 0 {
			zeroNum++
		} else {
			if _, ok := maps[numbers[i]]; ok {
				return false
			}
			if numbers[i] < min {
				min = numbers[i]
			}

			if numbers[i] > max {
				max = numbers[i]
			}
			maps[numbers[i]] = struct{}{}
		}
	}

	if max-min > 4 || zeroNum < 4-(max-min) {
		return false
	}

	return true
}

/**
NC31 第一个只出现一次的字符
 算法知识视频讲解
简单  通过率：31.00%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(123)
讨论(1k)
排行
描述
在一个长为 字符串中找到第一个只出现一次的字符,并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）


数据范围：0 \le n \le 100000≤n≤10000，且字符串只有字母组成。
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
"google"
复制
返回值：
4
复制
示例2
输入：
"aa"
复制
返回值：
-1
*/
func FirstNotRepeatingChar(str string) int {
	// write code here
	if len(str) == 0 {
		return -1
	}

	maps := make(map[rune]int)
	for _, b := range str {
		maps[b]++
	}

	for i, b := range str {
		if maps[b] == 1 {
			return i
		}
	}

	return -1
}

/**
NC139 孩子们的游戏(圆圈中最后剩下的数)
 算法知识视频讲解
中等  通过率：32.39%  时间限制：1秒  空间限制：64M
知识点
数学
题目
题解(83)
讨论(865)
排行
描述
    每年六一儿童节，牛客都会准备一些小礼物和小游戏去看望孤儿院的孩子们。其中，有个游戏是这样的：首先，让 n 个小朋友们围成一个大圈，小朋友们的编号是0~n-1。然后，随机指定一个数 m ，让编号为0的小朋友开始报数。每次喊到 m-1 的那个小朋友要出列唱首歌，然后可以在礼品箱中任意的挑选礼物，并且不再回到圈中，从他的下一个小朋友开始，继续0... m-1报数....这样下去....直到剩下最后一个小朋友，可以不用表演，并且拿到牛客礼品，请你试着想下，哪个小朋友会得到这份礼品呢？

数据范围：1 \le n \le 50001≤n≤5000，1 \le m \le 100001≤m≤10000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
5,3
复制
返回值：
3
复制
示例2
输入：
2,3
复制
返回值：
1
复制
说明：
有2个小朋友编号为0，1，第一次报数报到3的是0号小朋友，0号小朋友出圈，1号小朋友得到礼物
示例3
输入：
10,17
复制
返回值：
2
*/
// func LastRemaining_Solution( n int ,  m int ) int {
//     // write code here
// }

/**
NC71 旋转数组的最小数字
 算法知识视频讲解
简单  通过率：34.07%  时间限制：1秒  空间限制：64M
知识点
二分
题目
题解(152)
讨论(2k)
排行
描述
有一个长度为 n 的非降序数组，比如[1,2,3,4,5]，将它进行旋转，即把一个数组最开始的若干个元素搬到数组的末尾，变成一个旋转数组，比如变成了[3,4,5,1,2]，或者[4,5,1,2,3]这样的。请问，给定这样一个旋转数组，求数组中的最小值。

数据范围：1 \le n \le 100001≤n≤10000，数组中任意元素的值: 0 \le val \le 100000≤val≤10000
要求：空间复杂度：O(1)O(1) ，时间复杂度：O(logn)O(logn)
示例1
输入：
[3,4,5,1,2]
复制
返回值：
1
复制
示例2
输入：
[3,100,200,3]
复制
返回值：
3
*/
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	lens := len(rotateArray)
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) >> 1
		if rotateArray[mid] < rotateArray[right] { // 处于后半部分的递增区，最小值在这之前
			right = mid
		} else if rotateArray[mid] > rotateArray[right] { // 处于旋转前半部分，最小值在这之后
			left = mid + 1
			continue
		} else { // 相等时，不确定是在左边还是右边，让right = right - 1慢慢缩小区间，同时不会错过答案
			right -= 1
		}
	}
	return rotateArray[left]
}

/**
NC74 数字在升序数组中出现的次数
 算法知识视频讲解
简单  通过率：32.43%  时间限制：1秒  空间限制：64M
知识点
数组
二分
题目
题解(119)
讨论(1k)
排行
描述
给定一个长度为 n 的非降序数组和一个非负数整数 k ，要求统计 k 在数组中出现的次数

数据范围：0 \le n \le 1000 , 0 \le k \le 1000≤n≤1000,0≤k≤100，数组中每个元素的值满足 0 \le val \le 1000≤val≤100
要求：空间复杂度 O(1)O(1)，时间复杂度 O(logn)O(logn)
示例1
输入：
[1,2,3,3,3,3,4,5],3
复制
返回值：
4
复制
示例2
输入：
[1,3,4,5],6
复制
返回值：
0
*/
func GetNumberOfK(data []int, k int) int {
	// write code here
	lens := len(data)
	if lens == 0 {
		return 0
	}

	num := 0
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) >> 1
		if data[mid] == k {
			num++
			for i := mid - 1; i >= 0; i-- {
				if data[i] == k {
					num++
				}
			}

			for i := mid + 1; i <= right; i++ {
				if data[i] == k {
					num++
				}
			}
			break
		}

		if data[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return num
}

/**
NC79 丑数
 算法知识视频讲解
中等  通过率：23.13%  时间限制：1秒  空间限制：64M
知识点
数学
二分
题目
题解(74)
讨论(717)
排行
描述
把只包含质因子2、3和5的数称作丑数（Ugly Number）。例如6、8都是丑数，但14不是，因为它包含质因子7。 习惯上我们把1当做是第一个丑数。求按从小到大的顺序的第 n个丑数。

数据范围：0 \le n \le 20000≤n≤2000
要求：空间复杂度 O(n)O(n) ， 时间复杂度 O(n)O(n)
示例1
输入：
7
复制
返回值：
8
*/
// func GetUglyNumber_Solution( index int ) int {
//     // write code here
// }

/**
NC84 完全二叉树结点数
 算法知识视频讲解
中等  通过率：54.51%  时间限制：1秒  空间限制：64M
知识点
树
二分
题目
题解(11)
讨论(32)
排行
描述
给定一棵完全二叉树的头节点head，返回这棵树的节点个数。

完全二叉树指：设二叉树的深度为h，则 [1,h-1] 层的节点数都满足 2^{i-1}2
i−1
 个

数据范围：节点数量满足 0 \le n \le 1000000≤n≤100000，节点上每个值都满足 0 \le val \le 1000000≤val≤100000
进阶：空间复杂度 O(1)O(1) ， 时间复杂度 O(n)O(n)
示例1
输入：
{1,2,3}
复制
返回值：
3
复制
示例2
输入：
{}
复制
返回值：
0
*/
// func nodeNum( head *TreeNode ) int {
//     // write code here
// }

/**
NC125 未排序数组中累加和为给定值的最长子数组长度
 算法知识视频讲解
简单  通过率：42.45%  时间限制：2秒  空间限制：256M
知识点
哈希
题目
题解(17)
讨论(23)
排行
描述
给定一个无序数组 arr , 其中元素可正、可负、可0。给定一个整数 k ，求 arr 所有连续子数组中累加和为k的最长连续子数组长度。保证至少存在一个合法的连续子数组。
[1,2,3]的连续子数组有[1,2]，[2,3]，[1,2,3] ，但是[1,3]不是

数据范围： 0 \le n \le 1000000≤n≤100000，|val| \le 100∣val∣≤100，|k| \le 1000000000∣k∣≤1000000000
进阶：空间复杂度 O(n)O(n) ， 时间复杂度 O(n)O(n)
示例1
输入：
[1,-2,1,1,1],0
复制
返回值：
3
复制
示例2
输入：
[0,1,2,3],3
复制
返回值：
3
复制
备注：
\begin{array}{l}1 \leqslant N \leqslant 10^{5} \\ -10^{9} \leqslant k \leqslant 10^{9} \\ -100 \leqslant a r r_{i} \leqslant 100\end{array}
1⩽N⩽10
5

−10
9
 ⩽k⩽10
9

−100⩽arr
i
​
 ⩽100
​*/
func maxlenEqualK(arr []int, k int) int {
	// write code here
	lens := len(arr)
	if lens == 0 {
		return 0
	}

	maps := make(map[int]int)
	maps[0] = -1
	max, sum := 0, 0
	for i := 0; i < lens; i++ {
		sum += arr[i]
		if v, ok := maps[sum-k]; ok { // 是否存在v...i子数组和为k，有则存储下最大长度
			max = maxInt(max, i-v)
		}
		if _, ok := maps[sum]; !ok { // 保存值为sum的第一个索引，这样计算出的长度是最大值
			maps[sum] = i
		}
	}

	return max
}

/**
NC130 分糖果问题
 算法知识视频讲解
中等  通过率：37.54%  时间限制：2秒  空间限制：256M
知识点
贪心
题目
题解(10)
讨论(21)
排行
描述
一群孩子做游戏，现在请你根据游戏得分来发糖果，要求如下：

1. 每个孩子不管得分多少，起码分到一个糖果。
2. 任意两个相邻的孩子之间，得分较多的孩子必须拿多一些糖果。(若相同则无此限制)

给定一个数组 arrarr 代表得分数组，请返回最少需要多少糖果。

要求: 时间复杂度为 O(n)O(n) 空间复杂度为 O(n)O(n)

数据范围： 1 \le n \le 1000001≤n≤100000 ，1 \le a_i \le 10001≤a
i
​
 ≤1000

示例1
输入：
[1,1,2]
复制
返回值：
4
复制
说明：
最优分配方案为1,1,2
示例2
输入：
[1,1,1]
复制
返回值：
3
复制
说明：
最优分配方案是1,1,1
*/
// func candy( arr []int ) int {
//     // write code here
// }

/**
NC104 比较版本号
 算法知识视频讲解
中等  通过率：32.40%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(22)
讨论(38)
排行
描述
牛客项目发布项目版本时会有版本号，比如1.02.11，2.14.4等等
现在给你2个版本号version1和version2，请你比较他们的大小
版本号是由修订号组成，修订号与修订号之间由一个"."连接。1个修订号可能有多位数字组成，修订号可能包含前导0，且是合法的。例如，1.02.11，0.1，0.2都是合法的版本号
每个版本号至少包含1个修订号。
修订号从左到右编号，下标从0开始，最左边的修订号下标为0，下一个修订号下标为1，以此类推。

比较规则：
一. 比较版本号时，请按从左到右的顺序依次比较它们的修订号。比较修订号时，只需比较忽略任何前导零后的整数值。比如"0.1"和"0.01"的版本号是相等的
二. 如果版本号没有指定某个下标处的修订号，则该修订号视为0。例如，"1.1"的版本号小于"1.1.1"。因为"1.1"的版本号相当于"1.1.0"，第3位修订号的下标为0，小于1
三.  version1 > version2 返回1，如果 version1 < version2 返回-1，不然返回0.

数据范围：len(version1),len(version2）\le 1000len(version1),len(version2）≤1000，版本号中每一节可能超过int的表达范围
进阶： 空间复杂度 O(1)O(1) ， 时间复杂度 O(n)O(n)
示例1
输入：
"1.1","2.1"
复制
返回值：
-1
复制
说明：
version1 中下标为 0 的修订号是 "1"，version2 中下标为 0 的修订号是 "2" 。1 < 2，所以 version1 < version2，返回-1

示例2
输入：
"1.1","1.01"
复制
返回值：
0
复制
说明：
version2忽略前导0，为"1.1"，和version相同，返回0
示例3
输入：
"1.1","1.1.1"
复制
返回值：
-1
复制
说明：
"1.1"的版本号小于"1.1.1"。因为"1.1"的版本号相当于"1.1.0"，第3位修订号的下标为0，小于1，所以version1 < version2，返回-1
示例4
输入：
"2.0.1","2"
复制
返回值：
1
复制
说明：
version1的下标2>version2的下标2，返回1
示例5
输入：
"0.226","0.36"
复制
返回值：
1
复制
说明：
226>36，version1的下标2>version2的下标2，返回1
*/
// func compare( version1 string ,  version2 string ) int {
//     // write code here
// }

/**
NC149 kmp算法
 算法知识视频讲解
中等  通过率：44.94%  时间限制：1秒  空间限制：256M
知识点
字符串
题目
题解(32)
讨论(33)
排行
描述
给你一个文本串 T ，一个非空模板串 S ，问 T 在 S 中出现了多少次

数据范围：1 \le len(S) \le 500000, 1 \le len(T) \le 10000001≤len(S)≤500000,1≤len(T)≤1000000
要求：空间复杂度 O(len(S))O(len(S))，时间复杂度 O(len(S)+len(T))O(len(S)+len(T))
示例1
输入：
"ababab","abababab"
复制
返回值：
2
复制
示例2
输入：
"abab","abacabab"
复制
返回值：
1
复制
备注：
空间O(n)时间O(n)的算法
*/
// func kmp( S string ,  T string ) int {
//     // write code here
// }

/**
NC159 最小生成树
 算法知识视频讲解
中等  通过率：42.40%  时间限制：1秒  空间限制：256M
知识点
图论
图
题目
题解(8)
讨论(20)
排行
描述
一个有 n 户人家的村庄，有 m 条路相互连接着。村里现在要修路，每条路都有一个成本价格，现在请你帮忙计算下，最少需要花费多少钱，就能让这 n 户人家连接起来。

costcost 为一个二维数组，每个元素是一个长度为 3 的一维数组 aa ， a[0]a[0] 和 a[1]a[1] 表示村庄 a[0]a[0] 和村庄 a[1]a[1] 有一条路，修这条路的成本价格为 a[2]a[2] .

每户之间可能有多条道路连接，但不可能自己与自己相连

数据范围: 1 \le n \le 5 \times 10^31≤n≤5×10
3
  ， 1 \le m \le 5 \times 10^51≤m≤5×10
5
  ， 1 \le a[2] \le 10^41≤a[2]≤10
4

进阶： 时间复杂度 O(n+mlogm)O(n+mlogm) ， 空间复杂度 O(n)O(n)
示例1
输入：
3,3,[[1,3,3],[1,2,1],[2,3,1]]
复制
返回值：
2
复制
示例2
输入：
2,1,[[1,2,1]]
复制
返回值：
1
*/
// func miniSpanningTree( n int ,  m int ,  cost [][]int ) int {
//     // write code here
// }

/**
NC23 划分链表
 算法知识视频讲解
中等  通过率：35.92%  时间限制：1秒  空间限制：64M
知识点
链表
双指针
题目
题解(22)
讨论(129)
排行
描述
给出一个长度为 n 的单链表和一个值 x ，单链表的每一个值为 listi ，请返回一个链表的头结点，要求新链表中小于 x 的节点全部在大于等于 x 的节点左侧，并且两个部分之内的节点之间与原来的链表要保持相对顺序不变。

例如：
给出 1 \to 4 \to 3 \to 2 \to 5 \to 21→4→3→2→5→2 和 x = 3x=3
返回 1 \to 2 \to 2 \to 4 \to 3 \to 51→2→2→4→3→5

数据范围： n \le 200n≤200 ， -100 \le list[i] \le 100−100≤list[i]≤100
进阶：时间复杂度 O(n)O(n) ， 空间复杂度 O(1)O(1)
示例1
输入：
{1,4,3,2,5,2},3
复制
返回值：
{1,2,2,4,3,5}
复制
示例2
输入：
{1,2,3,4,1},5
复制
返回值：
{1,2,3,4,1}
*/
// func partition( head *ListNode ,  x int ) *ListNode {
//     // write code here
// }

/**
NC75 数组中只出现一次的两个数字
 算法知识视频讲解
中等  通过率：55.86%  时间限制：1秒  空间限制：256M
知识点
位运算
哈希
题目
题解(72)
讨论(95)
排行
描述
一个整型数组里除了两个数字只出现一次，其他的数字都出现了两次。请写程序找出这两个只出现一次的数字。

数据范围：数组长度 2\le n \le 10002≤n≤1000，数组中每个数的大小 0 < val \le 10000000<val≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

提示：输出时按非降序排列。
示例1
输入：
[1,4,1,6]
复制
返回值：
[4,6]
复制
说明：
返回的结果中较小的数排在前面
示例2
输入：
[1,2,3,3,2,9]
复制
返回值：
[1,9]
*/
func FindNumsAppearOnce(array []int) []int {
	// write code here
	type1, type2 := 0, 0
	xorsum := 0 // = type1 ^ type2
	for _, num := range array {
		xorsum ^= num
	}

	lsb := xorsum & -xorsum // 可找出xorsum中最低位的1，有1则表示该位上type1和type2值不同，一个是0，一个是1
	for _, num := range array {
		if num&lsb > 0 { // lsb位为1的分类到type1中
			type1 ^= num
		} else {
			type2 ^= num
		}
	}

	if type1 > type2 {
		type1, type2 = type2, type1
	}
	return []int{type1, type2}
}

/**
NC106 三个数的最大乘积
 算法知识视频讲解
简单  通过率：35.48%  时间限制：1秒  空间限制：64M
知识点
数组
数学
题目
题解(27)
讨论(42)
排行
描述
给定一个长度为 nn 的无序数组 AA ，包含正数、负数和 0 ，请从中找出 3 个数，使得乘积最大，返回这个乘积。

要求时间复杂度： O(n)O(n) ，空间复杂度： O(1)O(1) 。

数据范围：
3 \le n \le 10^43≤n≤10
4

-10^4 \le A[i] \le 10^4−10
4
 ≤A[i]≤10
4

示例1
输入：
[3,4,1,2]
复制
返回值：
24
*/
func solveThreeMaxProduct(A []int) int64 {
	// write code here
	lens := len(A)
	if lens == 3 {
		return int64(A[0] * A[1] * A[2])
	}

	min_1, min_2 := math.MaxInt64, math.MaxInt64
	max_1, max_2, max_3 := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < lens; i++ {
		if A[i] < min_1 {
			min_1, min_2 = A[i], min_1
		} else if A[i] < min_2 {
			min_2 = A[i]
		}

		if A[i] > max_1 {
			max_1, max_2, max_3 = A[i], max_1, max_2
		} else if A[i] > max_2 {
			max_2, max_3 = A[i], max_2
		} else if A[i] > max_3 {
			max_3 = A[i]
		}
	}

	return int64(maxInt(min_1*min_2*max_1, max_1*max_2*max_3))
}

/**
NC114 旋转字符串
 算法知识视频讲解
简单  通过率：62.47%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(20)
讨论(32)
排行
描述
字符串旋转:
给定两字符串A和B，如果能将A从中间某个位置分割为左右两部分字符串（可以为空串），并将左边的字符串移动到右边字符串后面组成新的字符串可以变为字符串B时返回true。

例如：如果A=‘youzan’，B=‘zanyou’，A按‘you’‘zan’切割换位后得到‘zanyou’和B相同，返回true。
再如：如果A=‘abcd’，B=‘abcd’，A切成‘abcd’和''（空串），换位后可以得到B，返回true。

数据范围：A,B字符串长度满足 n \le 1000n≤1000，保证字符串中仅包含小写英文字母和阿拉伯数字
进阶： 时间复杂度 O(n)O(n)，空间复杂度 O(n)O(n)
示例1
输入：
"youzan","zanyou"
复制
返回值：
true
复制
示例2
输入：
"youzan","zyouan"
复制
返回值：
false
复制
示例3
输入：
"nowcoder","nowcoder"
复制
返回值：
true
*/
func solveRotateString(A string, B string) bool {
	// write code here
	if A == B {
		return true
	}
	if A == "" || B == "" || len(A) != len(B) {
		return false
	}

	lens := len(A)
	for i := 0; i < lens; i++ {
		tmpA := A[i+1:] + A[0:i+1]
		if tmpA == B {
			return true
		}
	}

	return false
}

/**
NC115 栈和排序
 算法知识视频讲解
中等  通过率：39.81%  时间限制：1秒  空间限制：128M
知识点
栈
排序
题目
题解(11)
讨论(12)
排行
描述
给你一个 1 到 n 的排列和一个栈，并按照排列顺序入栈

你要在不打乱入栈顺序的情况下，仅利用入栈和出栈两种操作，输出字典序最大的出栈序列

排列：指 1 到 n 每个数字出现且仅出现一次

数据范围：  1 \le n \le 5 \times 10^41≤n≤5×10
4
 ，排列中的值都满足 1 \le val \le n1≤val≤n

进阶：空间复杂度 O(n)O(n) ，时间复杂度 O(n)O(n)
示例1
输入：
[2,1,5,3,4]
复制
返回值：
[5,4,3,1,2]
复制
说明：
操作       栈     结果
2 入栈；[2]       []
1 入栈；[2\1]     []
5 入栈；[2\1\5]   []
5 出栈；[2\1]     [5]
3 入栈；[2\1\3]   [5]
4 入栈；[2\1\3\4] [5]
4 出栈；[2\1\3]   [5,4]
3 出栈；[2\1]     [5,4,3]
1 出栈；[2]       [5,4,3,1]
2 出栈；[]        [5,4,3,1,2]
示例2
输入：
[1,2,3,4,5]
复制
返回值：
[5,4,3,2,1]
*/
// func solve( a []int ) []int {
//     // write code here
// }

/**
NC145 01背包
 算法知识视频讲解
简单  通过率：46.77%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(17)
讨论(31)
排行
描述
已知一个背包最多能容纳体积之和为v的物品

现有 n 个物品，第 i 个物品的体积为 vi , 重量为 wi

求当前背包最多能装多大重量的物品?

数据范围： 1 \le v \le 10001≤v≤1000 ， 1 \le n \le 10001≤n≤1000 ， 1 \le v_i \le 10001≤v
i
​
 ≤1000 ， 1 \le w_i \le 10001≤w
i
​
 ≤1000

进阶 ：O(n \cdot v)O(n⋅v)
示例1
输入：
10,2,[[1,3],[10,4]]
复制
返回值：
4
复制
说明：
第一个物品的体积为1，重量为3，第二个物品的体积为10，重量为4。只取第二个物品可以达到最优方案，取物重量为4
示例2
输入：
10,2,[[1,3],[9,8]]
复制
返回值：
11
复制
说明：
两个物品体积之和等于背包能装的体积，所以两个物品都取是最优方案
*/
func knapsack2(V int, n int, vw [][]int) int {
	// write code here
	if V == 0 || n == 0 || len(vw) == 0 {
		return 0
	}

	// 动态规划标准套路
	// 一、明确“状态”和“选择”
	// 状态有两个，背包的容量和可选择的物品
	// 选择，对于每件物品，可选择“装进背包”或“不装背包”
	// 二、明确dp数组的定义
	// dp[i][w]定义：对于前i个物品，当前背包的容量为w，这种情况下能装的最大价值是dp[i][w]
	// 根据定义，我们要求的就是dp[n][V]。base case是dp[0][...]=dp[...][0]=0，因为没有物品或背包容量为空时，能装的最大价值就是0
	// 三、根据“选择”，思考状态转移的逻辑
	// 针对dp[i][w]的值有两种选择：
	// 1. 如果没有把第i个物品装入背包，那么最大价值会是dp[i-1][w]，继承相同容量下前i-1物品能装的最大价值
	// 2. 如果把第i个物品装入背包，那么最大价值会是dp[i-1][w-vw[i][0]] + vw[i][1]，前i-1物品在背包容量为w-vw[i][0]时的最大价值+第i个物品价值

	// 初始化dp数组
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, V+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= V; w++ {
			if w-vw[i-1][0] < 0 { // 装不下，则直接继承dp[i-1][w]
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = maxInt(dp[i-1][w], dp[i-1][w-vw[i-1][0]]+vw[i-1][1])
			}
		}
	}
	return dp[n][V]
}

/**
NC151 最大公约数
 算法知识视频讲解
入门  通过率：56.21%  时间限制：1秒  空间限制：256M
知识点
数学
题目
题解(44)
讨论(42)
排行
描述
如果有一个自然数 a 能被自然数 b 整除，则称 a 为 b 的倍数， b 为 a 的约数。几个自然数公有的约数，叫做这几个自然数的公约数。公约数中最大的一个公约数，称为这几个自然数的最大公约数。

输入 a 和 b , 请返回 a 和 b 的最大公约数。

数据范围：1 \le a,b \le 10^91≤a,b≤10
9

进阶：空间复杂度 O(1)O(1)，时间复杂度 O(logn)O(logn)
示例1
输入：
3,6
复制
返回值：
3
复制
示例2
输入：
8,12
复制
返回值：
4
复制
备注：
a和b的范围是[1-109]
*/
// func gcd( a int ,  b int ) int {
//     // write code here
// }

/**
NC156 数组中只出现一次的数（其它数出现k次）
 算法知识视频讲解
简单  通过率：55.66%  时间限制：1秒  空间限制：256M
知识点
位运算
题目
题解(44)
讨论(76)
排行
描述
给定一个长度为 n 的整型数组 arr 和一个整数 k(k>1) 。
已知 arr 中只有 1 个数出现一次，其他的数都出现 k 次。
请返回只出现了 1 次的数。

数据范围: 1 \le n \le 2*10^51≤n≤2∗10
5
  ， 1 \lt k \lt 1001<k<100 ， -2*10^9 \le arr[i] \le 2*10^9−2∗10
9
 ≤arr[i]≤2∗10
9

进阶：时间复杂度 O(32n)O(32n)，空间复杂度 O(1)O(1)


示例1
输入：
[5,4,1,1,5,1,5],3
复制
返回值：
4
复制
示例2
输入：
[2,2,1],2
复制
返回值：
1
*/
func foundOnceNumber(arr []int, k int) int {
	// write code here
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range arr {
			total += int32(num) >> i & 1
		}
		if total%int32(k) != 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

/**
NC67 汉诺塔问题
 算法知识视频讲解
中等  通过率：37.61%  时间限制：3秒  空间限制：32M
知识点
递归
动态规划
题目
题解(14)
讨论(44)
排行
描述
我们有由底至上为从大到小放置的 n 个圆盘，和三个柱子（分别为左/中/右即left/mid/right），开始时所有圆盘都放在左边的柱子上，按照汉诺塔游戏的要求我们要把所有的圆盘都移到右边的柱子上，要求一次只能移动一个圆盘，而且大的圆盘不可以放到小的上面。

请实现一个函数打印最优移动轨迹。

给定一个 `int n` ，表示有 n 个圆盘。请返回一个 `string` 数组，其中的元素依次为每次移动的描述。描述格式为： `move from [left/mid/right] to [left/mid/right]`。

数据范围：1\le n \le 161≤n≤16
要求：时间复杂度 O(3^n)O(3
n
 ) ， 空间复杂度 O(3^n)O(3
n
 )
示例1
输入：
2
复制
返回值：
["move from left to mid","move from left to right","move from mid to right"]
*/
// func getSolution( n int ) []string {
//     // write code here
// }

/**
NC85 拼接所有的字符串产生字典序最小的字符串
 算法知识视频讲解
中等  通过率：48.36%  时间限制：2秒  空间限制：256M
知识点
贪心
题目
题解(13)
讨论(6)
排行
描述
给定一个长度为 n 的字符串数组 strs ，请找到一种拼接顺序，使得数组中所有的字符串拼接起来组成的字符串是所有拼接方案中字典序最小的，并返回这个拼接后的字符串。

数据范围：1 \le n \le 10000001≤n≤1000000 ， 1 \le len(strs_i)\le 101≤len(strs
i
​
 )≤10
进阶：空间复杂度 O(n)O(n) ， 时间复杂度 O(nlognlen(strs_i))O(nlognlen(strs
i
​
 ))
示例1
输入：
["abc","de"]
复制
返回值：
"abcde"
复制
示例2
输入：
["a","a","b"]
复制
返回值：
"aab"
复制
备注：
1 \leq n \leq 10^51≤n≤10
5

1 \leq |strs_i| \leq 101≤∣strs
i
​
 ∣≤10
*/
//  func minString( strs []string ) string {
//     // write code here
// }

/**
NC160 二分查找-I
 算法知识视频讲解
简单  通过率：48.36%  时间限制：1秒  空间限制：256M
知识点
二分
题目
题解(16)
讨论(10)
排行
描述
请实现无重复数字的升序数组的二分查找

给定一个 元素升序的、无重复数字的整型数组 nums 和一个目标值 target ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标（下标从 0 开始），否则返回 -1

数据范围：0 \le len(nums) \le 100000≤len(nums)≤10000 ， 数组中任意值满足 |val| \le 10000∣val∣≤10000
进阶：时间复杂度 O(\log n)O(logn) ，空间复杂度 O(1)O(1)

示例1
输入：
[-1,0,3,4,6,10,13,14],13
复制
返回值：
6
复制
说明：
13 出现在nums中并且下标为 6
示例2
输入：
[],3
复制
返回值：
-1
复制
说明：
nums为空，返回-1
示例3
输入：
[-1,0,3,4,6,10,13,14],2
复制
返回值：
-1
复制
说明：
2 不存在nums中因此返回 -1
备注：
数组元素长度在[0,10000]之间
数组每个元素都在 [-9999, 9999]之间。
*/
func searchEasy(nums []int, target int) int {
	// write code here
	if len(nums) == 0 {
		return -1
	}

	lens := len(nums)
	left, right := 0, lens-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

/**
NC161 二叉树的中序遍历
 算法知识视频讲解
中等  通过率：66.49%  时间限制：1秒  空间限制：256M
知识点
树
iOS工程师
小米
2021
题目
题解(21)
讨论(20)
排行
描述
给定一个二叉树的根节点root，返回它的中序遍历结果。

数据范围：树上节点数满足 0 \le n \le 10000≤n≤1000，树上每个节点的值满足 0 \le val \le 10000≤val≤1000
进阶：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,#,#,3}
复制
返回值：
[2,3,1]
复制
说明：

示例2
输入：
{}
复制
返回值：
[]
复制
示例3
输入：
{1,2}
复制
返回值：
[2,1]
复制
说明：

示例4
输入：
{1,#,2}
复制
返回值：
[1,2]
复制
说明：

备注：
树中节点数目在范围 [0, 100] 内
树中的节点的值在[-100,100]以内
*/
func inorderTraversal(root *TreeNode) []int {
	// write code here
	var ret []int
	if root == nil {
		return ret
	}
	var dfsInOrder func(node *TreeNode)
	dfsInOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfsInOrder(node.Left)
		ret = append(ret, node.Val)
		dfsInOrder(node.Right)
	}
	dfsInOrder(root)
	return ret
}

/**
NC147 主持人调度
 算法知识视频讲解
中等  通过率：32.82%  时间限制：2秒  空间限制：256M
知识点
贪心
题目
题解(17)
讨论(9)
排行
描述
有 n 个活动即将举办，每个活动都有开始时间与活动的结束时间，第 i 个活动的开始时间是 starti ,第 i 个活动的结束时间是 endi ,举办某个活动就需要为该活动准备一个活动主持人。

一位活动主持人在同一时间只能参与一个活动。并且活动主持人需要全程参与活动，换句话说，一个主持人参与了第 i 个活动，那么该主持人在 (starti,endi) 这个时间段不能参与其他任何活动。求为了成功举办这 n 个活动，最少需要多少名主持人。

数据范围: 1 \le n \le 10^51≤n≤10
5
  ， -2^{32} \le start_i,end_i \le 2^{31}-1−2
32
 ≤start
i
​
 ,end
i
​
 ≤2
31
 −1

复杂度要求：时间复杂度 O(n \log n)O(nlogn) ，空间复杂度  O(n)O(n)
示例1
输入：
2,[[1,2],[2,3]]
复制
返回值：
1
复制
说明：
只需要一个主持人就能成功举办这两个活动
示例2
输入：
2,[[1,3],[2,4]]
复制
返回值：
2
复制
说明：
需要两个主持人才能成功举办这两个活动
备注：
1 \leq n \leq 10^51≤n≤10
5

start_i,end_istart
i
​
 ,end
i
​
 在int范围内
*/
//  func minmumNumberOfHost( n int ,  startEnd [][]int ) int {
//     // write code here
// }

/**
NC148 几步可以从头跳到尾
 算法知识视频讲解
简单  通过率：38.49%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(14)
讨论(5)
排行
描述
给你一个长度为 n 的数组 a。
ai 表示从 i 这个位置开始最多能往后跳多少格。
求从 1 开始最少需要跳几次就能到达第 n 个格子。

数据范围： 1 \le n \le 1000001≤n≤100000 ， 1 \le a_i \le 10000000001≤a
i
​
 ≤1000000000
进阶： 空间复杂度O(1)O(1) ， 时间复杂度 O(n)O(n)
示例1
输入：
2,[1,2]
复制
返回值：
1
复制
说明：
从1号格子只需要跳跃一次就能到达2号格子
示例2
输入：
3,[2,3,1]
复制
返回值：
1
复制
说明：
从1号格子只需要跳一次就能直接抵达3号格子
*/
func Jump(n int, A []int) int {
	// write code here
	lens := len(A)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < lens-1; i++ {
		maxPosition = maxInt(maxPosition, i+A[i]) // 从0开始，一步步确认下一个跳跃点（能跳到的最远的距离）
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

/**
NC150 二叉树的个数
 算法知识视频讲解
中等  通过率：30.92%  时间限制：1秒  空间限制：256M
知识点
树
题目
题解(8)
讨论(11)
排行
描述
已知一棵节点个数为 n 的二叉树的中序遍历单调递增, 求该二叉树能能有多少种树形, 输出答案对 109 取模

数据范围：1 \le n \le 30001≤n≤3000
进阶：空间复杂度 O(1)O(1) ， 时间复杂度 O(nlogn)O(nlogn)
示例1
输入：
1
复制
返回值：
1
复制
示例2
输入：
2
复制
返回值：
2
复制
示例3
输入：
4
复制
返回值：
14
*/
// func numberOfTree( n int ) int {
//     // write code here
// }

/**
NC152 数的划分
 算法知识视频讲解
中等  通过率：45.38%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(10)
讨论(5)
排行
描述
将整数 n 分成 k 份，且每份不能为空，任意两个方案不能相同(不考虑顺序)。
例如： n=7,k=3 ，下面三种分法被认为是相同的。

1，1，5;1，1，5;
1，5，1;1，5，1;
5，1，1;5，1，1;

问有多少种不同的分法, 答案对 109 + 7 取模。

数据范围： 6 \lt n \le 5000，2 \le k \le 10006<n≤5000，2≤k≤1000
进阶：空间复杂度 O(nk)O(nk) ，时间复杂度 O(nk)O(nk)
示例1
输入：
7,3
复制
返回值：
4
复制
示例2
输入：
6,2
复制
返回值：
3
*/
// func divideNumber( n int ,  k int ) int {
//     // write code here
// }

/**
NC153 信封嵌套问题
 算法知识视频讲解
中等  通过率：58.94%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(10)
讨论(15)
排行
描述
给 n 个信封的长度和宽度。如果信封 a 的长和宽都小于信封 b ，那么信封 a 可以放到信封 b 里，请求出信封最多可以嵌套多少层。

数据范围： 1 \le n \le 1001≤n≤100 ， 1 \le letters[i][0], letters[i][1] \le 10^31≤letters[i][0],letters[i][1]≤10
3

要求：空间复杂度 O(n)O(n) ，时间复杂度 O(n^2)O(n
2
 )
进阶：空间复杂度 O(n)O(n) ，时间复杂度 O(nlogn)O(nlogn)
示例1
输入：
[[3,4],[2,3],[4,5],[1,3],[2,2],[3,6],[1,2],[3,2],[2,4]]
复制
返回值：
4
复制
说明：
从里到外分别是{1，2}，{2，3}，{3，4}，{4，5}。
示例2
输入：
[[1,4],[4,1]]
复制
返回值：
1
复制
备注：
时间复杂度O(nlog n)，空间复杂度O(n)。
*/
// func maxLetters( letters [][]int ) int {
//     // write code here
// }

/**
NC154 最长回文子序列
 算法知识视频讲解
中等  通过率：50.49%  时间限制：3秒  空间限制：256M
知识点
动态规划
题目
题解(10)
讨论(10)
排行
描述
给定一个字符串，找到其中最长的回文子序列，并返回该序列的长度。

注：回文序列是指这个序列无论从左读还是从右读都是一样的。
        本题中子序列字符串任意位置删除k（len(s)>=k>=0)个字符后留下的子串。

数据范围：字符串长度满足 1 \le n \le 50001≤n≤5000
进阶：空间复杂度 O(n^2)O(n
2
 ) ， 时间复杂度 O(n^2)O(n
2
 )
示例1
输入：
"abccsb"
复制
返回值：
4
复制
说明：
分别选取第2、3、4、6位上的字符组成“bccb”子序列是最优解
示例2
输入：
"abcdewa"
复制
返回值：
2
复制
说明：
分别选取第一个和最后一个是最优解
*/
// func longestPalindromeSubSeq( s string ) int {
//     // write code here
// }

/**
NC155 最长严格上升子数组(二)
 算法知识视频讲解
中等  通过率：33.62%  时间限制：1秒  空间限制：256M
知识点
贪心
动态规划
穷举
题目
题解(1)
讨论(2)
排行
描述
给定一个长度为n的正整数数组nums，可以任意改变数组的其中一个元素，然后返回nums的最长"严格上升"子数组的长度。
1.子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
2.严格上升指在数组上任意位置都满足 nums[i] < nums[i+1]，比如[1,2,2,3]，其中[1,2,2]不是严格上升的子数组，[1,2]是的
数据范围：  ，
要求： 空间复杂度 ，时间复杂度
示例1
输入：
[7,2,3,1,5,6]
复制
返回值：
5
复制
说明：
将1改为4，最长严格上升子数组为[2,3,4,5,6]
示例2
输入：
[1,2,3,4]
复制
返回值：
4
复制
说明：
最长严格上升子数组为[1,2,3,4]
示例3
输入：
[1,2,2,3]
复制
返回值：
3
复制
说明：
改变一个元素之后，最长严格上升子数组为[1,2,3]或者[2,3,4]，长度都为3
*/
// func maxSubArrayLengthTwo( nums []int ) int {
//     // write code here
// }

/**
NC157 单调栈
 算法知识视频讲解
中等  通过率：63.68%  时间限制：1秒  空间限制：256M
知识点
栈
题目
题解(21)
讨论(23)
排行
描述
给定一个长度为 n 的可能含有重复值的数组 arr ，找到每一个 i 位置左边和右边离 i 位置最近且值比 arri 小的位置。

请设计算法，返回一个二维数组，表示所有位置相应的信息。位置信息包括：两个数字 l 和 r，如果不存在，则值为 -1，下标从 0 开始。

数据范围：1 \le n \le 10^51≤n≤10
5
  ， -10^9 \le arr[i] \le 10^9−10
9
 ≤arr[i]≤10
9

进阶：空间复杂度 O(n)O(n) ，时间复杂度 O(n)O(n)

示例1
输入：
[3,4,1,5,6,2,7]
复制
返回值：
[[-1,2],[0,2],[-1,-1],[2,5],[3,5],[2,-1],[5,-1]]
复制
示例2
输入：
[1,1,1,1]
复制
返回值：
[[-1,-1],[-1,-1],[-1,-1],[-1,-1]]
*/
// func foundMonotoneStack( nums []int ) [][]int {
//     // write code here
// }

/**
NC158 单源最短路
 算法知识视频讲解
中等  通过率：38.07%  时间限制：1秒  空间限制：256M
知识点
图
题目
题解(12)
讨论(17)
排行
描述
在一个有 n 个点， m 个边的有向图中，已知每条边长，求出 1 到 n 的最短路径，返回 1 到 n 的最短路径值。如果 1 无法到 n ，输出 -1

图中可能有重边，无自环。

数据范围：1 < n \le 5001<n≤500 ， 1 \le m \le 50001≤m≤5000 ， 1 \le dist(n, m) \le 10001≤dist(n,m)≤1000
示例1
输入：
5,5,[[1,2,2],[1,4,5],[2,3,3],[3,5,4],[4,5,5]]
复制
返回值：
9
复制
示例2
输入：
2,1,[[1,2,4]]
复制
返回值：
4
复制
备注：
两个整数n和m,表示图的顶点数和边数。
一个二维数组，一维3个数据，表示顶点到另外一个顶点的边长度是多少
每条边的长度范围[0,1000]。
注意数据中可能有重边
*/
// func findShortestPath( n int ,  m int ,  graph [][]int ) int {
//     // write code here
// }
