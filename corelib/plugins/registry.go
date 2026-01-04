package plugins

import ip "github.com/myproject/myinterface"

var pluginMap = make(map[string]ip.Plugin)

func Register(p ip.Plugin) {
	pluginMap[p.Name()] = p
}

func All() []ip.Plugin {
	list := []ip.Plugin{}
	for _, p := range pluginMap {
		list = append(list, p)
	}
	return list
}
