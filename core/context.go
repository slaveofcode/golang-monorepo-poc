package core

import "fmt"

type pluginContext struct {
	plugins map[string]Plugin
}

func NewContext(registry map[string]Plugin) Context {
	return &pluginContext{plugins: registry}
}

func (c *pluginContext) GetPlugin(name string) (Plugin, error) {
	p, ok := c.plugins[name]
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", name)
	}
	return p, nil
}