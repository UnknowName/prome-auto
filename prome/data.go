package prome

const ApiAddr = "/api/v1/label/__name__/values"

//Prometheus Healthy API响应结构体
type ItemData struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type ServicesData struct {
	Status  string `json:"status"`
	Data Result     `json:"data"`
}

type Result struct {
	ResultType string `json:"resultType"`
	Result []Metric `json:"result"`
}

type Metric struct {
	DestinationService string `json:"destination_service"`
	ResponseCode       string `json:"response_code"`
	DestinationVersion string `json:"destination_version"`
}
