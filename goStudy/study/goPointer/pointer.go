
package main

import "fmt"

func main() {
    var a int = 1
    var b *int = &a
    var c **int = &b
    var x int = *b
    fmt.Println("a = ",a) //1
    fmt.Println("&a = ",&a) //地址
    fmt.Println("*&a = ",*&a)//1
    fmt.Println("b = ",b)//a变量的地址
    fmt.Println("&b = ",&b)//b变量的地址
    fmt.Println("*&b = ",*&b)//a变量的地址
    fmt.Println("*b = ",*b)//获取地址下的值
    fmt.Println("c = ",c)//b变量的地址
    fmt.Println("*c = ",*c)//b变量的值 a的地址
    fmt.Println("&c = ",&c)//c的地址
    fmt.Println("*&c = ",*&c)//c的值 b的地址
    fmt.Println("**c = ",**c)//1
    fmt.Println("***&*&*&*&c = ",***&*&*&*&*&c)//1
    fmt.Println("x = ",x)//1
}