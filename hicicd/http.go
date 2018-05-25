package hicicd

//OpenShift ConfigMap
type ConfMap struct {
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}

//Hicicd Login
type Token struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Hicicd Response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
