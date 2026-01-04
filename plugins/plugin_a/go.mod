module github.com/myproject/plugins/plugin_a

go 1.22

require github.com/myproject/corelib v0.0.0

require github.com/myproject/myinterface v0.0.0 // indirect

replace github.com/myproject/corelib => ../../corelib

replace github.com/myproject/myinterface => ../../myinterface
