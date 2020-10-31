package main

import ("fmt"
"sync"
"runtime"
"time"
)

func TestGoroutine(){

	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		wg.Done()
	}()

	go func() {
		fmt.Println(65)
		fmt.Println(66)
		// 设置个睡眠，让该协程执行超时而被挂起，引起超时调度
		time.Sleep(time.Second)
		fmt.Println(67)
		wg.Done()
	}()

	wg.Wait()


}

func main(){
	TestGoroutine()
}