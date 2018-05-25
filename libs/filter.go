package libs

import (
	"regexp"
)

//Filter  end with healthy item
func FilterItem(item string) string{
	reg := regexp.MustCompile(`\w.*http.*healthy$`)
	if reg.MatchString(item){
		return item
	}
	return ""
}

//Parse the project and app from item
func ParseProjectApp(item string) (project, app string) {
	reg := regexp.MustCompile(`(.*)\.(.*).svc.*`)
	matchs := reg.FindStringSubmatch(item)
	matchsLength := len(matchs)
	if matchsLength < 3 {
		project, app = "", ""
	}else{
		project = matchs[matchsLength-1]
		app = matchs[matchsLength-2]
	}
	return project,app
}