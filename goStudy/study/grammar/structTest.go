package main

import "fmt"

type person struct {
	name string
	age int
}

func newPerson(name1 string) *person {
	p := person{name: name1}
    p.age = 42
    return &p
}

func main() { 
	//按照定义字段顺序来创建person结构体
	fmt.Println(person{"Bob",20})
	//根据字段名来创建person
	fmt.Println(person{name:"alice",age:30})
	fmt.Println(&person{name:"alice",age:30})
	fmt.Println(newPerson("jon"))

	s:=person{name:"sean",age:50}
	fmt.Println(s.name)
	//todo 考虑传递指针和值的区别
	// sp := s 
	sp := &s 
	fmt.Println(sp.age)
	sp.age =51
	fmt.Println(s.age)
}