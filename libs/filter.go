package libs

import (
	"regexp"
)

//过滤出带有healthy结尾的指标
func FilterItem(item string) string{
	reg := regexp.MustCompile(`\w.*http.*healthy$`)
	if reg.MatchString(item){
		return item
	}
	return ""
}

//通过正则，将应用名与项目名称一起提取出来
func GetItemAttr(itemStr string) string {
	reg := regexp.MustCompile(`([a-z]+_){3}(\w+)_svc`)
	matchs := reg.FindStringSubmatch(itemStr)
	if len(matchs) != 0 {
		return matchs[len(matchs)-1]
	}
	return ""
}