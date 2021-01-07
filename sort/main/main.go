package main

import (
	"fmt"
	"gotest/sort"
	"math/rand"
	"time"
)

/*
10万个打乱的数组
冒泡排序: 15.6591668s
选择排序: 5.5382233s
希尔排序: 2.6768141s
插入排序: 1.6206671s
归并排序: 1.2855309s
sort.Stable排序: 54.8234ms
普通堆排序: 11.9741ms
自底向上堆排序: 10.971ms
*/

const TestNumCount = 100000 //数组元素个数

var arr []int

//测试排序算法的运算效率
func main() {
	for i := 0; i < TestNumCount; i++ {
		arr = append(arr, i+1)
	}
	var cost time.Duration
	cost = costTime(_sort.BubbleSort)
	fmt.Printf("冒泡排序: %v\n", cost)
	cost = costTime(_sort.SelectGoodSort)
	fmt.Printf("选择排序: %v\n", cost)
	cost = costTime(_sort.ShellSort)
	fmt.Printf("希尔排序: %v\n", cost)
	cost = costTime(_sort.InsertSort)
	fmt.Printf("插入排序: %v\n", cost)
	cost = costTimeMerge(_sort.MergeSort3)
	fmt.Printf("归并排序: %v\n", cost)
	cost = costTime(_sort.Stable)
	fmt.Printf("sort.Stable排序: %v\n", cost)
	cost = costTimeHeap()
	fmt.Printf("普通堆排序: %v\n", cost)
	cost = costTime(_sort.HeapSort)
	fmt.Printf("自底向上堆排序: %v\n", cost)
}

func costTime(f func([]int)) time.Duration {
	shuffle(arr)
	bT := time.Now()
	f(arr)
	eT := time.Since(bT)
	return eT
}

func costTimeMerge(f func([]int, int)) time.Duration {
	shuffle(arr)
	bT := time.Now()
	f(arr, TestNumCount)
	eT := time.Since(bT)
	return eT
}

func costTimeHeap() time.Duration {
	shuffle(arr)
	bT := time.Now()
	// 构建最大堆
	h := _sort.NewHeap(arr)
	for _, v := range arr {
		h.Push(v)
	}
	// 将堆元素移除
	for range arr {
		h.Pop()
	}
	eT := time.Since(bT)
	return eT
}

//随机打乱数组
func shuffle(s []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(s) > 0 {
		n := len(s)
		randIndex := r.Intn(n)
		s[n-1], s[randIndex] = s[randIndex], s[n-1]
		s = s[:n-1]
	}
}
