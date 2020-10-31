package main

import (

    "fmt"

)

func main () {
    var str string
    
    fmt.Println(len([]rune(str)))
    str = ""
    fmt.Println(len([]rune(str)))
    str = "1"
    fmt.Println(len([]rune(str)))

}