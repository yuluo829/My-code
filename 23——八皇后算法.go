package main		//八皇后算法Go实现		//code缺失，以后修改

import (
	"fmt"
)

func putn(i, j, n int){
	pos[i], b[j], c[j-i+7], d[i+j]=j, n, n, n
}

func checkPos(i, j int) int {
	if b[j] == 1 || c[j-i+7] == 1 || d[i+j] == 1 {
		return 0
	}
	return 1
}

func printQueen(n int){
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			if pos[i] == j{
				fmt.Print(1)
			}else{
				fmt.Print(0)
			}
		}
		fmt.Println()
	}
	fmt.Println("++++++++++++++++++++++++++")
}

func queen(i, n int, count *int){
	if i > 7{
		*count++
		printQueen(n)
		return
	}
	for j := 0; j < n;j ++{
		if checkPos(i, j) == 1{
			putn(i, j, 1)
			queen(i+1, n, count)
			putn(i, j, 0)
		}
	}
}

func main(){
	n, count := 8, 0
	queen(o, n, &count)
	fmt.Println(count)
}