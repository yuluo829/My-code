package main		//结构体面向对象编程设计

import "fmt"

type Person struct {	//创建结构体
	name string		
	age int
	sex byte
	addr string
}

type Student struct{
	Person		//只有变量名没有数据类型的为匿名变量，匿名字段就默认Student包含了Person的所有字段
	id int
}


func main(){
	var s1 Student = Student{Person{"mike", 18, 'm', "北京"}, 123}	//普通赋值型
	fmt.Println("s1 = ", s1)
	s2 := Student{Person{"mike", 18, 'm', "北京"}, 123}	//自动类型推导
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2 = %+v\n", s2)	//更加详细的输出s2的值
	s3 := Student{id:666}
	fmt.Println("s3 = ", s3)		//分别给指定成员初始化，未初始化成员默认赋值为0
	//成员操作
	s4 := Student{Person{"mike", 19, 'm', "深圳"}, 888}
	fmt.Printf("s4 = %+v\n", s4)
	//打印成员
	fmt.Println(s4.name, s4.sex, s4.addr, s4.id, s4.age)
	//结构体成员分别赋值
	s4.name = "yoyo"
	s4.sex = 'W'
	s4.addr = "上海"
	s4.id = 999
	s4.age = 20
	fmt.Println("s4 = ", s4)
}