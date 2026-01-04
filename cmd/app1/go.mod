module github.com/myproject/cmd/app1

go 1.22

require (
	github.com/myproject/corelib v0.0.0
	github.com/myproject/plugins/plugin_a v0.0.0
	github.com/myproject/plugins/plugin_b v0.0.0
)

require github.com/myproject/myinterface v0.0.0 //indirect

replace github.com/myproject/corelib => ../../corelib

replace github.com/myproject/myinterface => ../../myinterface

replace github.com/myproject/plugins/plugin_a => ../../plugins/plugin_a

replace github.com/myproject/plugins/plugin_b => ../../plugins/plugin_b
