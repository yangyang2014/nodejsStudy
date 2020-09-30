package main
import (
	"fmt"
	"math/rand"
)

//初始化时，邀请码都没有被占用，初始化为0，支持0-999999位
var flag [31251]uint32

/**
* 生成随机邀请码
* @param inviteCode 邀请码
**/
func GenInviteCode() string{
	while(true){
		//获取0-999999的随机数
		//检查是否为uint32
		inviteCodeNum := rand.Intn(1000000)
		inviteCodeIndex := inviteCodeNum/32
		inviteCodePos := inviteCode%32
		isExist := flag[inviteCodeIndex]&1<<() 
		if(isExist == 0) {
			//不存在
			continue
		} else {
			flag[inviteCodeIndex] = flag[inviteCodeIndex]|1<<inviteCodePos
			return formatInt32ToInviteCode(inviteCodeNum);
		}
	}
	return ""
}

/**
* 释放指定的邀请码
* @param inviteCode 邀请码
**/
func Drop(inviteCode string) {

	inviteCodeNum := covertInviteCodeToInt32(inviteCode)
	inviteCodeIndex = inviteCodeNum/32
	inviteCodePos = inviteCode%32
	flag[inviteCodeIndex] = flag[inviteCodeIndex]&~(1<<inviteCodePos)
	
}

/**
*初始化邀请码的使用情况
**/
func Init() {
//修改flag
//TODO
	flag[1] = 65535
}

/**
* 将整型uint32格式化为邀请码
* @param inviteCodeNum 邀请码数值 范围在0-999999
**/
func formatInt32ToInviteCode(inviteCodeNum uint32) string {
	inviteCodeNumStr = strconv.Itoa(inviteCodeNum)
	if(len(inviteCodeNumStr)<6){
	    int pos = len(inviteCodeNumStr)
		while(pos < 6) {
			pos ++
			inviteCodeNumStr = "0" + inviteCodeNumStr
		}
	} 
	return inviteCodeNumStr
}

/**
* 将邀请码转为整型uint32
* @param inviteCodeNum 邀请码
**/
func covertInviteCodeToInt32(inviteCode string) uint32{
   return strconv.Atoi(a)
}

func main {
	
	Init()
	fmt.Println(GenInviteCode())

}
