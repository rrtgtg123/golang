package main

import (
	p1 "code/package01"
	"code/package01/package001"
	_ "fmt"
)

func main() {
	//package xxx 随意命名 ，包名
	//import 内前缀 相当于重命名包名
	p1.Test02()
	p001.Test03()
}
