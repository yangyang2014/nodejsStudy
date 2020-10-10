package main

import (
	"fmt"
)


func main() {
	//创建一个空map make(map[key-type]val-type)
	m := make(map[string]int)
	// map中set方法
	m["k1"] = 0
	m["k2"] = 2
	fmt.Println("map",m)
	//map中get方法
	v1 := m["k1"]
	fmt.Println("v1 = ",v1)

	//使用len来实现获取数组长度
	fmt.Println("len = ",len(m))
	delete(m, "k2")
	fmt.Println(m)

	//可选的第二返回参数,表示是否存在
	value, prs := m["k2"]
	fmt.Println("prs:",value, prs)
	value1, prs1 := m["k1"]
	fmt.Println("prs:",value1, prs1)
	//新的写法 "_" 表示忽略某个值。单函数有多个返回值，用来获取某个特定的值
	_, prs2 := m["k2"]
	fmt.Println("prs:", prs2)
	

	//声明一个初始化的map
	n := map[string]int{"foo":1,"bar":2}
	fmt.Println("n:", n)
}