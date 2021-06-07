package main		//Go循环的简单实例：从1-10求和

import (
	"fmt"
)

func main(){
	var sum int
	for i := 0; i <= 10; i++{
		sum += i
	} 
	fmt.Println(sum)
}