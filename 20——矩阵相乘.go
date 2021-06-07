package main		//矩阵相乘

import (
	"fmt"
)

type SQ struct{
	m, n int
	data [][] int
}

func min(a SQ, b SQ) [][]int{
	res := [][]int{}
	for i := 0; i < a.n; i++{
		t := []int{}
		for j := 0; j < b.m; j++{
			r := 0;
			for k := 0; k < a.m; k++{
				r += a.data[i][k] * b.data[k][i]
			}
			t = append(t, r)
		}
		res = append(res, t)
	}
	return res
}

func main(){
	
	a := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	b := [][]int{
		{1, 2, 3},
		{3, 4, 1},
	}
	A := SQ{
		2, 3,
		a,
	}
	B := SQ{
		3, 2,
		b,
	}

	res := min(A, B)
	fmt.Println(res)
}