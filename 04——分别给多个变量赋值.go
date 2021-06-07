package main

import (
	"fmt"
)

func main(){
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"		//分别给a，b，c，变量赋值（多重赋值）

	area = LENGTH * WIDTH
	fmt.Println("面积为：", area)
	println()
	println(a, b, c)
}