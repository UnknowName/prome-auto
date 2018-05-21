package hicicd

import "fmt"

func Login(user, password string) string  {

}

func CreateConf(server string,data PostData) {
	path := "/configMap/add"
	//还要获取Token!!
	fmt.Print(server + path)

}

func UpdateConf(){

}

func MountConf(){

}
