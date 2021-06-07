package main

import "fmt"

const tar int = 10

func main(){
	for i := 1; i<tar;  i++{
		for j := 1; j<=i;j++{
			fmt.Printf("%d*%d=%d\t", i, j, i*j)
		}
		println()
	}
}

// 格式化输出：fmt.Sprintf(格式化样式，参数列表~)