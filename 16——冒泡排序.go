package main		//冒泡排序

import (
	"fmt"
)

func maoPaosort(arr [] int){
	for i := 0; i < len(arr); i ++ {
		for j := 0; j < len(arr) - i - 1; j ++{
			k := j + 1
			if arr[k] < arr[j] {
				// arr[k], arr[j] = arr[j], arr[k]
				t := arr[j]
				arr[j] = arr[k]
				arr[k] = t
			}
		}
	}
	fmt.Println(arr)
} 

func main()  {
	var arr = [] int{7,5,6,0,1,2,3,8,9,4}
	fmt.Println(arr)
	maoPaosort(arr)
}
// func main(){
// 	arr := []int{7,5,3,9,1,4,6,2,8}
// 	fmt.Println(arr)
// 	for i := 0; i < len(arr); i ++{
// 		for j := 0; j < len(arr) - i - 1; j ++{
// 			k := j + 1
// 			if arr[k] < arr[j]{
// 				t := arr[j]
// 				arr[j] = arr[k]
// 				arr[k] = t
// 			}
// 		}
// 	}
// 	fmt.Println(arr)
// }