package sutil

import (
	"fmt"
	"math"
	"sort"
)

// 定义一个结构体，用于封装一个整数切片
type IntSliceWrapper struct {
	data []int
}

// 获取切片的长度
func (s *IntSliceWrapper) Length() int {
	return len(s.data)
}

// 添加元素到切片末尾
func (s *IntSliceWrapper) Add(item int) {
	s.data = append(s.data, item)
}

// 根据索引删除元素
func (s *IntSliceWrapper) Remove(index int) error {
	if index < 0 || index >= len(s.data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	// 使用切片操作删除指定索引的元素
	s.data = append(s.data[:index], s.data[index+1:]...)
	return nil
}

// 修改指定索引处的元素
func (s *IntSliceWrapper) Update(index int, newItem int) error {
	if index < 0 || index >= len(s.data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	s.data[index] = newItem
	return nil
}

// 查询指定索引处的元素
func (s *IntSliceWrapper) Get(index int) (int, error) {
	if index < 0 || index >= len(s.data) {
		return 0, fmt.Errorf("index out of range: %d", index)
	}
	return s.data[index], nil
}

// 打印切片内容
func (s *IntSliceWrapper) Print() {
	fmt.Println("Slice contents:", s.data)
}

// 清空切片
func (s *IntSliceWrapper) Clear() {
	s.data = []int{}
}

// 检查元素是否存在
func (s *IntSliceWrapper) Contains(item int) bool {
	for _, v := range s.data {
		if v == item {
			return true
		}
	}
	return false
}

// 插入元素到指定位置
func (s *IntSliceWrapper) Insert(index int, item int) error {
	if index < 0 || index > len(s.data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	s.data = append(s.data[:index], append([]int{item}, s.data[index:]...)...)
	return nil
}

// 反转切片
func (s *IntSliceWrapper) Reverse() {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
}

// SortAscending 对切片进行排序（升序）
func (s *IntSliceWrapper) SortAscending() {
	sort.Ints(s.data)
}

func (s IntSliceWrapper) SortedAscending() IntSliceWrapper {
	ret := make([]int, len(s.data))
	copy(ret, s.data)
	sort.Ints(ret)
	return IntSliceWrapper{
		data: ret,
	}
}

// SortDescending 对切片进行排序（降序）
func (s *IntSliceWrapper) SortDescending() {
	sort.Sort(sort.Reverse(sort.IntSlice(s.data)))
}

// SortedDescending return the sorted result, and it does not affect the original slice
func (s IntSliceWrapper) SortedDescending() IntSliceWrapper {
	ret := make([]int, len(s.data))
	copy(ret, s.data)
	sort.Sort(sort.Reverse(sort.IntSlice(ret)))
	return IntSliceWrapper{
		data: ret,
	}
}

// Filter 过滤切片（保留满足条件的元素）
func (s *IntSliceWrapper) Filter(predicate func(int) bool) {
	filtered := []int{}
	for _, v := range s.data {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	s.data = filtered
}

// Filtered return the filtered result, and this method does not affect the original slice
func (s IntSliceWrapper) Filtered(f func(int) int) IntSliceWrapper {
	ret := make([]int, len(s.data))
	for idx, v := range s.data {
		ret[idx] = f(v)
	}
	return IntSliceWrapper{
		data: ret,
	}
}

// Sum 计算切片中所有元素的和
func (s *IntSliceWrapper) Sum() int {
	sum := 0
	for _, v := range s.data {
		sum += v
	}
	return sum
}

// Average 计算切片中所有元素的平均值
func (s *IntSliceWrapper) Average() float64 {
	if len(s.data) == 0 {
		return 0
	}
	return float64(s.Sum()) / float64(len(s.data))
}

// Max 查找切片中的最大值
func (s *IntSliceWrapper) Max() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	max := s.data[0]
	for _, v := range s.data {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// Min 查找切片中的最小值
func (s *IntSliceWrapper) Min() (int, error) {
	if len(s.data) == 0 {
		return 0, fmt.Errorf("slice is empty")
	}
	min := s.data[0]
	for _, v := range s.data {
		if v < min {
			min = v
		}
	}
	return min, nil
}

// StdDev 计算切片中所有元素的标准差
func (s *IntSliceWrapper) StdDev() float64 {
	if len(s.data) == 0 {
		return 0
	}
	mean := s.Average()
	variance := 0.0
	for _, v := range s.data {
		variance += math.Pow(float64(v)-mean, 2)
	}
	return math.Sqrt(variance / float64(len(s.data)))
}

func (s IntSliceWrapper) Find(predicate func(int) bool) (int, bool) {
	for _, v := range s.data {
		if predicate(v) {
			return v, true
		}
	}
	return -1, false
}


