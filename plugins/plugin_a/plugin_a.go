package plugin_a

import (
	"fmt"
	//ip "github.com/gyy198/go-plugins-pro/myinterface"
	"github.com/gyy198/go-plugins-pro/corelib/plugins"
	"github.com/gyy198/go-plugins-pro/corelib/utils"
)

type PluginA struct{}

func (p PluginA) Name() string { return "PluginA" }
func (p PluginA) Run() {
	fmt.Println("PluginA running. Utils test:", utils.ToUpper("plugin a"))
}

func init() {
	plugins.Register(PluginA{})
}
