package main

import (
	"github.com/unknowname/prome-auto/hicicd"
	"github.com/unknowname/prome-auto/libs"
	"github.com/unknowname/prome-auto/prome"
	"github.com/zieckey/goini"
	"os"
	"time"
	"log"
)

const path = "/api/v1/query?query=istio_request_count"

func main() {
	log.Print("程序启动...")
	var latestMd5 string
	args := os.Args[:]
	if !libs.CheckParams(args) {
		return
	}
	confFile := os.Args[1]
	latestMd5 = newRun(confFile, latestMd5)
	for {
		currMd5 := newRun(confFile, latestMd5)
		latestMd5 = currMd5
		time.Sleep(time.Second * 15)
	}
}

func newRun(confFile, latestMd5 string) (md5 string) {
	//初始化一些变量待用
	serviceData := prome.NewServicesData()
	ini := goini.New()
	err := ini.ParseFile(confFile)
	if err != nil {
		log.Print("读取配置文件失败！请确认配置文件合法！\n")
		return
	}
	prometheusHost, _ := ini.SectionGet("prometheus", "host")
	hicicdHost, _ := ini.SectionGet("hicicd", "host")
	username, _ := ini.SectionGet("hicicd", "username")
	password, _ := ini.SectionGet("hicicd", "password")
	if libs.CheckHost(prometheusHost) && libs.CheckHost(hicicdHost) {
		if err := prome.GetServicesData(prometheusHost+path, serviceData); err == nil {
			confStr := prome.CreateRuleConfig(*serviceData)
			if libs.GetMd5(confStr) != latestMd5 {
				log.Print("检测到Prometheus API接口有新数据，执行更新ConfgMap动作\n")
				if token, err := hicicd.Login(hicicdHost, username, password); err == nil {
					if err := hicicd.CreateConf(token, hicicdHost, "", "", confStr); err == nil {
						log.Print("更新ConfigMap成功\n睡眠30秒，稍后重载Prometheus\n")
						time.Sleep(time.Second * 30)
						if err := prome.ReloadServer(prometheusHost); err == nil {
							log.Print("重载Prometheus成功\n")
							md5 = libs.GetMd5(confStr)
						} else {
							log.Print("重载Prometheus失败！", err, "\n")
						}
					} else {
						log.Print("创建ConfgMap失败 ", err, "\n")
					}
				} else {
					log.Print("获取创建ConfigMap的Token失败，请确保hicicd服务正常。", err, "\n")
					time.Sleep(time.Second * 20)
				}
			} else {
				//log.Print("未检测到有新指标可加入监控,休眠15秒后再监测\n")
				md5 = latestMd5
			}
		} else {
			log.Print("获取接口数据异常，原因：", err, "\n")
		}

	} else {
		log.Print("配置文件Prometheus host格式有误，请以http://或者hppts://为前缀\n")
	}
	return md5
}
