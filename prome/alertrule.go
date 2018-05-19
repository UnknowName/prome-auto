package prome

//Prometheus AlertRule配置文件结构体
type Rule struct {
	Alert       string `yaml:"alert"`
	Expr        string  `yaml:"expr"`
	Labels      map[string]string  `yaml:"labels"`
	Annotations map[string]string   `yaml:"annotations"`
}

type Group struct {
	Name  string
	Rules []Rule
}

type Groups struct {
	Groups []Group
}

func (g *Group) AddRule(rule Rule) {
	g.Rules = append(g.Rules, rule)
}