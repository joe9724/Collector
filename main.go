package main

/**
定义路由
**/

import (
	"net/http"
	"log"
	"fmt"
	"Collector/lib"
	
)


func main() {
	startHttpServer()
}

func startHttpServer() {
	//后台登录
	http.HandleFunc("/login", lib.Login)


	//http.Handle("/img/", http.StripPrefix("/img", http.FileServer(http.Dir("/usr/img"))))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("collector server start success...")
}



