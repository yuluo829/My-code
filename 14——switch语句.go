package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){		//从标准输入读出数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil{
		fmt.Printf("An error occurred:%s\n", err)	//异常错误后推出
	os.Exit(1)
	}else{
		//用切片操作删最后的\n
		name := input[:len(input) - 1]
		fmt.Printf("Hello, %swhat can I did for you?\n", name)
	}
	for {
		input, err = inputReader.ReadString('\n')
		if err != nil{
			fmt.Printf("An error occurred:%s\n", err)
			continue
		}
		input = input[:len(input) - 1]
		//全部转换为小写
		input = strings.ToLower(input)
		switch input{
		case "":
			continue
		case "nothing", "bye":
			fmt. Println("Bye!")
			//正常退出
			os.Exit(0)
		default:
			fmt.Println("Sorry,I did not catch you.")
		}
	}
}