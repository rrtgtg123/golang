package main

import (
	"fmt"
	"math"
)
import "math/rand"

func main() {
	//分配内存空间，赋0值
	var i *int = new(int)

	//address in RAM
	fmt.Println(i)
	//value
	fmt.Println(*i)

	//切片 切片类型，切片容量，数组容量
	list := make([]int, 5, 10)
	fmt.Println(len(list), cap(list))
	//不能超越切片长度赋值，如需，需要重新切片
	//list[5] = 443
	list[4] = 443
	list1 := list[:10]
	fmt.Println(len(list1), cap(list1))

	//map 初始容量 map扩容机制？
	mp := make(map[string]int, 0)
	//通道 通道缓冲能力
	ch := make(chan int64, 1)

	fmt.Println(list, mp, ch)

	//不建议使用
	println(rand.Intn(100))

	fmt.Println(math.Pi)

	fmt.Println(add(995, 443))

	fmt.Println(split(18))

}

// 返回类型为any？
func add(a int, b int) int {
	return a + b
}

// go的返回值可以被命名，它们会被视作定义在函数顶部的变量
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
