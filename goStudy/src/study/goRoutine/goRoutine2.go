package main
import (
	"sync"
	"log"
)
func structRaceWrong() {

	x := struct {
	A int
	B int
	}{1, 1}
	
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	if x.A != x.B {
	log.Println("A != B", x.A, x.B)
	}
	if i < 3 {
	log.Printf("%p\n", &x)
	}
	}
	wg.Done()
	}()
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	x = struct {
	A int
	B int
	}{x.A + 1, x.B + 1} // <---------- 导致 A、B 可能不一致
	}
	wg.Done()
	}()
	
	wg.Wait()
	}
	
	// 结论：Go 是值拷贝，结构体会出现复制一半就被其它读取了
	func structRaceWrong2() {
	
	x := struct {
	A int
	B int
	}{1, 1}
	
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	cp := x
	if cp.A != cp.B {
	log.Println("A != B", cp.A, cp.B)
	}
	if i < 3 {
	log.Printf("%p\n", &x)
	}
	}
	wg.Done()
	}()
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	cp := x // <---------- 拷贝也不行，可能只拷贝一半
	x = struct {
	A int
	B int
	}{cp.A + 1, cp.B + 1}
	}
	wg.Done()
	}()
	
	wg.Wait()
	}
	
	// 结论：通过指针读取的数据是完整的，但不一定是最新的。
	func structRaceOk() {
	
	x := &struct {
	A int
	B int
	}{0, 0}
	
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	cp := x
	if cp.A != cp.B { // <---------- 通过指针读取的数据是完整的，但不一定是最新的。
	log.Fatal("A != B", cp.A, cp.B) // 不会出现
	}
	if i < 10 { // 抽取前 10 条数据检查
	log.Printf("%p -> %p\n", &x, x)
	log.Printf("-A%d - i%d = %d", cp.A, i, cp.A - i) // 不一定是最新的
	}
	if i - cp.A > 1 {
	log.Printf("A:%d - i:%d = %d \n", cp.A, i, cp.A - i)
	}
	}
	wg.Done()
	}()
	
	wg.Add(1)
	go func() {
	for i := 0; i < 1000000; i++ {
	cp := x
	x = &struct {
	A int
	B int
	}{cp.A + 1, cp.B + 1}
	}
	wg.Done()
	}()
	
	wg.Wait()
	}