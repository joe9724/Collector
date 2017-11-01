package model

type LoginReturn struct {
	Code int  `json:"code"`  //返回状态码200=ok
	Validate bool `json:"validate"`  //是否合法
	Token string `json:"token"`   //返回当前用户token
	Role string `json:"role"`  //用户角色

}