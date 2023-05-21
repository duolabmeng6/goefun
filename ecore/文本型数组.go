package ecore

import (
	"bytes"
	"encoding/json"
	"github.com/duolabmeng6/goefun/src/rwmutex"
	"math"
	"sort"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
)

type StrArray struct {
	mu    *rwmutex.RWMutex
	Array []string
}

// NewStrArray creates and returns an empty Array.
// The parameter <safe> is used to specify whether using Array in concurrent-safety,
// which is false in default.
func NewStrArray(safe ...bool) *StrArray {
	return NewStrArraySize(0, 0, safe...)
}

// NewStrArraySize create and returns an Array with given size and cap.
// The parameter <safe> is used to specify whether using Array in concurrent-safety,
// which is false in default.
func NewStrArraySize(size int, cap int, safe ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: make([]string, size, cap),
	}
}

// NewStrArrayFrom creates and returns an Array with given slice <Array>.
// The parameter <safe> is used to specify whether using Array in concurrent-safety,
// which is false in default.
func NewStrArrayFrom(Array []string, safe ...bool) *StrArray {
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: Array,
	}
}

// NewStrArrayFromCopy creates and returns an Array from a copy of given slice <Array>.
// The parameter <safe> is used to specify whether using Array in concurrent-safety,
// which is false in default.
func NewStrArrayFromCopy(Array []string, safe ...bool) *StrArray {
	newArray := make([]string, len(Array))
	copy(newArray, Array)
	return &StrArray{
		mu:    rwmutex.New(safe...),
		Array: newArray,
	}
}

// Get returns the value of the specified index,
// the caller should notice the boundary of the Array.
func (a *StrArray) Get(index int) string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	value := a.Array[index]
	return value
}

// Set sets value to specified index.
func (a *StrArray) Set(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Array[index] = value
	return a
}

// SetArray sets the underlying slice Array with the given <Array>.
func (a *StrArray) SetArray(Array []string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Array = Array
	return a
}

// Replace replaces the Array items by given <Array> from the beginning of Array.
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

// Sum returns the sum of values in an Array.
func (a *StrArray) Sum() (sum int) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.Array {
		sum += gconv.Int(v)
	}
	return
}

// Sort sorts the Array in increasing order.
// The parameter <reverse> controls whether sort
// in increasing order(default) or decreasing order
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

// SortFunc sorts the Array by custom function <less>.
func (a *StrArray) SortFunc(less func(v1, v2 string) bool) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	sort.Slice(a.Array, func(i, j int) bool {
		return less(a.Array[i], a.Array[j])
	})
	return a
}

// InsertBefore inserts the <value> to the front of <index>.
func (a *StrArray) InsertBefore(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.Array[index:]...)
	a.Array = append(a.Array[0:index], value)
	a.Array = append(a.Array, rear...)
	return a
}

// InsertAfter inserts the <value> to the back of <index>.
func (a *StrArray) InsertAfter(index int, value string) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	rear := append([]string{}, a.Array[index+1:]...)
	a.Array = append(a.Array[0:index+1], value)
	a.Array = append(a.Array, rear...)
	return a
}

// Remove removes an item by index.
func (a *StrArray) Remove(index int) string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if index < 0 || index >= len(a.Array) {
		return ""
	}
	// Determine Array boundaries when deleting to improve deletion efficiency。
	if index == 0 {
		value := a.Array[0]
		a.Array = a.Array[1:]
		return value
	} else if index == len(a.Array)-1 {
		value := a.Array[index]
		a.Array = a.Array[:index]
		return value
	}
	// If it is a non-boundary delete,
	// it will involve the creation of an Array,
	// then the deletion is less efficient.
	value := a.Array[index]
	a.Array = append(a.Array[:index], a.Array[index+1:]...)
	return value
}

// RemoveValue removes an item by value.
// It returns true if value is found in the Array, or else false if not found.
func (a *StrArray) RemoveValue(value string) bool {
	if i := a.Search(value); i != -1 {
		a.Remove(i)
		return true
	}
	return false
}

// PushLeft pushes one or multiple items to the beginning of Array.
func (a *StrArray) PushLeft(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(value, a.Array...)
	a.mu.Unlock()
	return a
}

// PushRight pushes one or multiple items to the end of Array.
// It equals to Append.
func (a *StrArray) PushRight(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(a.Array, value...)
	a.mu.Unlock()
	return a
}

// PopLeft pops and returns an item from the beginning of Array.
func (a *StrArray) PopLeft() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	value := a.Array[0]
	a.Array = a.Array[1:]
	return value
}

// PopRight pops and returns an item from the end of Array.
func (a *StrArray) PopRight() string {
	a.mu.Lock()
	defer a.mu.Unlock()
	index := len(a.Array) - 1
	value := a.Array[index]
	a.Array = a.Array[:index]
	return value
}

// PopRand randomly pops and return an item out of Array.
func (a *StrArray) PopRand() string {
	return a.Remove(grand.Intn(len(a.Array)))
}

// PopRands randomly pops and returns <size> items out of Array.
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

// PopLefts pops and returns <size> items from the beginning of Array.
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

// PopRights pops and returns <size> items from the end of Array.
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

// Range picks and returns items by range, like Array[start:end].
// Notice, if in concurrent-safe usage, it returns a copy of slice;
// else a pointer to the underlying data.
//
// If <end> is negative, then the offset will start from the end of Array.
// If <end> is omitted, then the sequence will have everything from start up
// until the end of the Array.
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

// SubSlice returns a slice of elements from the Array as specified
// by the <offset> and <size> parameters.
// If in concurrent safe usage, it returns a copy of the slice; else a pointer.
//
// If offset is non-negative, the sequence will start at that offset in the Array.
// If offset is negative, the sequence will start that far from the end of the Array.
//
// If length is given and is positive, then the sequence will have up to that many elements in it.
// If the Array is shorter than the length, then only the available Array elements will be present.
// If length is given and is negative then the sequence will stop that many elements from the end of the Array.
// If it is omitted, then the sequence will have everything from offset up until the end of the Array.
//
// Any possibility crossing the left border of Array, it will fail.
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

// See PushRight.
func (a *StrArray) Append(value ...string) *StrArray {
	a.mu.Lock()
	a.Array = append(a.Array, value...)
	a.mu.Unlock()
	return a
}

// Len returns the length of Array.
func (a *StrArray) Len() int {
	a.mu.RLock()
	length := len(a.Array)
	a.mu.RUnlock()
	return length
}

// Slice returns the underlying data of Array.
// Note that, if it's in concurrent-safe usage, it returns a copy of underlying data,
// or else a pointer to the underlying data.
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

// Interfaces returns current Array as []interface{}.
func (a *StrArray) Interfaces() []interface{} {
	a.mu.RLock()
	defer a.mu.RUnlock()
	Array := make([]interface{}, len(a.Array))
	for k, v := range a.Array {
		Array[k] = v
	}
	return Array
}

// Clone returns a new Array, which is a copy of current Array.
func (a *StrArray) Clone() (newArray *StrArray) {
	a.mu.RLock()
	Array := make([]string, len(a.Array))
	copy(Array, a.Array)
	a.mu.RUnlock()
	return NewStrArrayFrom(Array, !a.mu.IsSafe())
}

// Clear deletes all items of current Array.
func (a *StrArray) Clear() *StrArray {
	a.mu.Lock()
	if len(a.Array) > 0 {
		a.Array = make([]string, 0)
	}
	a.mu.Unlock()
	return a
}

// Contains checks whether a value exists in the Array.
func (a *StrArray) Contains(value string) bool {
	return a.Search(value) != -1
}

// Search searches Array by <value>, returns the index of <value>,
// or returns -1 if not exists.
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

// Unique uniques the Array, clear repeated items.
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

// LockFunc locks writing by callback function <f>.
func (a *StrArray) LockFunc(f func(Array []string)) *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	f(a.Array)
	return a
}

// RLockFunc locks reading by callback function <f>.
func (a *StrArray) RLockFunc(f func(Array []string)) *StrArray {
	a.mu.RLock()
	defer a.mu.RUnlock()
	f(a.Array)
	return a
}

// Fill fills an Array with num entries of the value <value>,
// keys starting at the <startIndex> parameter.
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

// Chunk splits an Array into multiple Arrays,
// the size of each Array is determined by <size>.
// The last chunk may contain less than size elements.
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

// Pad pads Array to the specified length with <value>.
// If size is positive then the Array is padded on the right, or negative on the left.
// If the absolute value of <size> is less than or equal to the length of the Array
// then no padding takes place.
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

// Rand randomly returns one item from Array(no deleting).
func (a *StrArray) Rand() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.Array[grand.Intn(len(a.Array))]
}

// Rands randomly returns <size> items from Array(no deleting).
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

// Shuffle randomly shuffles the Array.
func (a *StrArray) Shuffle() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, v := range grand.Perm(len(a.Array)) {
		a.Array[i], a.Array[v] = a.Array[v], a.Array[i]
	}
	return a
}

// Reverse makes Array with elements in reverse order.
func (a *StrArray) Reverse() *StrArray {
	a.mu.Lock()
	defer a.mu.Unlock()
	for i, j := 0, len(a.Array)-1; i < j; i, j = i+1, j-1 {
		a.Array[i], a.Array[j] = a.Array[j], a.Array[i]
	}
	return a
}

// Join joins Array elements with a string <glue>.
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

// CountValues counts the number of occurrences of all values in the Array.
func (a *StrArray) CountValues() map[string]int {
	m := make(map[string]int)
	a.mu.RLock()
	defer a.mu.RUnlock()
	for _, v := range a.Array {
		m[v]++
	}
	return m
}

// Iterator is alias of IteratorAsc.
func (a *StrArray) Iterator(f func(k int, v string) bool) {
	a.IteratorAsc(f)
}

// IteratorAsc iterates the Array in ascending order with given callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (a *StrArray) IteratorAsc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for k, v := range a.Array {
		if !f(k, v) {
			break
		}
	}
}

// IteratorDesc iterates the Array in descending order with given callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (a *StrArray) IteratorDesc(f func(k int, v string) bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	for i := len(a.Array) - 1; i >= 0; i-- {
		if !f(i, a.Array[i]) {
			break
		}
	}
}

// String returns current Array as a string, which implements like json.Marshal does.
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

// MarshalJSON implements the interface MarshalJSON for json.Marshal.
func (a *StrArray) MarshalJSON() ([]byte, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return json.Marshal(a.Array)
}

// UnmarshalJSON implements the interface UnmarshalJSON for json.Unmarshal.
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

// UnmarshalValue is an interface implement which sets any type of value for Array.
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

func New文本型数组() *StrArray {
	return NewStrArraySize(0, 0, []bool{true}...)
}
func (a *StrArray) E加入成员(s string) {
	a.Append(s)
}
func (a *StrArray) E取数组成员数() int {
	return a.Len()
}

func (a *StrArray) E删除成员(欲删除的位置 int) {
	a.Remove(欲删除的位置)
}

func (a *StrArray) E清除数组() {
	a.Clear()
}

// E数组排序
// 对指定数值数组变量内的所有数组成员进行快速排序，不影响数组的维定义信息，排序结果存放回该数组变量。本命令为初级命令。
// 参数<1>的名称为“排序方向是否为从小到大”，类型为“逻辑型（bool）”，可以被省略。如果参数值为真，排序方向为从小到大，否则为从大到小。如果本参数被省略，默认值为真。
func (a *StrArray) E数组排序(排序方向是否为从小到大 bool) {
	a.Sort(排序方向是否为从小到大)
}

// E插入成员
// 参数<1>的名称为“欲插入的位置”，类型为“整数型（int）”。位置值从 1 开始，如果小于 1 或大于第一个参数所指定数组变量的成员数目 + 1，将不会插入任何数据。
// 参数<2>的名称为“欲插入的成员数据”，类型为“通用型（all）”，提供参数数据时可以同时提供数组或非数组数据。参数值所指定成员数据的数据类型必须能够与第一个参数所指定的数组变量相匹配。
func (a *StrArray) E插入成员(欲插入的位置 int, 欲插入的成员数据 string) {
	a.InsertBefore(欲插入的位置, 欲插入的成员数据)
}

func (a *StrArray) E取值(i int) string {
	return a.Get(i)
}
