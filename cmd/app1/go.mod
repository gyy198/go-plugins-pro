module github.com/gyy198/go-plugins-pro.git/cmd/app1

go 1.22

require (
	github.com/gyy198/go-plugins-pro.git/corelib v0.0.0
	github.com/gyy198/go-plugins-pro.git/plugins/plugin_a v0.0.0
	github.com/gyy198/go-plugins-pro.git/plugins/plugin_b v0.0.0
)

require github.com/gyy198/go-plugins-pro.git/myinterface v0.0.0 //indirect

replace github.com/gyy198/go-plugins-pro.git/corelib => ../../corelib

replace github.com/gyy198/go-plugins-pro.git/myinterface => ../../myinterface

replace github.com/gyy198/go-plugins-pro.git/plugins/plugin_a => ../../plugins/plugin_a

replace github.com/gyy198/go-plugins-pro.git/plugins/plugin_b => ../../plugins/plugin_b
