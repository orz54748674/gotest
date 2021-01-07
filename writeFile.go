package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

const FilePath = "./test.txt" //写入数组的txt文件路径

func main() {
	writeFile()
	arr := readFile()
	fmt.Println(arr)
}

//写入数组至txt文件
func writeFile() {
	f, _ := os.Create(FilePath)
	_, _ = io.WriteString(f, "[")
	var arr []int
	for i := 0; i < 10000; i++ {
		arr = append(arr, i+1)
	}
	for k, v := range arr {
		_, _ = io.WriteString(f, strconv.Itoa(v))
		if k != len(arr)-1 {
			_, _ = io.WriteString(f, ",")
		}
	}
	_, _ = io.WriteString(f, "]")
	_ = f.Close()
}

//读取txt文件并转为数组返回
func readFile() (arr []int) {
	f, _ := os.Open(FilePath)
	content, _ := ioutil.ReadAll(f)
	_ = json.Unmarshal(content, &arr)
	_ = f.Close()
	return
}
