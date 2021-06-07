// package main		//选择排序

// import (
// 	"fmt"
// )

// func main(){
// 	arr := []int{0,9,8,5,3,2,1,6,7,4}
// 	fmt.Println("初始数组为：", arr)
// 	SSort(arr)
// }

// func SSort(arr [] int){
// 	var m, site int
// 	n := len(arr)
// 	for i := 0; i < n - 1; i ++{
// 		m = arr[i]
// 		site = i
// 		for j := i + 1; j < n; j++{
// 			if m > arr[j]{
// 				m = arr[j]
// 				site = j
// 			}
// 		}
// 		arr[site] = arr[i]
// 		arr[i] = m
// 	}
// 	fmt.Println("排序之后的数组为：", arr)
// }

// package main

// import (
// 	"fmt"
// )

// func main(){
// 	arr := [] int {7,8,3,5,6,0,1,9,4,2}
// 	fmt.Println("初始数组为：", arr)
// 	xuanZe(arr)
// }

// func xuanZe(arr [] int){
// 	n := len(arr)
// 	var m, site int
// 	for i := 0; i < n; i ++{
// 		m = arr[i]
// 		site = i
// 		for j := i; j < n; j ++{
// 			if m > arr[j]{
// 				m = arr[j]
// 				site = j
// 			}
// 		}
// 		arr[site] = arr[i]
// 		arr[i] = m
// 	}
// 	fmt.Println("排序后的数组为：", arr)
// }

package main

import (
	"fmt"
)

func main(){
	var arr = [] int {1,5,3,4,2,8,0,9,6,7}
	fmt.Println("初始数组为：", arr)
	xuanze(arr)
}
//降序排列之后的数组
func xuanze(arr [] int){
	n := len(arr)
	var m, site int
	for i := 0; i < n; i ++{
		m = arr[i]
		site = i
		for j := i; j < n; j ++{
			if m < arr[j]{
				m = arr[j]
				site = j
			}
		}
		arr[site] = arr[i]
		arr[i] = m
	}
	fmt.Println("排序后的数组为：", arr)
}