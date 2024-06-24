package plugins

import (
	"reflect"
	"strings"
)

type Plugins struct{}

// Collection is a struct that contains the name of the method and the method itself
// this is used to invoke the method if the command is found
// in the user input
type Collection struct {
	Name   string
	Method interface{}
}

// CollectionInfo is a struct that contains the name of the method and the description of the method
// this is used to display the list of available commands and their descriptions
// in the help command
type CollectionInfo struct {
	Name          string
	Description   string
	PositionIndex int
}

// Plugins returns a slice of Collection
// which contains the name of the method and the method itself
// this is used to invoke the method if the command is found
// in the user input
func (p *Plugins) Plugins() []Collection {
	var plugins []Collection
	pluginValue := reflect.ValueOf(p)
	pluginType := reflect.TypeOf(p)
	for i := 0; i < pluginType.NumMethod(); i++ {
		method := pluginType.Method(i)
		plugins = append(plugins, Collection{
			// strings.ToLower converts the string to lowercase
			// this is used to make the command case-insensitive
			Name: strings.ToLower(method.Name),
			// MethodByName returns a Value corresponding to the method with the given name
			// if the method is found, it will return a valid value
			// if the method is not found, it will return a zero Value
			Method: pluginValue.MethodByName(method.Name).Interface(),
		})
	}
	return plugins
}

// PluginsInfo returns a slice of CollectionInfo
// which contains the name of the method and the description of the method
// this is used to display the list of available commands and their descriptions
// in the help command
func (p *Plugins) PluginsInfo() []CollectionInfo {
	var pluginsInfo []CollectionInfo
	pluginValue := reflect.ValueOf(p)
	pluginType := reflect.TypeOf(p)
	for i := 0; i < pluginType.NumMethod(); i++ {
		method := pluginType.Method(i)
		// type assertion to check if the method is a function
		if fn, ok := pluginValue.MethodByName(method.Name).Interface().(func() CollectionInfo); ok {
			pluginsInfo = append(pluginsInfo, fn())
		}
	}
	return pluginsInfo
}
