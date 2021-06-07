package main

import(
	"fmt"
)

func main(){
	var totallever int
	fmt.Print("请输入三角形的行数：")
	fmt.Scanln(&totallever)
	for i := 1; i <= totallever; i++{
		// var k int = 1
		for k := 1; k <= totallever - i; k++{
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i - 1; j++{
			if j == 1 || j <= 2*i - 1 || i == totallever{
				fmt.Print("*")
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	} 
}