package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"runtime/debug"
	"sync/atomic"
)

//初始化时，邀请码都没有被占用，初始化为0，支持0-999999位
var flag [31250]uint32

/**
* 生成随机邀请码
* @return inviteCode 邀请码
**/
func GenInviteCode() string{
 
	for {
		// 解决随机数不随机的问题 耗性能
		rand.Seed(time.Now().UnixNano())
		//获取0-999999的随机数
		inviteCodeNum := rand.Intn(1000000)
		inviteCodeIndex := inviteCodeNum>>5
		inviteCodePos := uint32(inviteCodeNum&31)

		originValue := flag[inviteCodeIndex]
		originValueAddr := &flag[inviteCodeIndex]
		isExist := originValue&(1<<inviteCodePos) 
		//TODO to delete 用于测试并发问题
		// time.Sleep(time.Second)
		if(isExist == 0) {
			newValue := originValue|1<<inviteCodePos
			swagFlag := atomic.CompareAndSwapUint32(originValueAddr, originValue, newValue)
			if swagFlag {
				return formatIntToInviteCode(inviteCodeNum)
			} 
		} 
	}
}

/**
* 释放指定的邀请码
* @param inviteCode 邀请码
* @return 错误信息，如果有错误信息就返回，否则返回空字符串
**/
func Drop(inviteCode string) string{

	inviteCodeNum , err := covertInviteCodeToUint32(inviteCode)
	if err != nil {
		//TODO 提示用户该邀请码不存在
		return "该邀请码[" + inviteCode + "]不存在"
	} else {
		inviteCodeIndex := inviteCodeNum>>5
		inviteCodePos := inviteCodeNum&31
		atomic.StoreUint32(&flag[inviteCodeIndex],atomic.LoadUint32(&flag[inviteCodeIndex])&^(1<<inviteCodePos))
		// flag[inviteCodeIndex] = flag[inviteCodeIndex]&^(1<<inviteCodePos)
		return ""
	}

}

/**
*初始化邀请码的使用情况
**/
func Init() {
//修改flag
//TODO
	flag[0] = 3
}

/**
*整型转邀请码字符串 不足6位时 高位以零补齐
*@param value 整型数字，范围0-999999
*@return 6位邀请码
**/
func formatIntToInviteCode(value int)string{
	str := strconv.Itoa(value);
	length := len([]rune(str))
	switch length {
		case 1:
			return "00000"+str
		case 2:
			return "0000"+str
		case 3:
			return "000"+str
		case 4:
			return "00"+str
		case 5:
			return "0"+str
		case 6:
			return str
		default:
			return "-1"
	}
}
/**
* 将邀请码转为整型uint32
* @param inviteCode 邀请码
* @return inviteCodeNum 邀请数值
* @return error 错误
**/
func covertInviteCodeToUint32(inviteCode string) (uint32, error){
   i, err := strconv.Atoi(inviteCode)
   if err != nil {
		fmt.Println(err)
	    fmt.Printf("%s", debug.Stack())
		return 0,err
   } 
   return uint32(i),nil
}

// var count uint64= 0
func testGoRoutine(flag int){
	for i:=0;i<10;i++{
		// count ++	
		// atomic.AddUint64(&count, 1)
		// fmt.Println("t",flag," ",GenInviteCode())
		// begintime := time.Now().UnixNano() / 1e6
		begintime := time.Now().Unix()
		code := GenInviteCode()
		fmt.Println(code,"gen success!","begintime=",begintime,"endtime=",time.Now().Unix(),"flag = ",flag)
		// GenInviteCode()
	}
}

func testDropCode(flag int){
	 for{
		time.Sleep(time.Second)
		res := Drop("000000")
		 Drop("000001")
		Drop("000002")
		fmt.Println("drop res:",res)
		// GenInviteCode()
	}
}

func main() {
	// Init()
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
	for i:=0;i<10000;i++{
		go testGoRoutine(i)
	}
	//测试高并发情况下drop问题
	for i:=0;i<1;i++{
		go testDropCode(i)
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