package main
import (
	"fmt"
	"time"
)

func say (s string) {
	for i := 0;i<5;i++{
		time.Sleep(10000*time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
//有疑问 会启动goroutine线程，如果方法中只有go xxx执行方法，那么就不继续执行，
//go say("hello")
	say("world")
	
}