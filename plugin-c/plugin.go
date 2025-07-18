package pluginc

import (
	"fmt"

	corev2 "github.com/slaveofcode/golang-monorepo-poc/core/v2"
)

type PluginC struct{}

func (p *PluginC) Name() string {
	return "plugin-c"
}

func (p *PluginC) Init(ctx corev2.Context) error {
	fmt.Println("[plugin-c] initialized with core v2")
	return nil
}

func (p *PluginC) Dependencies() []string {
	return nil
}

func (p *PluginC) Execute() error {
	fmt.Println("[plugin-c] executed")
	return nil
}

func (p *PluginC) Version() string {
	return "2.0.0"
}