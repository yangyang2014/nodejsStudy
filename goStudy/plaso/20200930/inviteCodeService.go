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
		return "该邀请码[" + inviteCode + "]不存在"
	} else {
		inviteCodeIndex := inviteCodeNum>>5
		inviteCodePos := inviteCodeNum&31
		for {
			originValue := flag[inviteCodeIndex]
			originValueAddr := &flag[inviteCodeIndex]
			isExist := originValue&(1<<inviteCodePos) 
			if isExist == 0 {
				//如果资源未被占用
				break
			}
			newValue := originValue&^(1<<inviteCodePos)
			swagFlag := atomic.CompareAndSwapUint32(originValueAddr, originValue, newValue)
			//修改成功则ok，失败则继续将对应位改为零直到成功
			if swagFlag {
				break
			}
		}
		return ""
	}

}

/**
*初始化邀请码的使用情况
**/
func Init() {
	//TODO
}

/**
*整型转邀请码字符串 不足6位时 高位以零补齐
*@param value 整型数字，范围0-999999
*@return 6位邀请码 -1表示异常
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
* todo 测试下 入参为带负号的字符串
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

