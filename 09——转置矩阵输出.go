// package main

// import (
// 	"fmt"
// )

// func main(){
// 	var arr = [3][3]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	// fmt.Println(arr)
// 	// for循环遍历数组
// 	for i := 0; i < len(arr); i ++ {
// 		for j := 0; j < len(arr[i]); j ++ {
// 			fmt.Print("\t",arr[i][j])
// 		}
// 		fmt.Println( )
// 	}
// 	fmt.Println(" ")
// 	fun(arr)
// 	// 转置矩阵输出
// }

// func fun(arr [3][3]int){
// 	for g := 0; g < len(arr); g ++{
// 		for k := 0; k < len(arr[g]); k ++{
// 			fmt.Print("\t",arr[k][g])
// 		}
// 		fmt.Println( )
// 	}
// }


package main

import (
	"fmt"
)

func main(){
	var arr = [][] int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for i := 0; i < len(arr); i ++ {
		for j := 0; j < len(arr); j ++ {
			fmt.Print("\t", arr[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	for g := 0; g < len(arr); g ++ {
		for k := 0; k < len(arr); k  ++{
			fmt.Print("\t", arr[k][g])
		}
		fmt.Println()
	} 
}