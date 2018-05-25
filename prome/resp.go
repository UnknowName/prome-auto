package prome

//Prometheus ALL API response
type ItemData struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

//Prometheus HTTP API response
type ServicesData struct {
	Status string `json:"status"`
	Data   Result `json:"data"`
}

type Result struct {
	ResultType string    `json:"resultType"`
	Result     []Results `json:"result"`
}

type Results struct {
	Metric Metric `json:"metric"`
}

type Metric struct {
	DestinationService string `json:"destination_service"`
	ResponseCode       string `json:"response_code"`
	DestinationVersion string `json:"destination_version"`
}

func NewServicesData() *ServicesData {
	return &ServicesData{}
}
