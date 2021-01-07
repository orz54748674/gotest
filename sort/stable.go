package _sort

import "sort"

type stable struct {
	arr []int
}

func (s stable) Len() int { return len(s.arr) }
func (s stable) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}
func (s stable) Less(i, j int) bool { return s.arr[i] < s.arr[j] }

func Stable(arr []int) {
	s := new(stable)
	s.arr = arr
	sort.Stable(s)
}
