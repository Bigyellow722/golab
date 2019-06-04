package algo


//通过数组指针对数组进行反转
func ReverseArray(s *[5]int){
	i,j := 0,len(*s)-1
	for i<j{
		(*s)[i],(*s)[j]=(*s)[j],(*s)[i]
		i+=1
		j-=1
	}
}