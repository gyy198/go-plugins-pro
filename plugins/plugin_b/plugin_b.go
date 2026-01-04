package plugin_b

import (
	"fmt"
	//ip "github.com/gyy198/go-plugins-pro.git/myinterface"
	"github.com/gyy198/go-plugins-pro.git/corelib/plugins"
)

type PluginB struct{}

func (p PluginB) Name() string { return "PluginB" }
func (p PluginB) Run() {
	fmt.Println("PluginB running.")
}

func init() {
	plugins.Register(PluginB{})
}
