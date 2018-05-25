package prome

import (
	"fmt"
	"github.com/unknowname/prome-auto/libs"
	"gopkg.in/yaml.v2"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

const exprBase = `changes(istio_request_count{destination_service="%s",destination_version="%s",response_code="%s"}[1m])` +
	` != 0`

//Create Prometheus AlertRule string
func CreateRuleConfig(data ServicesData) string {
	groups := Groups{}
	group := Group{}
	group.Name = "Prometheues Alert Rule"
	for _, v := range data.Data.Result {
		item := v.Metric.DestinationService
		appVersion := v.Metric.DestinationVersion
		project, app := libs.ParseProjectApp(item)
		httpCode := v.Metric.ResponseCode
		alertName := fmt.Sprintf("%s_%s_%s_%s", app, project, appVersion, httpCode)
		if strings.HasPrefix(httpCode, "2") || strings.HasPrefix(httpCode, "3") {
			continue
		}
		expr := fmt.Sprintf(exprBase, item, appVersion, httpCode)
		annot := make(map[string]string)
		labels := make(map[string]string)
		labels["project"] = project
		labels["app"] = app
		annot["desc"] = fmt.Sprintf("%s项目的%s应用%s版本,发现%s响应", project, app, appVersion, httpCode)
		rule := NewRule(alertName, expr, labels, annot)
		group.AddRule(*rule)
	}
	groups.AddGroup(group)
	confByte, _ := yaml.Marshal(groups)
	return string(confByte)
}

//Init a New alert rule
func CreateRule(item string) *Rule {
	labe := make(map[string]string)
	annot := make(map[string]string)
	project, app := libs.ParseProjectApp(item)
	labe["project"] = project
	labe["app"] = app
	return &Rule{
		Alert:       item,
		Expr:        item,
		Labels:      labe,
		Annotations: annot,
	}
}

//Access API and return struct data
func GetServicesData(url string, i *ServicesData) error {
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

//Reload Prometheus
func ReloadServer(server string) error {
	reloadPath := "/-/reload"
	_, err := http.Post(server+reloadPath, "", bytes.NewBuffer([]byte("")))
	return err
}
