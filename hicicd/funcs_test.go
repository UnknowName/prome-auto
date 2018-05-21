package hicicd

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/widuu/goini"
)

func TestLogin(t *testing.T) {
	host := "http://localhost:8081"
	token, err := Login(host, "chengjianneng", "1234qwer")
	fmt.Print("token is ",token,"\n")
	fmt.Print("err is", err, "\n")
}

func TestCreateConf(t *testing.T) {
	server := "http://localhost:8081"
	token,err := Login(server, "chengjianneng", "1234qwer")
	name, namespace := "prometheus-rule", ""
	confStr := "fffs"
	err = CreateConf(token, server, name, namespace, confStr)
	assert.Equal(t, err, nil)
}

func TestSetValue(t *testing.T) {
	confFile := "./conf.ini"
	conf := goini.SetConfig(confFile)
	conf.SetValue("prometheus", "md5", "md5value")
	value := conf.GetValue("prometheus", "host")
	fmt.Print(value)

}