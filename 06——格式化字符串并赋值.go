package main

import (
	"fmt"
	"io"
	"os"
)

func main(){		//go中格式化字符串并赋值给新串，使用fmt.Sprintf方法	%s表示字符串
	var stockcode = "000987"
	var enddate = "2020-12-31"
	var url = "Code=%s&endDate=%s"
	var target_url = fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)
	//%d表示整型
	const name, age = "Kim",22
	s:=fmt.Sprintf("%s is %d years old.\n",name, age)
	io.WriteString(os.Stdout,s)		//忽略了一些错误
}