package main
import (
	"fmt"
	"math/rand"
	"strconv"
	// "time"
	"runtime/debug"
	"log"
    "net/http"
)

//初始化时，邀请码都没有被占用，初始化为0，支持0-999999位
var flag [31251]uint32

/**
* 生成随机邀请码
* @param inviteCode 邀请码
**/
func GenInviteCode() string{
	// var i uint32 = 0  
	for {
		// 解决随机数不随机的问题 很耗性能
		rand.Seed(time.Now().Unix())
		//获取0-999999的随机数
		inviteCodeNum := rand.Intn(1000000)
		// inviteCodeNum := i
		// i++
		// fmt.Println(inviteCodeNum)
		inviteCodeIndex := inviteCodeNum/32
		inviteCodePos := inviteCodeNum%32
		isExist := flag[inviteCodeIndex]&(1<<inviteCodePos) 
		if(isExist != 0) {
			//不存在
			continue
		} else {
			flag[inviteCodeIndex] = flag[inviteCodeIndex]|1<<inviteCodePos
			return formatInt32ToInviteCode(uint32(inviteCodeNum))
		}
	}
}

/**
* 释放指定的邀请码
* @param inviteCode 邀请码
**/
func Drop(inviteCode string) {

	inviteCodeNum , err := covertInviteCodeToInt32(inviteCode)
	if err != nil {
		return ;
	} else {
		inviteCodeIndex := inviteCodeNum/32
		inviteCodePos := inviteCodeNum%32
		flag[inviteCodeIndex] = flag[inviteCodeIndex]&^(1<<inviteCodePos)
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
* 将整型uint32格式化为邀请码
* @param inviteCodeNum 邀请码数值 范围在0-999999
**/
func formatInt32ToInviteCode(inviteCodeNum uint32) string {
	inviteCodeNumStr := strconv.Itoa(int(inviteCodeNum))
	if(len(inviteCodeNumStr)<6){
	    pos := len(inviteCodeNumStr)
		for {
			if pos >= 6 {
				break
			}
			pos ++
			inviteCodeNumStr = "0" + inviteCodeNumStr
		}
	} 
	return inviteCodeNumStr
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

func main() {
	
	// Init()
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	// Drop ("q00001")
	// Drop ("000002")
	// Drop ("000003")
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	// fmt.Println(GenInviteCode())
	http.HandleFunc("/", dealRequest) 
	// index 为向 url发送请求时，调用的函数
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func dealRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, GenInviteCode())
}
