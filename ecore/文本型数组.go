// Copyright 2023 The duolabmeng6 Authors. All rights reserved.
// license that can be found in the LICENSE file.

// Package ecore 提供文本型数组相关功能
package ecore

import (
	"bytes"
	"encoding/json"
	"math"
	"sort"
	"strings"

	"github.com/duolabmeng6/goefun/src/rwmutex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

// StrArray 文本型数组结构体，提供并发安全的字符串数组操作
type StrArray struct {
	mu    *rwmutex.RWMutex
	Array []string
}

// NewStrArray 创建并返回一个空的文本型数组。
//
// 参数：
//   - safe: 可选参数，是否启用并发安全模式，默认为false
//
// 返回值：
//   - *StrArray: 新创建的文本型数组
func NewStrArray(safe ...bool) *StrArray {
	return NewStrArraySize(0, 0, safe...)
}

// NewStrArraySize 创建并返回一个指定大小和容量的文本型数组。
//
// 参数：
//   - size: 数组的初始大小
//   - cap: 数组的初始容量
//   - safe: 可选参数，是否启用并发安全模式，默认为false
//
// 返回值：
//   - *StrArray: 新创建的文本型数组
func NewStrArraySize(size int, cap int, safe ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: make([]string, size, cap),
	}
}

// NewStrArrayFrom 从现有切片创建并返回一个文本型数组。
//
// 参数：
//   - Array: 字符串切片
//   - safe: 可选参数，是否启用并发安全模式，默认为false
//
// 返回值：
//   - *StrArray: 新创建的文本型数组
func NewStrArrayFrom(Array []string, safe ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: Array,
	}
}

// NewStrArrayFromCopy 从现有切片的副本创建并返回一个文本型数组。
//
// 参数：
//   - Array: 字符串切片
//   - safe: 可选参数，是否启用并发安全模式，默认为false
//
// 返回值：
//   - *StrArray: 新创建的文本型数组
func NewStrArrayFromCopy(Array []string, safe ...bool) *StrArray {
	newArray := make([]string, len(Array))
	copy(newArray, Array)
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: newArray,
	}
}

// Get 获取指定索引位置的元素值。
//
// 参数：
//   - index: 索引位置
//
// 返回值：
//   - string: 指定位置的元素值
func (a *StrArray) Get(index int) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.Array[index]
	return value
}

// Set 设置指定索引位置的元素值。
//
// 参数：
//   - index: 索引位置
//   - value: 要设置的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Set(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Array[index] = value
	return a
}

// SetArray 设置底层数组。
//
// 参数：
//   - Array: 要设置的字符串切片
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) SetArray(Array []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Array = Array
	return a
}

// Replace 从数组开头替换元素。
//
// 参数：
//   - Array: 用于替换的字符串切片
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Replace(Array []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	max := len(Array)
	if max > len(a.Array) {
		max = len(a.Array)
	}
	for i := 0; i < max; i++ {
		a.Array[i] = Array[i]
	}
	return a
}

// Sum 计算数组中所有元素转换为整数后的总和。
//
// 返回值：
//   - int: 元素总和
func (a *StrArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.Array {
		sum += gconv.Int(v)
	}
	return
}

// Sort 对数组进行排序。
//
// 参数：
//   - reverse: 可选参数，是否降序排列，默认为升序
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Sort(reverse ...bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(reverse) > 0 && reverse[0] {
		sort.Slice(a.Array, func(i, j int) bool {
			if strings.Compare(a.Array[i], a.Array[j]) < 0 {
				return false
			}
			return true
		})
	} else {
		sort.Strings(a.Array)
	}
	return a
}

// SortFunc 使用自定义函数对数组进行排序。
//
// 参数：
//   - less: 自定义比较函数，返回true表示v1应该排在v2前面
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) SortFunc(less func(v1, v2 string) bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.Array, func(i, j int) bool {
		return less(a.Array[i], a.Array[j])
	})
	return a
}

// InsertBefore 在指定索引位置之前插入元素。
//
// 参数：
//   - index: 索引位置
//   - value: 要插入的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) InsertBefore(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.Array[index:]...)
	a.Array = append(a.Array[0:index], value)
	a.Array = append(a.Array, rear...)
	return a
}

// InsertAfter 在指定索引位置之后插入元素。
//
// 参数：
//   - index: 索引位置
//   - value: 要插入的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) InsertAfter(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.Array[index+1:]...)
	a.Array = append(a.Array[0:index+1], value)
	a.Array = append(a.Array, rear...)
	return a
}

// Remove 删除指定索引位置的元素。
//
// 参数：
//   - index: 索引位置
//
// 返回值：
//   - string: 被删除的元素值
func (a *StrArray) Remove(index int) string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.Array) {
		return ""
	}
	// 确定数组边界时删除以提高删除效率
	if index == 0 {
		value := a.Array[0]
		a.Array = a.Array[1:]
		return value
	} else if index == len(a.Array)-1 {
		value := a.Array[index]
		a.Array = a.Array[:index]
		return value
	}
	// 如果是非边界删除，会涉及数组的创建，删除效率较低
	value := a.Array[index]
	a.Array = append(a.Array[:index], a.Array[index+1:]...)
	return value
}

// RemoveValue 根据值删除元素。
//
// 参数：
//   - value: 要删除的值
//
// 返回值：
//   - bool: 如果找到并删除成功返回true，否则返回false
func (a *StrArray) RemoveValue(value string) bool {
	if i := a.Search(value); i != -1 {
		a.Remove(i)
		return true
	}
	return false
}

// PushLeft 在数组开头压入一个或多个元素。
//
// 参数：
//   - value: 要压入的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) PushLeft(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(value, a.Array...)
	a.mu.Unlock()
	return a
}

// PushRight 在数组末尾压入一个或多个元素，等同于Append。
//
// 参数：
//   - value: 要压入的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) PushRight(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(a.Array, value...)
	a.mu.Unlock()
	return a
}

// PopLeft 弹出并返回数组开头的元素。
//
// 返回值：
//   - string: 弹出的元素值
func (a *StrArray) PopLeft() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.Array[0]
	a.Array = a.Array[1:]
	return value
}

// PopRight 弹出并返回数组末尾的元素。
//
// 返回值：
//   - string: 弹出的元素值
func (a *StrArray) PopRight() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.Array) - 1
	value := a.Array[index]
	a.Array = a.Array[:index]
	return value
}

// PopRand 随机弹出并返回一个元素。
//
// 返回值：
//   - string: 弹出的元素值
func (a *StrArray) PopRand() string {
	return a.Remove(grand.Intn(len(a.Array)))
}

// PopRands 随机弹出并返回指定数量的元素。
//
// 参数：
//   - size: 要弹出的元素数量
//
// 返回值：
//   - []string: 弹出的元素切片
func (a *StrArray) PopRands(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size > len(a.Array) {
		size = len(a.Array)
	}
	Array := make([]string, size)
	for i := 0; i < size; i++ {
		index := grand.Intn(len(a.Array))
		Array[i] = a.Array[index]
		a.Array = append(a.Array[:index], a.Array[index+1:]...)
	}
	return Array
}

// PopLefts 从数组开头弹出指定数量的元素。
//
// 参数：
//   - size: 要弹出的元素数量
//
// 返回值：
//   - []string: 弹出的元素切片
func (a *StrArray) PopLefts(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	length := len(a.Array)
	if size > length {
		size = length
	}
	value := a.Array[0:size]
	a.Array = a.Array[size:]
	return value
}

// PopRights 从数组末尾弹出指定数量的元素。
//
// 参数：
//   - size: 要弹出的元素数量
//
// 返回值：
//   - []string: 弹出的元素切片
func (a *StrArray) PopRights(size int) []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.Array) - size
	if index < 0 {
		index = 0
	}
	value := a.Array[index:]
	a.Array = a.Array[:index]
	return value
}

// Range 按范围获取元素，类似于数组[start:end]。
//
// 参数：
//   - start: 起始索引
//   - end: 结束索引（可选）
//
// 返回值：
//   - []string: 范围内的元素切片
func (a *StrArray) Range(start int, end ...int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	offsetEnd := len(a.Array)
	if len(end) > 0 && end[0] < offsetEnd {
		offsetEnd = end[0]
	}
	if start > offsetEnd {
		return nil
	}
	if start < 0 {
		start = 0
	}
	Array := ([]string)(nil)
	if a.mu.IsSafe() {
		Array = make([]string, offsetEnd-start)
		copy(Array, a.Array[start:offsetEnd])
	} else {
		Array = a.Array[start:offsetEnd]
	}
	return Array
}

// SubSlice 获取子切片。
//
// 参数：
//   - offset: 偏移量
//   - length: 长度（可选）
//
// 返回值：
//   - []string: 子切片
func (a *StrArray) SubSlice(offset int, length ...int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	size := len(a.Array)
	if len(length) > 0 {
		size = length[0]
	}
	if offset > len(a.Array) {
		return nil
	}
	if offset < 0 {
		offset = len(a.Array) + offset
		if offset < 0 {
			return nil
		}
	}
	if size < 0 {
		offset += size
		size = -size
		if offset < 0 {
			return nil
		}
	}
	end := offset + size
	if end > len(a.Array) {
		end = len(a.Array)
		size = len(a.Array) - offset
	}
	if a.mu.IsSafe() {
		s := make([]string, size)
		copy(s, a.Array[offset:])
		return s
	} else {
		return a.Array[offset:end]
	}
}

// Append 在数组末尾追加元素，等同于PushRight。
//
// 参数：
//   - value: 要追加的值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Append(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(a.Array, value...)
	a.mu.Unlock()
	return a
}

// Len 返回数组的长度。
//
// 返回值：
//   - int: 数组长度
func (a *StrArray) Len() int {
	a.mu.RLock()
	length := len(a.Array)
	a.mu.RUnlock()
	return length
}

// Slice 返回底层数组切片。
//
// 返回值：
//   - []string: 数组切片
func (a *StrArray) Slice() []string {
	Array := ([]string)(nil)
	if a.mu.IsSafe() {
		a.mu.RLock()
		defer a.mu.RUnlock()
		Array = make([]string, len(a.Array))
		copy(Array, a.Array)
	} else {
		Array = a.Array
	}
	return Array
}

// Interfaces 将数组转换为[]interface{}类型。
//
// 返回值：
//   - []interface{}: 转换后的切片
func (a *StrArray) Interfaces() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	Array := make([]interface{}, len(a.Array))
	for k, v := range a.Array {
		Array[k] = v
	}
	return Array
}

// Clone 返回数组的副本。
//
// 返回值：
//   - *StrArray: 新的数组副本
func (a *StrArray) Clone() (newArray *StrArray) {
	a.mu.RLock()
	Array := make([]string, len(a.Array))
	copy(Array, a.Array)
	a.mu.RUnlock()
	return NewStrArrayFrom(Array, !a.mu.IsSafe())
}

// Clear 清空数组中的所有元素。
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Clear() *StrArray {
	a.mu.Lock()
	if len(a.Array) > 0 {
		a.Array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains 检查数组中是否包含指定值。
//
// 参数：
//   - value: 要检查的值
//
// 返回值：
//   - bool: 如果包含返回true，否则返回false
func (a *StrArray) Contains(value string) bool {
	return a.Search(value) != -1
}

// Search 在数组中搜索指定值。
//
// 参数：
//   - value: 要搜索的值
//
// 返回值：
//   - int: 找到返回索引位置，未找到返回-1
func (a *StrArray) Search(value string) int {
	if len(a.Array) == 0 {
		return -1
	}
	a.mu.RLock()
	result := -1
	for index, v := range a.Array {
		if strings.Compare(v, value) == 0 {
			result = index
			break
		}
	}
	a.mu.RUnlock()
	return result
}

// Unique 去除数组中的重复元素。
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Unique() *StrArray {
	a.mu.Lock()
	for i := 0; i < len(a.Array)-1; i++ {
		for j := i + 1; j < len(a.Array); j++ {
			if a.Array[i] == a.Array[j] {
				a.Array = append(a.Array[:j], a.Array[j+1:]...)
			}
		}
	}
	a.mu.Unlock()
	return a
}

// LockFunc 使用写锁执行回调函数。
//
// 参数：
//   - f: 回调函数
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) LockFunc(f func(Array []string)) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.Array)
	return a
}

// RLockFunc 使用读锁执行回调函数。
//
// 参数：
//   - f: 回调函数
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) RLockFunc(f func(Array []string)) *StrArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.Array)
	return a
}

// Fill 用指定值填充数组的指定范围。
//
// 参数：
//   - startIndex: 起始索引
//   - num: 填充数量
//   - value: 填充值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Fill(startIndex int, num int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if startIndex < 0 {
		startIndex = 0
	}
	for i := startIndex; i < startIndex+num; i++ {
		if i > len(a.Array)-1 {
			a.Array = append(a.Array, value)
		} else {
			a.Array[i] = value
		}
	}
	return a
}

// Chunk 将数组分割为多个指定大小的子数组。
//
// 参数：
//   - size: 每个子数组的大小
//
// 返回值：
//   - [][]string: 分割后的二维数组
func (a *StrArray) Chunk(size int) [][]string {
	if size < 1 {
		return nil
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	length := len(a.Array)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, a.Array[i*size:end])
		i++
	}
	return n
}

// Pad 使用指定值将数组填充到指定长度。
//
// 参数：
//   - size: 目标长度，正数向右填充，负数向左填充
//   - value: 填充值
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Pad(size int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	if size == 0 || (size > 0 && size < len(a.Array)) || (size < 0 && size > -len(a.Array)) {
		return a
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(a.Array)
	tmp := make([]string, n)
	for i := 0; i < n; i++ {
		tmp[i] = value
	}
	if size > 0 {
		a.Array = append(a.Array, tmp...)
	} else {
		a.Array = append(tmp, a.Array...)
	}
	return a
}

// Rand 随机返回一个元素（不删除）。
//
// 返回值：
//   - string: 随机元素
func (a *StrArray) Rand() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Array[grand.Intn(len(a.Array))]
}

// Rands 随机返回指定数量的元素（不删除）。
//
// 参数：
//   - size: 返回数量
//
// 返回值：
//   - []string: 随机元素切片
func (a *StrArray) Rands(size int) []string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if size > len(a.Array) {
		size = len(a.Array)
	}
	n := make([]string, size)
	for i, v := range grand.Perm(len(a.Array)) {
		n[i] = a.Array[v]
		if i == size-1 {
			break
		}
	}
	return n
}

// Shuffle 随机打乱数组元素的顺序。
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Shuffle() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.Perm(len(a.Array)) {
		a.Array[i], a.Array[v] = a.Array[v], a.Array[i]
	}
	return a
}

// Reverse 反转数组元素的顺序。
//
// 返回值：
//   - *StrArray: 返回数组本身以支持链式调用
func (a *StrArray) Reverse() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.Array)-1; i < j; i, j = i+1, j-1 {
		a.Array[i], a.Array[j] = a.Array[j], a.Array[i]
	}
	return a
}

// Join 使用指定分隔符连接数组元素。
//
// 参数：
//   - glue: 分隔符
//
// 返回值：
//   - string: 连接后的字符串
func (a *StrArray) Join(glue string) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	for k, v := range a.Array {
		buffer.WriteString(v)
		if k != len(a.Array)-1 {
			buffer.WriteString(glue)
		}
	}
	return buffer.String()
}

// CountValues 统计每个值在数组中出现的次数。
//
// 返回值：
//   - map[string]int: 值到出现次数的映射
func (a *StrArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.Array {
		m[v]++
	}
	return m
}

// Iterator 迭代数组（升序迭代器的别名）。
//
// 参数：
//   - f: 回调函数，返回false停止迭代
func (a *StrArray) Iterator(f func(k int, v string) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc 升序迭代数组。
//
// 参数：
//   - f: 回调函数，返回false停止迭代
func (a *StrArray) IteratorAsc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.Array {
		if !f(k, v) {
			break
		}
	}
}

// IteratorDesc 降序迭代数组。
//
// 参数：
//   - f: 回调函数，返回false停止迭代
func (a *StrArray) IteratorDesc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.Array) - 1; i >= 0; i-- {
		if !f(i, a.Array[i]) {
			break
		}
	}
}

// String 返回数组的字符串表示。
//
// 返回值：
//   - string: 数组的JSON格式字符串
func (a *StrArray) String() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	buffer := bytes.NewBuffer(nil)
	buffer.WriteByte('[')
	for k, v := range a.Array {
		buffer.WriteString(`"` + gstr.QuoteMeta(v, `"\`) + `"`)
		if k != len(a.Array)-1 {
			buffer.WriteByte(',')
		}
	}
	buffer.WriteByte(']')
	return buffer.String()
}

// MarshalJSON 实现json.Marshaler接口。
func (a *StrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.Array)
}

// UnmarshalJSON 实现json.Unmarshaler接口。
func (a *StrArray) UnmarshalJSON(b []byte) error {
	if a.mu == nil {
		a.mu = rwmutex.New()
		a.Array = make([]string, 0)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if err := json.Unmarshal(b, &a.Array); err != nil {
		return err
	}
	return nil
}

// UnmarshalValue 设置任意类型的值。
func (a *StrArray) UnmarshalValue(value interface{}) error {
	if a.mu == nil {
		a.mu = rwmutex.New()
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	switch value.(type) {
	case string, []byte:
		return json.Unmarshal(gconv.Bytes(value), &a.Array)
	default:
		a.Array = gconv.SliceStr(value)
	}
	return nil
}

// New文本型数组 创建一个新的文本型数组（并发安全）。
//
// 返回值：
//   - *StrArray: 新创建的文本型数组
func New文本型数组() *StrArray {
	return NewStrArraySize(0, 0, []bool{true}...)
}

// E加入成员 向数组末尾添加一个成员。
//
// 参数：
//   - s: 要添加的字符串
func (a *StrArray) E加入成员(s string) {
	a.Append(s)
}

// E取数组成员数 获取数组的成员数量。
//
// 返回值：
//   - int: 成员数量
func (a *StrArray) E取数组成员数() int {
	return a.Len()
}

// E删除成员 删除指定位置的成员。
//
// 参数：
//   - 欲删除的位置: 要删除的成员索引
func (a *StrArray) E删除成员(欲删除的位置 int) {
	a.Remove(欲删除的位置)
}

// E清除数组 清除数组中的所有成员。
func (a *StrArray) E清除数组() {
	a.Clear()
}

// E数组排序 对数组进行排序。
//
// 参数：
//   - 排序方向是否为从小到大: true为升序，false为降序
func (a *StrArray) E数组排序(排序方向是否为从小到大 bool) {
	a.Sort(排序方向是否为从小到大)
}

// E插入成员 在指定位置插入成员。
//
// 参数：
//   - 欲插入的位置: 插入位置（从0开始）
//   - 欲插入的成员数据: 要插入的字符串
func (a *StrArray) E插入成员(欲插入的位置 int, 欲插入的成员数据 string) {
	a.InsertBefore(欲插入的位置, 欲插入的成员数据)
}

// E取值 获取指定位置的成员值。
//
// 参数：
//   - i: 索引位置
//
// 返回值：
//   - string: 该位置的成员值
func (a *StrArray) E取值(i int) string {
	return a.Get(i)
}
