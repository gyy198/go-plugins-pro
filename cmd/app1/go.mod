module github.com/gyy198/go-plugins-pro/cmd/app1

go 1.22

require (
	github.com/gyy198/go-plugins-pro/corelib v0.0.3
	github.com/gyy198/go-plugins-pro/myinterface v0.0.3 //indirect
	github.com/gyy198/go-plugins-pro/plugins/plugin_a v0.0.3
	github.com/gyy198/go-plugins-pro/plugins/plugin_b v0.0.3
)

// replace github.com/gyy198/go-plugins-pro/corelib => ../../corelib

// replace github.com/gyy198/go-plugins-pro/myinterface => ../../myinterface

// replace github.com/gyy198/go-plugins-pro/plugins/plugin_a => ../../plugins/plugin_a

// replace github.com/gyy198/go-plugins-pro/plugins/plugin_b => ../../plugins/plugin_b
