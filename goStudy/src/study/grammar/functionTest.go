package main
import (
	"fmt"
)
//可变参数的函数
func sum(nums ...int) {
	fmt.Print(nums," ")
	total := 0
	for _,num := range nums {
		total += num
	}
	fmt.Println(total)
}

//测试返回函数
func returnFunc() func() int {
	var i int = 0;
	return func () int {
		i++
		return i
	}
}


//测试递归功能
func recursion(n int)int {
	if n==1 || n==2 {
		return 1
	}
	return recursion(n-1) + recursion(n-2)
}

func main()  {
	sum(1, 2)
	sum(1, 2, 3)
	sum(1,2)
	nums := []int{1,2,3,4}
	//如果可变的参数在切片中，可以使用func(slice...)语法，如下
	sum(nums...)

	nextFunc := returnFunc()
	fmt.Println(nextFunc())
	fmt.Println(nextFunc())
	fmt.Println(nextFunc())
	nextFunc2 := returnFunc()
	fmt.Println(nextFunc2())
	
	fmt.Println(recursion(5))
}