package prome

import (
	"fmt"
	"github.com/unknowname/prome-auto/libs"
	"gopkg.in/yaml.v2"
	"strings"
)

const exprBase = `changes(istio_request_count{destination_service="%s",destination_version="%s",response_code="%s"}[1m])` +
                 ` != 0`

//创建Prometheus的AlertRule配置文件字符串
func CreateRuleConfig(data ServicesData) string {
	groups := Groups{}
	group := Group{}
	labels := make(map[string]string)
	annot := make(map[string]string)
	for _, v := range data.Data.Result {
		item := v.Metric.DestinationService
		httpCode := v.Metric.ResponseCode
		project, app := libs.ParseProjectApp(item)
		labels["project"] = project
		labels["app"] = app
		if strings.HasPrefix(httpCode, "2") || strings.HasPrefix(httpCode, "3") {
			continue
		}
		expr := fmt.Sprintf(exprBase, item, v.Metric.DestinationVersion, httpCode)
		rule := NewRule(item, expr, labels, annot)
		group.AddRule(*rule)
	}
	groups.AddGroup(group)
	confByte, _ := yaml.Marshal(groups)
	fmt.Print(string(confByte))
	return string(confByte)
}
