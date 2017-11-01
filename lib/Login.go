package lib

import (
	"net/http"
	"Collector/util"
	"Collector/model"
	"encoding/json"
	"log"
	"upper.io/db.v3/mongo"
	"fmt"
)

//vue 调用 post方法 参数:username password 返回
func Login(w http.ResponseWriter, r *http.Request) {

	var lr model.LoginReturn
	var user model.User

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	authstr := r.FormValue("authstr")

	if (util.GetMD5Hash(username+password) == authstr){
		//进一步从数据库验证用户账号合法性
		lr.Code = 200

		var settings = mongo.ConnectionURL{
			Host:"106.14.2.153",
			Database:"Collector",
			User:"",
			Password:"",
		}

		sess,err := mongo.Open(settings)
		if err!=nil{
			fmt.Println("connect to mongo err"+err.Error())
			lr.Code = 401
		}
		defer sess.Close()

		err = sess.Collection("rbac_r").Find(model.User{Name:"admin"}).One(&user)
		if err!=nil{
			fmt.Println("rbac_r.err"+err.Error())
			lr.Code = 402
		}else{
			lr.Role = user.Role
			//进一步从Role中获取对应权限树


		}




	}else{
		//登陆加密错误
		lr.Code = 500
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


