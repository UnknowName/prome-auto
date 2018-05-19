package libs

import (
	"testing"
	"fmt"
)

func TestFilterItem(t *testing.T) {
	item := "envoy_cluster_out_demo_provider_demo_dev_svc_cluster_local_http_7575_membership_healthy"
	fmt.Print(FilterItem(item))
}

func TestGetItemAttr(t *testing.T) {
	item := "envoy_cluster_out_demo_provider_demo_dev_svc_cluster_local_http_7575_membership_healthy"
	fmt.Print(GetItemAttr(item))
}