package sutil

import (
	"fmt"
	"sort"
	"strings"
)

type SliceWrapper struct {
	Data []string
}

func (s *SliceWrapper) Length() int {
	return len(s.Data)
}

func (s *SliceWrapper) Add(item string) {
	s.Data = append(s.Data, item)
}

func (s *SliceWrapper) Remove(index int) error {
	if index < 0 || index >= len(s.Data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	s.Data = append(s.Data[:index], s.Data[index+1:]...)
	return nil
}

func (s *SliceWrapper) Update(index int, newItem string) error {
	if index < 0 || index >= len(s.Data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	s.Data[index] = newItem
	return nil
}

func (s *SliceWrapper) Get(index int) (string, error) {
	if index < 0 || index >= len(s.Data) {
		return "", fmt.Errorf("index out of range: %d", index)
	}
	return s.Data[index], nil
}

func (s *SliceWrapper) Clear() {
	s.Data = []string{}
}

func (s *SliceWrapper) Contains(item string) bool {
	for _, v := range s.Data {
		if v == item {
			return true
		}
	}
	return false
}

func (s *SliceWrapper) Insert(index int, item string) error {
	if index < 0 || index > len(s.Data) {
		return fmt.Errorf("index out of range: %d", index)
	}
	s.Data = append(s.Data[:index], append([]string{item}, s.Data[index:]...)...)
	return nil
}

func (s *SliceWrapper) Reverse() {
	for i, j := 0, len(s.Data)-1; i < j; i, j = i+1, j-1 {
		s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
	}
}

// Sort in-place sort
func (s *SliceWrapper) Sort() {
	sort.Strings(s.Data)
}
func (s *SliceWrapper) Sorted() SliceWrapper {
	ret := make([]string, len(s.Data))

	copy(ret, s.Data)
	sort.Strings(ret)

	return SliceWrapper{
		Data: ret,
	}
}

// Filter 原地过滤
func (s *SliceWrapper) Filter(predicate func(string) bool) {
	filtered := []string{}
	for _, v := range s.Data {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	s.Data = filtered
}

// Filtered 返回过滤的结果
func (s *SliceWrapper) Filtered(predicate func(string) bool) SliceWrapper {
	filtered := []string{}
	for _, v := range s.Data {
		if predicate(v) {
			filtered = append(filtered, v)
		}
	}
	return SliceWrapper{Data: filtered}
}

// IndexOf return the index of specified value
func (s *SliceWrapper) IndexOf(item string) int {
	for i, v := range s.Data {
		if v == item {
			return i
		}
	}
	return -1
}

// Unique 去重
func (s *SliceWrapper) Unique() {
	seen := make(map[string]bool)
	result := []string{}
	for _, v := range s.Data {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	s.Data = result
}

func (s *SliceWrapper) Join(separator string) string {
	return strings.Join(s.Data, separator)
}

// Map 将每一个item map为 新值
func (s *SliceWrapper) Map(mapper func(string) string) {
	for idx, v := range s.Data {
		s.Data[idx] = mapper(v)
	}
}

// Mapped return a new slice that are mapped by mapper
func (s *SliceWrapper) Mapped(mapper func(string) string) SliceWrapper {
	ret := make([]string, len(s.Data))
	for idx, v := range s.Data {
		ret[idx] = mapper(v)
	}
	return SliceWrapper{Data: ret}
}

func (s SliceWrapper) All(fn func(string) bool) bool {
	for _, v := range s.Data {
		if !fn(v) {
			return false
		}
	}
	return true
}
func (s SliceWrapper) None(fn func(string) bool) bool {
	for _, v := range s.Data {
		//满足了反而不是None
		if fn(v) {
			return false
		}
	}
	return true
}

func (s SliceWrapper) Any(fn func(string) bool) bool {
	for _, v := range s.Data {
		if fn(v) {
			return true
		}
	}
	return false
}

func (s SliceWrapper) Find(predicate func(string) bool) (string, bool) {
	for _, v := range s.Data {
		if predicate(v) {
			return v, true
		}
	}
	return "", false
}
