package main	//格式化字符串

import (
	"fmt"
)
/*%d表示整型数字，%s表示字符串*/
func main()  {		//main函数报错的原因是一个文件夹里边只能有一个Go文件，参考go语言中文网！
	var stockcode = 123
	var enddate = "2020-12-31"
	var url = "code = %d&enddate = %s"
	var target_url = fmt.Sprintf(url,stockcode,enddate)
	fmt.Println(target_url)
}