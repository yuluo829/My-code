package main		//快速排序，其复杂度与数组长度有关,,,,,有问题的排序

import (
	"fmt"
)

func main(){
	arr := [] int {1,8,4,3,6,5,4,9,2,0,7}
	fmt.Println(arr)
	kuaiSu(arr)
}

func kuaiSu(arr []int, left, right int) {
	if left >= right {
		return 
	}
	index, key := left, arr[left]

	// n := len(arr)
	for i := left + 1; i <= right; i++ {
		if arr[i] < key {
			index ++
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[left], arr[index] = arr[index], arr[left]

	kuaiSu(arr, left, index - 1)
	kuaiSu(arr, index + 1, right)
}

func kuaiSu1(arr [] int) {
	kuaiSu(arr, left := 0, len(arr) - 1)
}