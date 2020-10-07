package main 
import (
	"fmt"
	"log"
	"net/http"
	"strings"
	// "html"
)

func main() {

	Init()
	http.HandleFunc("/genInviteCode",  genInviteCodeHandler)
	http.HandleFunc("/drop",  dropInviteCodeHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}

func genInviteCodeHandler(w http.ResponseWriter, r *http.Request) { 
	inviteCode := GenInviteCode()
	fmt.Fprint(w, inviteCode)
}

func dropInviteCodeHandler(w http.ResponseWriter, r *http.Request) { 
	//解析url参数 获取邀请码
	var code string
	r.ParseForm() 
	for k, v := range r.Form {
		if(strings.Compare("code",k) == 0) {
			fmt.Print("key:", k, ";")
		    fmt.Println("val:", v[0])
			code = v[0]
			break
		}
	}
	
	if(strings.Compare(code,"") != 0 && strings.Count(code, "") - 1 == 6) {
		//邀请码存在且长度为6
		info := Drop(code)
		fmt.Fprint(w, info)
	} else {
		fmt.Fprint(w, "无效的邀请码,请检查参数")
	}
}