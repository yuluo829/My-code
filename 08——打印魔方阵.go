package main		//打印魔方阵

import (
	"fmt"
)

func main(){
	N := 3
	var arr [][] int
	for i := 0; i < N; i++{
		temp := make([] int, N)
		arr = append(arr, temp)
	}

	i := 0
	j := N / 2
	arr[i][j] = 1
	for d := 2; d <=(N * N); d++{		//data
		t1 := i
		t2 := j
		i = i - 1
		j = j - 1
		if i < 0 {
			i = N - 1
		}
		if j < 0{
			j = N - 1
		}
		if arr[i][j] == 0{
			arr[i][j] = d
			fmt.Println(i, j, arr[i][j])
		}else {
			i = (t1 + 1) % N
			j = t2
			arr[i][j] = d
			fmt.Println(arr[i][j])
		}
	}
	for i := 0; i < N; i++{
		for j := 0; j < N; j++{
			fmt.Println(arr[i][j])
		}
		fmt.Println()
	}
}