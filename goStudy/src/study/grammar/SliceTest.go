package main

import "fmt"

func main() { 
	//创建一个长度为3的切片
	s := make([]int, 3)
	fmt.Println(s)

	//可以像操作数组一样操作它
	s[0] = 1
	fmt.Println(s)
	fmt.Println(s[1])
	fmt.Println("len",len(s))
	//除了上面的基础操作外，还有比数组更丰富的操作
	// 1、支持追加
	s = append(s, 2)
	s = append(s, 3, 4)
	fmt.Println("append:", s)
	//TODO 测试失败 append来实现元素的删除
	//  将删除点index前后的元素连接起来
	// seq = append(seq[:index], seq[index+1:]...)
	s = append(s[:1], s[3:])
	// 2、支持切片的copy
	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("cpy",c)

	// 3、支持切的操作 slice[low:high] 获取部分数据
	l := s[1:4]
	fmt.Println(l)
	l2 := s[:5]
	fmt.Println(l2)
	t := s[2:]
	fmt.Println(t)

	//创建切片的另外一种方式
	// s2 := []string {"1","2","3"}

	//和数组一样支持二维的数据结构，但是一维度和二维上切片的长度是可变的
	twoD := make([][]int, 3)
	for i:=0;i<3;i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	subSlice := []int{1}
	twoD = append(twoD,subSlice)
	fmt.Println(twoD)

	

}