package main		//稀疏矩阵额压缩存储

import (
	"fmt"
)

func main(){
	type Res struct {
		i int
		j int
		v int
	}
	var res []Res
	a := [][]int{		//原矩阵
		{1,0,0,0,0},
		{0,7,0,0,0},
		{0,0,0,0,4},
		{0,3,1,0,1},
		{0,0,8,0,0},
	}
	for i := 0; i < len(a[0]); i++{
		for j := 0; j < len(a); j++{
			if a[i][j] != 0 {
				res = append(res, Res{i, j, a[i][j]})
			}
		}
	}
	for _,r := range res{
		fmt.Println(r.i, r.j, r.v)
	}
}