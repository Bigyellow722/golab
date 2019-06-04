package base

import (
	"testing"
	"strings"
)

/*
 * 两种字符串拼接方法的性能对比
 * 1、使用strings包中的Join方法
 * 2、直接使用+操作符
 */

func BenchmarkString2Join(b *testing.B){
	//fmt.Println(b.N)
	for i := 0; i < b.N; i++{
		input := []string{"Welcome","To","China"}
		result := strings.Join(input," ")
		if result != "Welcome To China"{
			b.Error("Unexcepted result:" + result)
		}
	}
}

func BenchmarkString2Plus(b *testing.B){
	//fmt.Println(b.N)
	for i := 0; i < b.N; i++{
		input := []string{"Welcome","To","China"}
		var s, sep string
		for j := 0; j < len(input); j++{
			s += sep +input[j]
			sep = " "
		}
		if s != "Welcome To China"{
			b.Error("Unexcepted result:" + s)
		}
	}
}