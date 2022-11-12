package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
   bool

   string

   int  int8  int16  int32  int64
   uint uint8 uint16 uint32 uint64 uintptr

   byte // uint8 的别名

   rune // int32 的别名
       // 表示一个 Unicode 码点

   float32 float64

   complex64 complex128

   int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	i := 42
	f := float64(i)
	u := uint(f)
	fmt.Println(u)

	//当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了
	v := 42
	fmt.Printf("v 的类型 %T\n", v)

	//常量 常量不能用 := 语法声明
	const (
		Pi    = 3.14
		World = "世界"
	)

	fmt.Println(Pi, World)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum1 := 1
	//for ; sum1 < 1000; {
	for sum1 < 1000 {
		sum1 += sum1
	}
	fmt.Println(sum1)

	//if 语句可以在条件表达式前执行一个简单的语句。 该语句声明的变量作用域仅在 if 之内。
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	//defer 语句会将函数推迟到外层函数返回之后执行。
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	//推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照 后进先出 的顺序调用。
	defer fmt.Print("world ")
	defer fmt.Print("new ")
	fmt.Print("hello ")
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
