package libs

import (
	"testing"
	"fmt"
	"github.com/magiconair/properties/assert"
)

func TestGetMd5(t *testing.T) {
	originStr := "bc92c0617d9434fa337d5dd0f2e9ee42"
	conf := `groups:
- name: Prometheues Alert Rule
  rules:
  - alert: demo-consumer_demo-dev_v1_404
    expr: changes(istio_request_count{destination_service="demo-consumer.demo-dev.svc.cluster.local",destination_version="v1",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的demo-consumer应用v1版本,发现404响应
      value: '{{ $value }}'
  - alert: demo-consumer_demo-dev_v2_401
    expr: changes(istio_request_count{destination_service="demo-consumer.demo-dev.svc.cluster.local",destination_version="v2",response_code="401"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的demo-consumer应用v2版本,发现401响应
      value: '{{ $value }}'
  - alert: demo-consumer_demo-dev_v2_404
    expr: changes(istio_request_count{destination_service="demo-consumer.demo-dev.svc.cluster.local",destination_version="v2",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的demo-consumer应用v2版本,发现404响应
      value: '{{ $value }}'
  - alert: demo-provider_demo-dev_v1_404
    expr: changes(istio_request_count{destination_service="demo-provider.demo-dev.svc.cluster.local",destination_version="v1",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的demo-provider应用v1版本,发现404响应
      value: '{{ $value }}'
  - alert: demo-provider_demo-dev_v2_404
    expr: changes(istio_request_count{destination_service="demo-provider.demo-dev.svc.cluster.local",destination_version="v2",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的demo-provider应用v2版本,发现404响应
      value: '{{ $value }}'
  - alert: details_tutorial_v1_404
    expr: changes(istio_request_count{destination_service="details.tutorial.svc.cluster.local",destination_version="v1",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: tutorial项目的details应用v1版本,发现404响应
      value: '{{ $value }}'
  - alert: hello-world_demo-dev_v1_404
    expr: changes(istio_request_count{destination_service="hello-world.demo-dev.svc.cluster.local",destination_version="v1",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的hello-world应用v1版本,发现404响应
      value: '{{ $value }}'
  - alert: hello-world_demo-dev_v2_404
    expr: changes(istio_request_count{destination_service="hello-world.demo-dev.svc.cluster.local",destination_version="v2",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的hello-world应用v2版本,发现404响应
      value: '{{ $value }}'
  - alert: hello-world_demo-dev_v3_404
    expr: changes(istio_request_count{destination_service="hello-world.demo-dev.svc.cluster.local",destination_version="v3",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: demo-dev项目的hello-world应用v3版本,发现404响应
      value: '{{ $value }}'
  - alert: productpage_tutorial_v1_404
    expr: changes(istio_request_count{destination_service="productpage.tutorial.svc.cluster.local",destination_version="v1",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: tutorial项目的productpage应用v1版本,发现404响应
      value: '{{ $value }}'
  - alert: reviews_tutorial_v3_404
    expr: changes(istio_request_count{destination_service="reviews.tutorial.svc.cluster.local",destination_version="v3",response_code="404"}[1m])
      != 0
    labels:
      app: reviews
      project: tutorial
    annotations:
      desc: tutorial项目的reviews应用v3版本,发现404响应
      value: '{{ $value }}'
`
	str := GetMd5(conf)
	fmt.Print(str)
	assert.Equal(t,str,originStr)
}