package main

import(
	"fmt"
)

func main(){
	var g, s, b int
	for i := 100; i < 1000; i ++{
		b = i / 100
		s = (i % 100) / 10
		g = i % 10
		sum1 := b * b * b
		sum2 := s * s * s
		sum3 := g * g * g
		if sum1 + sum2 + sum3 == i {
			fmt.Println(i, "是水仙花数")
		}
	}
}