package main

import "fmt"

//https://www.cnblogs.com/aaronthon/p/10673841.html
//切片似乎就是获取数组的子数组。a[begin:end]就是获取数组索引为begin到end-1的元素形成一个新的数组.如果begin和end相同那么就无元素。
//切片为一个基于数组的数据结构，支持扩容、删除操作
func arraySum(x [3]int) int {
	sum := 0
	for i, v := range x {
		println(i, v)
		sum = sum + v
	}
	return sum
}

func main() {
	//数组中初始化大小也是类型的一部分
	//num := [3]int {1,2,3}
	//println(arraySum(num))
	//slice3ByMake()
	sliceAppend()
}

func slice() {
	//声明切片
	var a []string
	var b = []int{}
	var c = []bool{false, true}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func slice2() {
	//基于切片再次切片，索引不能超过原数组长度
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("type of b:%T\n", b)
	c := b[1:]
	fmt.Println(c)
	e := c[1:]
	fmt.Println(e)
	f := a[:]
	fmt.Println(f)
	//验证原数组某个元素值改变，是否影响切片后的数组。确实有影响，说明切片还是基于原来数组做的
	a[1] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func slice3ByMake() {
	a := make([]int, 2, 10)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

}

func sliceAppend() {

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v len:%d cap:%d ptr：%p \n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

}

//make切片的常用方法
func makeTest(m int) int64 {
	a := make([]int64, 0)
	if m > len(a)-1 {
		for i := len(a) - 1; i < m; i++ {
			a = append(a, 1)
		}
	}
	return a[m]
}
