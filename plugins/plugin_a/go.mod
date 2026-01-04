module github.com/gyy198/go-plugins-pro.git/plugins/plugin_a

go 1.22

require github.com/gyy198/go-plugins-pro.git/corelib v0.0.0

require github.com/gyy198/go-plugins-pro.git/myinterface v0.0.0 // indirect

replace github.com/gyy198/go-plugins-pro.git/corelib => ../../corelib

replace github.com/gyy198/go-plugins-pro.git/myinterface => ../../myinterface
