package plugin_a

import (
	"fmt"
	//ip "github.com/gyy198/go-plugins-pro.git/myinterface"
	"github.com/gyy198/go-plugins-pro.git/corelib/plugins"
	"github.com/gyy198/go-plugins-pro.git/corelib/utils"
)

type PluginA struct{}

func (p PluginA) Name() string { return "PluginA" }
func (p PluginA) Run() {
	fmt.Println("PluginA running. Utils test:", utils.ToUpper("plugin a"))
}

func init() {
	plugins.Register(PluginA{})
}
