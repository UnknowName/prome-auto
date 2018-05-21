package libs

import (
	"testing"
	"fmt"
	"github.com/magiconair/properties/assert"
)

func TestFilterItem(t *testing.T) {
	item := "envoy_cluster_out_demo_provider_demo_dev_svc_cluster_local_http_7575_membership_healthy"
	fmt.Print(FilterItem(item))
}

func TestGetItemAttr(t *testing.T) {
	item := "envoy_cluster_out_demo_provider_demo_dev_svc_cluster_local_http_7575_membership_healthy"
	fmt.Print(GetItemAttr(item))
}

func TestParseProjectApp(t *testing.T) {
	namespace,app := ParseProjectApp("details.tutorial.svc.cluster.local")
	assert.Equal(t, "tutorial", namespace)
	assert.Equal(t, app, "details")
	project,app := ParseProjectApp("demo-provider.demo-dev.svc.cluster.local")
	assert.Equal(t,project,"demo-dev")
	assert.Equal(t, app, "demo-provider")
}