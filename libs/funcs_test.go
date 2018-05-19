package libs

import (
	"testing"
	"fmt"
)

func TestGetPromeDatas(t *testing.T) {
	host := "http://prometheus-istio-system.apps.oc.com"
	datas := GetPromeItems(host)
	for _,item :=  range datas {
		if FilterItem(item) != "" {
			fmt.Println(item, "\n")
		}
	}
}

func TestToHttpData(t *testing.T) {
	data := make(map[string]string)
	data["name"] = "cheng"
	postData := &PostData{
		Name:`test.yaml`,
		NameSpace:"istio-system",
		Data:data,
	}
	jsonStr,err := ToHttpData(postData)
	if err == nil {
		fmt.Print(jsonStr)
	}
}
