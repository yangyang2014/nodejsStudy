package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"runtime/debug"
)

//初始化时，邀请码都没有被占用，初始化为0，支持0-999999位
var flag [31250]uint32

/**
* 生成随机邀请码
* @return inviteCode 邀请码
**/
func GenInviteCode() string{
	// var i uint32 = 0  
	for {
		// 解决随机数不随机的问题 耗性能
		rand.Seed(time.Now().Unix())
		//获取0-999999的随机数
		inviteCodeNum := rand.Intn(1)
		// inviteCodeNum := i
		// i++
		// fmt.Println(inviteCodeNum)
		inviteCodeIndex := inviteCodeNum>>5
		inviteCodePos := uint32(inviteCodeNum&31)
		isExist := flag[inviteCodeIndex]&(1<<inviteCodePos) 
		if(isExist == 0) {
			flag[inviteCodeIndex] = flag[inviteCodeIndex]|1<<inviteCodePos
			//整型转字符串 不足6位时 高位以零补齐
			return fmt.Sprintf("id:%06d", inviteCodeNum)
		} 
	}
}

/**
* 释放指定的邀请码
* @param inviteCode 邀请码
* @return 错误信息，如果有错误信息就返回，否则返回空字符串
**/
func Drop(inviteCode string) string{

	inviteCodeNum , err := covertInviteCodeToInt32(inviteCode)
	if err != nil {
		//TODO 提示用户该邀请码不存在
		return "该邀请码[" + inviteCode + "]不存在"
	} else {
		inviteCodeIndex := inviteCodeNum>>5
		inviteCodePos := inviteCodeNum&31
		flag[inviteCodeIndex] = flag[inviteCodeIndex]&^(1<<inviteCodePos)
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
* 将邀请码转为整型uint32
* @param inviteCode 邀请码
* @return inviteCodeNum 邀请数值
* @return error 错误
**/
func covertInviteCodeToInt32(inviteCode string) (uint32, error){
   i, err := strconv.Atoi(inviteCode)
   if err != nil {
		fmt.Println(err)
	    fmt.Printf("%s", debug.Stack())
		return 0,err
   } 
   return uint32(i),nil
}

// var count = 0
// func testGoRoutine(flag int){
// 	for i:=0;i<1000;i++{
// 		count ++
// 		// fmt.Println("t",flag," ",GenInviteCode())
// 	}
// }

// func main() {
// 	// Init()
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(Drop ("q00001"))
// 	// fmt.Println(Drop ("000002"))
// 	// fmt.Println(Drop ("000003"))
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())
// 	// fmt.Println(GenInviteCode())

// 	//测试高并发的情况下生产邀请码
// 	go testGoRoutine(1)
// 	go testGoRoutine(2)
// 	go testGoRoutine(3)
// 	go testGoRoutine(4)
// 	go testGoRoutine(5)
// 	go testGoRoutine(6)
// 	go testGoRoutine(7)
// 	go testGoRoutine(8)
// 	for{
// 		time.Sleep(1*time.Second)
// 		fmt.Println(count)
// 	}
// }
