package main

import (
	"os"
	"github.com/unknowname/prome-auto/libs"
	"github.com/unknowname/prome-auto/prome"
	"fmt"
	"time"
)

const path  = "/api/v1/query?query=changes(istio_request_count{}[15m])"

func main() {
	//latestConf从文件中读取。key为文件名，value为文件内容
	var latestConf string
	for {
		run(latestConf)
		time.Sleep(time.Second * 5)
	}
}

func run(conf string)  {
	fmt.Print("Start Programing...\n")
    serviceData := prome.NewServicesData()
    args := os.Args[:]
    if len(args) != 2 {
    	fmt.Print("请提供Prometheus服务器地址参数！\n")
    	os.Exit(5)
	}
	prometheusHost := args[1]
    if libs.CheckHost(prometheusHost) {
    	if err := libs.GetServicesData(prometheusHost+path, serviceData);err == nil {
    		fmt.Print("OK")
    		confStr := prome.CreateRuleConfig(*serviceData)
    		if conf == "" {
    			fmt.Print("新建ConfigMap")
			}
    		fmt.Print(confStr)
    		//如果confStr与上一次不一致，说明新的服务加入进来。需要修改configMap
		}else {
			fmt.Print("错误，获取接口数据异常。15秒后将进行重试\n")
		}
	}else {
		fmt.Print("请提供正确的Prometheus服务器地址。如 http://localhost:9090\n")
		os.Exit(5)
	}
}
