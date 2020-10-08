package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"runtime/debug"
	"sync/atomic"
)
// var count uint64= 0
func testGoRoutine(flag int){
	// for i:=0;i<10;i++{
		// count ++	
		// atomic.AddUint64(&count, 1)
		// fmt.Println("t",flag," ",GenInviteCode())
		// begintime := time.Now().UnixNano() / 1e6
		// begintime := time.Now().Unix()
		code := GenInviteCode()
		fmt.Println(code,"gen success!")
		// fmt.Println(code,"gen success!","begintime=",begintime,"endtime=",time.Now().Unix(),"flag = ",flag)
		// GenInviteCode()
	// }
}

		//drop 过程中并发drop相同的数字 会出现swag失败问题
			//drop 在同一个数组元素范围内比如0-31,32-63...也会出现swag失败的问题
			//drop 不在同一个数组元素范围内的数字，不会有影响
func testDropCode(flag1 int){
	//  for{
		// time.Sleep(time.Second)
		code := fmt.Sprintf("%06d",flag1)
		Drop(code)
		fmt.Printf("drop code =" + code + " %b %b\n",flag[1],flag[0])
		// Drop("00000")
		// res := Drop("000000")
		// fmt.Println("drop res:",res)
		// res2 := Drop("000002")
		// fmt.Println("drop res:",res2)
		// GenInviteCode()
	// }
}

func main() {
	Init()
	// go fmt.Println(GenInviteCode())
	// fmt.Println("enter")
	// go Drop ("000002")
	// go fmt.Println(GenInviteCode())
	// go fmt.Println(GenInviteCode())
	// fmt.Println("enter")
	// go Drop ("000001")
	// fmt.Println("enter")
	// // go fmt.Println(Drop ("000003"))
	// go fmt.Println(GenInviteCode())
	// go fmt.Println(GenInviteCode())
	// go fmt.Println(GenInviteCode())
	// go fmt.Println(GenInviteCode())
	// fmt.Println("enter")
	// go fmt.Println(GenInviteCode())

	// 测试高并发的情况下生产重复的邀请码
	for i:=0;i<100;i++{
		go testGoRoutine(i)
	}
	//测试高并发情况下drop问题
	for i:=0;i<20;i++{
		if i != 31{
			go testDropCode(i%10)
		}
		
	}
	// go testGoRoutine(2)
	// go testGoRoutine(3)
	// go testGoRoutine(4)
	// go testGoRoutine(5)
	// go testGoRoutine(6)
	// go testGoRoutine(7)
	// go testGoRoutine(8)
	// go testGoRoutine(1)
	// go testGoRoutine(2)
	// go testGoRoutine(3)
	// go testGoRoutine(4)
	// go testGoRoutine(5)
	// go testGoRoutine(6)
	// go testGoRoutine(7)
	// go testGoRoutine(8)
	// go testGoRoutine(1)
	// go testGoRoutine(2)
	// go testGoRoutine(3)
	
	for{
		time.Sleep(1*time.Second)
		// fmt.Println(count)
	}

	// 整型数组某个元素的赋值给某个变量，为 测试出变量申明时 分配了内存空间。并且为值传递。
	// var testArray = [3]uint32 {1,2,3}
	// e1 := testArray[0]
	// fmt.Println("testArray[0] ",testArray[0]," ",&testArray[0])
	// fmt.Println("e1 ", e1," ",&e1)
	// testArray[0] = 10
	// e1 = 100
	// fmt.Println("testArray[0] ",testArray[0]," ",&testArray[0])
	// fmt.Println("e1 ", e1," ",&e1)
	// var m int
	// fmt.Println(m," ",&m)
}