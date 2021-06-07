package main 	//用指针方法传递数组


import (
	"fmt"
)

func main(){
	 arr := [3][3] int{
		{1,5,9},
		{3,5,7},
		{4,5,6},
	}
	// fmt.Println(arr)
	exchanges(&arr) 
	fmt.Println(arr)
}

func exchanges(arr *[3][3] int){
	fmt.Println(arr)
}