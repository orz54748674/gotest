package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s1 := []int32{1, 2, 3}
	//s2 := []int32{1, 2}
	//fmt.Printf("s1&s2 is equal:%v\n", reflect.DeepEqual(s1, s2))
	//m1 := map[int]int{1: 2}
	//m2 := map[int]int{1: 2, 2: 0}
	//fmt.Printf("m1&m2 is equal:%v\n", reflect.DeepEqual(m1, m2))

	rate := float64(1789)/float64(1386) - 1
	fmt.Println(rate, "=>", ParsePercent(rate, 4))
}

//保留小数点后n位
func ParseFloatN(rate float64, n int) float64 {
	formatStr := fmt.Sprintf("%%.%df", n)
	if newRate, err := strconv.ParseFloat(fmt.Sprintf(formatStr, rate), 64); err == nil {
		return newRate
	}
	return rate
}

//输出百分比且保留小数点后n位
func ParsePercent(rate float64, n int) string {
	return fmt.Sprintf("%v%%", ParseFloatN(rate*100, n))
}
