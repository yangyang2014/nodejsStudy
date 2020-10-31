package main

import (
	"fmt"
	// "time"

)

func newTask(){
	i:=0
	for{
		i++
		fmt.Printf("new goroutine:i%d\n",i)
		// time.Sleep(1*time.Second)
	}
}
//主协程
// func main() {
// 	//启动子协程
// 	go newTask()
// 	i:=0
// 	for{
// 		i++
// 		fmt.Printf("main goroutine：i=%d\n",i)
// 		// time.Sleep(1*time.Second)
// 	}

// }

// 主协程和子协程交替执行，调度器会自动将其安排到合适的系统线程上执行