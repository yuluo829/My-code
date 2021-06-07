package main		//常量的使用

import(
	"unsafe"
)

const (		//常量表达式中，函数必须是内置函数，否则编译不过
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main(){
	println(a,b,c)
}