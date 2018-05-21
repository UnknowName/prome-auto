package hicicd

//ConfigMap接口要求的数据结构
type ConfMap struct {
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}

//Login数据结构
type Token struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//服务响应数据结构
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
