package main

import (
	"github.com/unknowname/prome-auto/libs"
	"fmt"
	"github.com/unknowname/prome-auto/prome"
	"gopkg.in/yaml.v2"
)

const host = "http://prometheus-istio-system.apps.oc.com"

func main() {
	dict := make(map[string]string)
	group := prome.Group{}
	group.Name = "Prometheus Alert by Auto Create"
	groups := prome.Groups{}
	dict["value"] = `{{ $value }}`
	datas := libs.GetPromeItems(host)
	for _, item := range datas {
		if libs.FilterItem(item) != "" {
			appAndNamespace := libs.GetItemAttr(item)
			alert := prome.Rule{
				Alert:       appAndNamespace,
				Expr:        fmt.Sprintf("(%s) != 1", item),
				Annotations: dict,
				Labels:      dict,
			}
			group.AddRule(alert)
		}
	}
	groups.Groups = []prome.Group{group}
	d, err := yaml.Marshal(&groups)
	if err != nil {
		fmt.Print(err)
	}
	yamlStr := string(d)
	fmt.Print(yamlStr)
}
