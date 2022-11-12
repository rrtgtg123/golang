package main

import "fmt"

var x, y int
var ( // 这种因式分解关键字的写法一般用于声明全局变量
	a int
	b bool
)

const (
	i = 1 << iota
	j = 3 << iota
	//不写会默认复制上一行的运算逻辑
	k
	l
)

var c, d = 1, 2
var e, f = 123, "hello"

func main() {
	fmt.Println("hello world!")

	// %d 表示整型数字，%s 表示字符串
	var stockCode = 123
	var endDate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var targetUrl = fmt.Sprintf(url, stockCode, endDate)
	fmt.Println(targetUrl)

	// := 声明变量的快捷方式,这种不带声明格式的只能在函数体中出现
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)

	_, numb, str4 := numbers()
	//只获取函数返回值的后两个
	fmt.Println(numb, str4)

	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var char rune = 66
	fmt.Println(string(char))

	var inf interface{} = "100"
	//断言？ 接口的强转与int类型不一样
	i2, err := inf.(int)
	fmt.Println(i2, err)
}

// 一个可以返回多个值的函数
func numbers() (int, int, string) {
	a, b, c := 1, 2, "str"
	return a, b, c
}
