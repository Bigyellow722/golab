package basestruct

import (
	"fmt"
	"testing"
)

func TestCache_Set(t *testing.T) {
	t.Parallel()
	go func() {

	}()
}
func BenchmarkCache_Set(b *testing.B) {

}

func TestLoop(t *testing.T) {
	sum := 0
	nums := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var numfunc []func()

	for i := 0; i < 10; i++ {
		sum = sum + nums[i]
		numfunc = append(numfunc, func() {
			fmt.Println(nums[i])
		})
		fmt.Println("匿名函数外", nums[i])

	}
	fmt.Println(sum)
	for _, mynum := range numfunc {
		mynum()
	}
}
