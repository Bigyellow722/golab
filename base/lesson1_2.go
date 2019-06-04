package base

import (
	"os"
	"fmt"
)

//os.Arg变量是一个字符串的切片，第一个参数是命令本身的名字
//其他的参数是程序启动时传给它的参数

//实现Unix里echo命令的功能
func echo(){
	for i := 0; i < len(os.Args); i++{
		fmt.Println(i,os.Args[i])
	}

}