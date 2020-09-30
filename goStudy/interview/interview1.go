package main

import "fmt"

type Person struct {
	name string
	age int
}

func change(a Person,b Person) {
	a.name,b.name = b.name,a.name
	a.age,b.age = b.age,a.age
	fmt.Println(a.name,a.age,b.name,b.age)
}

func main (){
	var a Person
	var b Person
	a.name = "xqy"
	a.age = 28

	b.name = "wyy"
	b.age = 27
	change(a,b)
	fmt.Println(a.name,a.age,b.name,b.age)
}