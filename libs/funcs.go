package libs

import (
	"encoding/json"
	"bytes"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/unknowname/prome-auto/prome"
	"strings"
)

//传入待JSON化的struct数据，转换成HTTP POST Data
func ToHttpData(i interface{}) (buffer *bytes.Buffer, err error) {
	jsonByte, err := json.Marshal(i)
	if err == nil {
		buffer = bytes.NewBuffer(jsonByte)
	}
	return buffer, err
}

//传入返回JSON的URL以及定义后的struct。返回JSON化后struct
func GetServicesData(url string, i *prome.ServicesData) error {
	resp, err := http.Get(url)
	if err == nil {
		respByte, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respByte, i)
		defer resp.Body.Close()
	} else {
		fmt.Print("获取URL失败 ", err)
	}
	return err
}

//检查用户提供的Prometheus服务器格式正确
func CheckHost(host string) bool {
	if strings.HasPrefix(host, "http://") || strings.HasPrefix(host, "https://") {
		return true
	}
	return false
}

//构造并初始化Rule
func CreateRule(item string) *prome.Rule {
	labe := make(map[string]string)
	annot := make(map[string]string)
	project, app := ParseProjectApp(item)
	labe["project"] = project
	labe["app"] = app
	return &prome.Rule{
		Alert:       item,
		Expr:        item,
		Labels:      labe,
		Annotations: annot,
	}
}
