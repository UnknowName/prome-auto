package libs

import (
	"testing"
	"github.com/unknowname/prome-auto/prome"
	"fmt"
)

func TestGetJsonData(t *testing.T) {
	url := "http://prometheus-istio-system.apps.oc.com/api/v1/query?query=changes(istio_request_count{}[15m])"
	s := &prome.ServicesData{}
	err := GetServicesData(url,s)
	if err != nil {
		fmt.Print(err)
	}
	for _,v := range s.Data.Result {
		if v.Metric.ResponseCode == "200" {
			continue
		}
		fmt.Print(v.Metric.DestinationService, v.Metric.ResponseCode, v.Metric.DestinationVersion, "\n")
	}
}