package main 
import (
	"fmt"
)

//测试指针在函数的用法，基本数据类型参数也可以通过函数修改值
//TODO 解释待完善 这与没有指针的java不同，
func zeroval (ival int) {
	ival = 0
}

func zeroptr (ptr *int) {
	*ptr = 0
}

func main () {

	i := 1
	zeroval(i)
	fmt.Println(i)
	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)

}