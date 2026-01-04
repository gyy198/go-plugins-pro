package plugin_b

import (
	"fmt"
	//ip "github.com/myproject/myinterface"
	"github.com/myproject/corelib/plugins"
)

type PluginB struct{}

func (p PluginB) Name() string { return "PluginB" }
func (p PluginB) Run() {
	fmt.Println("PluginB running.")
}

func init() {
	plugins.Register(PluginB{})
}
