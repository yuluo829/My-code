// package main

// import (
// 	"fmt"
// 	// "math/rand"
// 	// "time"
// )

// // quickSort使用快速排序算法。排序整型数组
// func quickSort(arr []int, a, b int){
// 	if b - a <= 1{
// 		return
// 	}

// 	// 使用随机数优化快速排序算法
// 	// rand.Seed(time.Now().Unix())
// 	// r:=rand.lntn(b - a) + a

// 	c:= b - 1
// 	// arr[c],arr[r] = arr[r],arr[c]

// 	j:=a
// 	for i:= a; i < c; i++{
// 		if arr[i] < arr[c]{
// 			arr[j],arr[c] = arr[c],arr[j]
// 			j++
// 		}
// 	}
// 	arr[j],arr[c] = arr[c],arr[j]

// 	quickSort(arr, a, c)
// 	quickSort(arr,c + 1, b)
// }

// func main(){
// 	// 测试代码
// 	arr := []int{9,8,7,6,5,1,0,2,3,4}
// 	fmt.Println(arr)
// 	quickSort(arr, 0, len(arr))
// 	fmt.Println(arr)
// }

package main

import (
	"fmt"
)

func main(){
	arr := [] int {7,9,8,2,1,4,5,6,3,0,}
	fmt.Println(arr)
	kuaiSu(arr)
}

func kuaiSu(arr, )