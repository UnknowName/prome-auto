package prome

import (
	"testing"
	"fmt"
	"gopkg.in/yaml.v2"
)

func TestCheng(t *testing.T){
	label := make(map[string]string)
	label["desc"] = "{{ value }}"
	label["item"] = "item"
    rule := Rule{
    	Alert:"test rule",
    	Expr:"go_threads() > 10",
    	Labels:label,
    	Annotations:map[string]string{},
    }
    group := Group{
    	Name:"localhost",
	}
	group.AddRule(rule)
	groups := Groups{[]Group{group}}
	d, err := yaml.Marshal(&groups)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(string(d))
}