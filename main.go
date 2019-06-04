package main

import (
	"fmt"
	"golab/algo"
)

func main() {
	var s1 algo.MySliceInt64 = []int64{1,3,5,6}
	var s2 algo.MySliceInt64 = []int64{2,3,6,8,9,10}
	s := s1.MergeSortedSlice(s2)

	var s3 = [5]int{1,2,3,4,5}
	algo.ReverseArray(&s3)

	fmt.Println(s)
	fmt.Println(s[1:])
	fmt.Println(s3)
}
