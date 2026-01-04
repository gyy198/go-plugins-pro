package plugin_a

import (
	"fmt"
	//ip "github.com/myproject/myinterface"
	"github.com/myproject/corelib/plugins"
	"github.com/myproject/corelib/utils"
)

type PluginA struct{}

func (p PluginA) Name() string { return "PluginA" }
func (p PluginA) Run() {
	fmt.Println("PluginA running. Utils test:", utils.ToUpper("plugin a"))
}

func init() {
	plugins.Register(PluginA{})
}
