package algo

type MySliceInt64 []int64


//将两个有序的序列合并成一个有序的序列
func (s1 MySliceInt64) MergeSortedSlice(s2 []int64) []int64{
	lens1 := len(s1)
	lens2 := len(s2)

	lenres := lens1 + lens2

	res := make([]int64, lenres)

	var i,j,k int

	for i,j,k = 0,0,0; i < lens1 && j<lens2; {
		var cur int64
		if s1[i] < s2[j]{
			cur = s1[i]
			i++
		}else{
			cur = s2[j]
			j++
		}
		res[k] = cur
		k++
	}

	for i < lens1 {
		res[k] = s1[i]
		k++
		i++
	}

	for j < lens2 {
		res[k] = s2[j]
		k++
		j++
	}
	return res
}

//通过设置过滤函数，对序列s进行过滤，返回一个新的序列
//主要就是对序列中的每个元素做检查，在判断是否保留
func (s MySliceInt64) FilterSlice(filter func(x int64) bool) []int64{
	res := s[:0]

	for _, x := range s{
		if !filter(x){
			res = append(res,x)
		}
	}
	return res
}


//原地反转Slice
func (s MySliceInt64) ReverseSlice(){
	for i, j:=0,len(s)-1;i<j;i,j=i+1,j-1{
		s[i],s[j] = s[j],s[i]
	}
}

//删去序列中第i+1个元素
func (s MySliceInt64) RemoveAt(i int)[]int64{
	copy(s[i:],s[i+1:])
	return s[:len(s)-1]
}


