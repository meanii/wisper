package plugins

import (
	"reflect"
	"strings"
)

type Plugins struct{}

type Collection struct {
	Name   string
	Method interface{}
}

func (p *Plugins) Plugins() []Collection {
	var plugins []Collection
	pluginValue := reflect.ValueOf(p)
	pluginType := reflect.TypeOf(p)
	for i := 0; i < pluginType.NumMethod(); i++ {
		method := pluginType.Method(i)
		plugins = append(plugins, Collection{
			Name:   strings.ToLower(method.Name),
			Method: pluginValue.MethodByName(method.Name).Interface(),
		})
	}
	return plugins
}
