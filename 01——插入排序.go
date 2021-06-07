package main

import (
	"fmt"
)

func main(){
	arr := [] int { 1,0,8,9,7,5,4,6,3,2}
	fmt.Println(arr)
	chaRu(arr)
}

func chaRu(arr [] int) {
	n := len(arr)
	for i := 1; i < n; i ++ {
		 for j := i; j >= 1; j-- {
			k := j - 1
			if arr[k] < arr[j]{
				break
			}
			// arr[k], arr[j] = arr[k], arr[j]
			t := arr[k]
			arr[k] = arr[j]
			arr[j] = t
		}
	}
	fmt.Println(arr)
}