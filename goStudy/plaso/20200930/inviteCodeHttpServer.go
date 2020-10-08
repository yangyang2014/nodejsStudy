package main 
import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"encoding/json"
)

type JsonResult  struct{
    Code int `json:"code"`
    Obj  string `json:"obj"`
}


func main() {

	Init()
	http.HandleFunc("/genInviteCode",  genInviteCodeHandler)
	http.HandleFunc("/drop",  dropInviteCodeHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}

func genInviteCodeHandler(w http.ResponseWriter, r *http.Request) { 
	inviteCode := GenInviteCode()
	res := new(JsonResult)
	res.Code = 0
	res.Obj = inviteCode
	res_json,_ := json.Marshal(res)
	fmt.Fprint(w, string(res_json))
}

func dropInviteCodeHandler(w http.ResponseWriter, r *http.Request) { 
	//解析url参数 获取邀请码
	var code string
	r.ParseForm() 
	for k, v := range r.Form {
		if(strings.Compare("inviteCode",k) == 0) {
			code = v[0]
			break
		}
	}
	
	res := new(JsonResult)
	
	if(len([]rune(code)) == 6) {
		//邀请码存在且长度为6
		info := Drop(code)
		res.Code = 0
		res.Obj = info
	} else {
		res.Code = 400
		res.Obj = "无效的邀请码,请检查参数"
	}
	res_json,_ := json.Marshal(res)
	fmt.Fprint(w, string(res_json))
}