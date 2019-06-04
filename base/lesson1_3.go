package base

import (
	"bufio"
	"os"
	"fmt"
)

/*
 * 实现类似于Unix中的uniq命令
 * 打印标准输入中多次出现的行，以重复次数开头
 */

 //第一个版本，主要处理标准输入中的数据

/* func main(){
 	counts := make(map[string]int)
 	input := bufio.NewScanner(os.Stdin)
 	//类似于while()，golang中没有while循环
 	for input.Scan(){
 		counts[input.Text()]++
	}

	for line,n := range  counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
 }*/

 //第二个版本，可以处理文件中的数据（可以是多个文件）

 //用一个map结构去统计f代表的文件中每行字符串出现的重复次数
 func countLines(f *os.File, counts map[string]int){
	 input := bufio.NewScanner(f)
	 for input.Scan(){
	 	counts[input.Text()]++
	 }
 }

 func main(){
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0{
		countLines(os.Stdin,counts)
	}else{
		for _, arg := range files{
			f, err := os.Open(arg)
			if err != nil{
				fmt.Fprintf(os.Stderr,"dup2:%v\n",err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line,n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

 }