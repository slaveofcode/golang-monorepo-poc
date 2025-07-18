package pluginb

import (
	"fmt"

	"github.com/slaveofcode/golang-monorepo-poc/core"
)

type PluginB struct{}

func (p *PluginB) Name() string {
	return "plugin-b"
}

func (p *PluginB) Init(ctx core.Context) error {
	fmt.Println("[plugin-b] initialized")
	return nil
}

func (p *PluginB) Dependencies() []string {
	return nil
}

func (p *PluginB) Execute() error {
	fmt.Println("[plugin-b] executed")
	return nil
}
