package lib

import (
	"net/http"
	"Collector/util"
	"Collector/model"
	"encoding/json"
	"log"
)

//vue 调用 post方法 参数:username password 返回
func Login(w http.ResponseWriter, r *http.Request) {

	var lr model.LoginReturn

	r.ParseForm()
    username := r.FormValue("username")
    password := r.FormValue("password")
    authstr := r.FormValue("authstr")

    if (util.GetMD5Hash(username+password) == authstr){
    	//字符串合法
    	lr.Code = 500

	}else{
		lr.Code = 200
	}


	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "NanjingYouzi")
	data, err := json.Marshal(lr)
	if err != nil {
		log.Fatal("err marshal: ", err)
	}
	defer func() {
		w.Write(data)
	}()
}


