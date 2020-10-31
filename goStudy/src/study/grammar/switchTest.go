package main
import(
	"fmt"
	// "time"
)

func main() {
	//c和c++中如果遇到合适的case，就会执行之后的语句，直到遇到break，如果没有遇到那么就会执行下一个case（如果存在）
	//但是在go 语言中，进入一个分支后，不管是否有break都直接结束了 这和ifelse相同，
	//除非使用 fall through 才会继续执行紧跟的下一个case 
	i:=8
	// i:=1
	//1.左花括号{必须与switch处于同一行； 
// 2.条件表达式不限制为常量或者整数； 
// 3.单个case中，可以出现多个结果选项；

//switch 优点 ，switch是在编译阶段将子函数的地址和判断条件绑定了，只要直接将a的直接映射到子函数地址去执行就可以了
	//https://blog.csdn.net/Zbylant/article/details/104787765/

	//1、如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型，建议使用switch语句
// 2、对区间判断和bool类型的判断，使用if，if的使用范围更广。

//https://zhuanlan.zhihu.com/p/65305380
	switch i {
	case 0 :
		fmt.Printf("0")
		fallthrough
	case 1 :
		fmt.Println("1")
	case 2 :
		fmt.Println("2")
		fallthrough
	case 3 :
		fmt.Println("3")
	case 4,5,6:
		fmt.Println("4,5,6")
		fallthrough
	default :
	    fmt.Println("default")
	}
	
}