package hicicd

//ConfigMap接口要求的数据结构
type PostData struct {
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}

type Token struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

