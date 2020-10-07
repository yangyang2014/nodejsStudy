# goStudy

##go基本数据类型有哪些
1、有符号整型 int8 int16 int32 int64
2、无符号整型 uint8 uint16 uint32 uint64
3、与平台相关的数据类型 int 和 uint 
4、特殊的无符号整型 uintptr，在32位平台是4个字节，在64位平台为8个字节，用来存储指针
不同整型进行算术运算和比较运算会出现编译错误。需要进行强制类型转换
5、浮点型 float32 float64 complex64 complex128
6、常量 编译时生成
  iota常量生成器
7、常量
8、布尔类型

##特殊的类型
1、指针类型
unsafe.Pointer类型 指针类型 用于实现定位和读写内存的基础 可以与uintptr互转
如何使用unsafe.Pointer和uintptr功能
2、数组类型
3、结构化类型
4、channel类型
5、函数类型
6、切片类型
7、接口类型
8、map类型

##常用语法有哪些？
go中没有while
如果依赖包没有使用到也会报出语法错误


数组的声明
var variable_name [SIZE] variable_type
var score [10] float32
数组的初始化
score = [3]float32{100.0,2.0,3.4}

##go中的nil是什么？和java中的null有什么区别？

##如何实现并发？

## go语言中的异常处理


## go 中方法与函数的区别
方法
Func main( a, b int) (int) {
}
函数
func (p myint) mysquare() int {  
  
}  

## struct 初始化

//第一种，在Go语言中，可以直接以 var 的方式声明结构体即可完成实例化
var t T
t.a = 1
t.b = 2

//第二种，使用 new() 实例化
t := new(T)

//第三种，使用字面量初始化
t := T{a, b}
t := &T{} //等效于 new(T)


示例
type Person struct {
    name string
    age int
    address string
}

ms := &Person{"name", 20, "bj"}
ms2 := &Person{name:"zhangsan"}
&Person{a, b, c} 是一种简写，底层仍会调用 new()，这里值的顺序必须按照字段顺序来写，同样它也可以使用在值前面加上字段名和冒号的写法（见上文的方式B，C）

## go中如何实现接口

## go中的线程和协程
多个协程可有一个或者多个线程管理，协程的调度都在其所在的线程中
调度策略可有应用层代码定义
执行效率高、占用内存不能小
https://juejin.im/post/6844903662553137165