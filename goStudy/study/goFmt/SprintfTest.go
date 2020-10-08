package main
import (
"fmt"
"time"
"strconv"
)

func main(){
	var limit int = 200000000
	// fmt.Println("sprintf ")
	// fmt.Println(time.Now().Unix())
	// for i:=0;i<limit;i++{
	// 	callSprintf(i)
	// 	// fmt.Print(callSprintf(i)," ")
	// }
	// fmt.Println()
	// fmt.Println(time.Now().Unix())

	fmt.Println("callSwitch ")

	
	fmt.Println(time.Now().Unix())
	begintime1 := time.Now().UnixNano()
	for i:=0;i<limit;i++{
		callSwitch(100)
		// fmt.Print(callSwitch(i)," ")
	}
	
	// fmt.Println()
	endtime1 := time.Now().UnixNano()
	fmt.Println(time.Now().Unix())
	fmt.Println(begintime1-endtime1)

	fmt.Println("callForLoop ")
	fmt.Println(time.Now().Unix())
	begintime2 := time.Now().UnixNano()
	for i:=0;i<limit;i++{
		callForLoop(100)
		// fmt.Print(callForLoop(i)," ")
	}
	// fmt.Println()
	endtime2 := time.Now().UnixNano()
	fmt.Println(time.Now().Unix())
	fmt.Println(begintime2-endtime2)

}

func callSprintf(i int)string{
	res := fmt.Sprintf("%06d",i)
	return res;
}

func callSwitch(i int)string{
	str := strconv.Itoa(i);
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

func callForLoop(i int)string{ 
	str := strconv.Itoa(i);
	length := len([]rune(str))
	for {
		if(length < 6 ){
			length ++
			str = "0" + str
		} else {
			break
		}
	}
	return str
}