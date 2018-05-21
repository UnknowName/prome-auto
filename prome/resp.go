package prome

//Prometheus 所有指标API响应结构体
type ItemData struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

//Prometheus HTTP状态统计API响应结构体
type ServicesData struct {
	Status  string `json:"status"`
	Data Result     `json:"data"`
}

type Result struct {
	ResultType string `json:"resultType"`
	Result []Results `json:"result"`
}

type Results struct {
	Metric  Metric `json:"metric"`
}

type Metric struct {
	DestinationService string `json:"destination_service"`
	ResponseCode       string `json:"response_code"`
	DestinationVersion string `json:"destination_version"`
}

func NewItemData() *ItemData{
	return  &ItemData{}
}

func NewServicesData() *ServicesData {
	return &ServicesData{}
}
