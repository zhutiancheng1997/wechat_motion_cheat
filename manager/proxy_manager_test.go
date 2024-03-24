package manager

import (
	"testing"
)

func TestDoGetDy1000Play(t *testing.T) {
	manager := &ProxyManager{}

	proxy := manager.getProxy()

	manager.DoGetDy1000Play(proxy)
}
