package main

import (
	"bytes"
	"fmt"
)

type ArrayInterface interface {
	// 指定位置添加
	Add(int, interface{})
	// 头插
	AddFirst(interface{})
	// 尾插
	AddLast(interface{})
	// 删除指定位置
	Remove(int) interface{}
	// 删除头部
	RemoveFirst() interface{}
	// 删除尾部
	RemoveLast() interface{}
	// 查找第一个索引
	Find(interface{}) int
	// 查找所有索引
	FindAll(interface{}) []int
	// 查找是否含有某个元素
	Contains(interface{}) bool
	// 获取指定索引的值
	Get(int) interface{}
	// 修改指定索引的值
	Set(int, interface{})
	// 容量
	Cap() int
	// 当前大小
	Len() int
	// 是否为空
	IsEmpty() bool
}

type Array struct {
	data []interface{}
	size int
}

// new一个动态数组
func NewArray(capacity int) *Array {
	arr := &Array{}
	arr.data = make([]interface{}, capacity)
	arr.size = 0
	return arr
}

// 指定位置添加值
func (a *Array) Add(idx int, val interface{}) {
	if idx < 0 || idx >= a.Cap() {
		panic("Add failed, index is illegal")
	}

	for i := a.size - 1; i > idx; i-- { // idx往后移1位
		a.data[i] = a.data[i-1]
	}
	a.data[idx] = val
	a.size++
	if a.size == a.Cap() { // size和数组长度一致时，需要扩容
		a.rezise(2 * a.Cap())
	}
}

// 头插
func (a *Array) AddFirst(val interface{}) {
	a.Add(0, val)
}

// 尾插
func (a *Array) AddLast(val interface{}) {
	a.Add(a.size, val)
}

// 删除指定位置的值
func (a *Array) Remove(idx int) interface{} {
	if idx < 0 || idx >= a.size {
		panic("Remove failed, index is illegal")
	}

	removeEle := a.data[idx]
	for i := idx + 1; i < a.size; i++ { // i+1值赋值给i
		a.data[i-1] = a.data[i]
	}
	a.size--             // 长度减少
	a.data[a.size] = nil // 清理最后一个元素（防止指针）
	if a.size == a.Cap()/4 && a.Cap()/2 > 0 {
		a.rezise(a.Cap() / 2)
	}
	return removeEle
}

// 头删
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 尾删
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 查找元素返回第一个索引
func (a *Array) Find(val interface{}) int {
	for i := 0; i < a.size; i++ {
		if val == a.data[i] {
			return i
		}
	}
	return -1
}

// 查找元素返回所有索引
func (a *Array) FindAll(val interface{}) (indexs []int) {
	for i := 0; i < a.size; i++ {
		if val == a.data[i] {
			indexs = append(indexs, i)
		}
	}
	return
}

// 校验数组是否包含特定值
func (a *Array) Contains(val interface{}) bool {
	if a.Find(val) == -1 {
		return false
	}
	return true
}

// 获取指定索引的值
func (a *Array) Get(idx int) interface{} {
	if idx < 0 || idx >= a.size {
		panic("Get failed, index is illegal.")
	}
	return a.data[idx]
}

// 设置指定索引的值
func (a *Array) Set(idx int, val interface{}) {
	if idx < 0 || idx >= a.size {
		panic("Set failed, index is illegal.")
	}
	a.data[idx] = val
}

// 获取容量
func (a *Array) Cap() int {
	return len(a.data)
}

// 获取当前数组长度
func (a *Array) Len() int {
	return a.size
}

// 数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 扩缩容
func (a *Array) rezise(newCapacity int) {
	newdata := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newdata[i] = a.data[i]
	}
	a.data = newdata
}

func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.Len(), a.Cap()))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}

var _ ArrayInterface = (*Array)(nil) // 确保Array实现了ArrayInterface接口

func main() {
	array := NewArray(5)
	dump(array)
	array.Add(2, 2)
	array.AddFirst(1)
	array.AddFirst(0)
	array.AddLast(3)
	dump(array)
	array.AddLast(4)
	dump(array)
	array.AddLast(5)
	dump(array)
	array.Remove(2)
	dump(array)
	array.RemoveLast()
	dump(array)
	array.RemoveFirst()
	dump(array)
	array.RemoveFirst()
	dump(array)
	array.RemoveFirst()
	dump(array)
	fmt.Println(array.Contains(4))
	fmt.Println(array.RemoveLast())

	/*
		Array: size = 0, capacity = 5
		[]
		Array: size = 4, capacity = 5
		[0,1,2,3]
		Array: size = 5, capacity = 10
		[0,1,2,3,4]
		Array: size = 6, capacity = 10
		[0,1,2,3,4,5]
		Array: size = 5, capacity = 10
		[0,1,3,4,5]
		Array: size = 4, capacity = 10
		[0,1,3,4]
		Array: size = 3, capacity = 10
		[1,3,4]
		Array: size = 2, capacity = 5
		[3,4]
		Array: size = 1, capacity = 2
		[4]
		true
		4
	*/
}

func dump(a *Array) {
	fmt.Println(a)
}
