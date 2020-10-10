package main

import ("fmt")
//使用range遍历数组、切片、map、字符串
func main () {
	//#遍历切片，数组类似
	nums := []int{1,2,3,4,5}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i,num := range nums {
		fmt.Println(i," ",num)

	}
	//#遍历map
	maps := map[string]int{"k1":1,"k2":2}
	//遍历k v元素
	for k,v := range maps {
		fmt.Println(k, " ", v)
	}
	//也支持只遍历k和v

	//#遍历字符串中单个字符 默认将字符串转为rune数组，第一个返回值为数组索引，第二个为unicode值
	str := "ABCDEFGabcdefg"
	for i, charvalue := range str {
		fmt.Println(i," ",charvalue)
	}

}