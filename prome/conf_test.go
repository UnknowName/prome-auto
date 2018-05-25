package prome

import (
	"testing"
	"fmt"
	"github.com/magiconair/properties/assert"
)

func TestCreateRuleConfig(t *testing.T) {
	fmt.Print("ok")
}

func TestReloadServer(t *testing.T) {
	server := "http://prometheus-istio-system.apps.oc.com"
	err := ReloadServer(server)
	assert.Equal(t, err, nil)
}
