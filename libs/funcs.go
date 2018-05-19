package libs

import (
	"github.com/unknowname/prome-auto/prome"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

const path = "configMap/add"

//ConfigMap接口要求数据结构
type PostData struct {
	Name      string            `json:"name"`
	NameSpace string            `json:"namespace"`
	Data      map[string]string `json:"data"`
}

//传入待JSON化的struct数据，转换成HTTP POST Data
func ToHttpData(i interface{}) (buffer *bytes.Buffer, err error) {
	jsonByte, err := json.Marshal(i)
	if err == nil {
		buffer = bytes.NewBuffer(jsonByte)
	}
	return buffer, err
}

//访问Prometheus的指标接口，获取所有指标
func GetPromeItems(host string) (datas []string) {
	data := prome.ItemData{}
	url := host + prome.ApiAddr
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		fmt.Print("Error ", err)
	}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error", err)
	}
	if err := json.Unmarshal(respByte, &data); err == nil {
		if data.Status == "success" {
			datas = data.Data
		}
	}
	return datas
}