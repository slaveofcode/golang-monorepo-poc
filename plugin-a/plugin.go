package plugina

import (
	"fmt"

	"github.com/slaveofcode/golang-monorepo-poc/core"
)

type PluginA struct {
	dep core.Plugin
}

func (p *PluginA) Name() string {
	return "plugin-a"
}

func (p *PluginA) Init(ctx core.Context) error {
	dep, err := ctx.GetPlugin("plugin-b")
	if err != nil {
		return err
	}
	p.dep = dep
	fmt.Println("[plugin-a] initialized")
	return nil
}

func (p *PluginA) Dependencies() []string {
	return []string{"plugin-b"}
}

func (p *PluginA) Execute() error {
	fmt.Println("[plugin-a] executing after...")
	return p.dep.Execute()
}