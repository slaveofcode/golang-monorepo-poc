# Golang Monorepo Prove of Concept

This is just POC of golang multi package support using single repository as a monorepo.

## Publish Tags

Here we create git tag using prefix of the submodule (core, plugin-a, plugin-b), followed by the semantic versioning.

```
git tag core/v1.0.0
git tag plugin-a/v1.0.0
git tag plugin-b/v1.0.0
git push --tags
```

## Prove to Use

You can test those packages by creating a simple golang project with `main.go` file, and write the example code below.

```go
package main

import (
	"fmt"
	"github.com/your-org/plugins/core"
	"github.com/your-org/plugins/plugin-a"
	"github.com/your-org/plugins/plugin-b"
)

func main() {
	registry := map[string]core.Plugin{
		"plugin-b": &pluginb.PluginB{},
		"plugin-a": &plugina.PluginA{},
	}

	ctx := core.NewContext(registry)

	// Init all plugins in dependency order
	for _, name := range []string{"plugin-b", "plugin-a"} {
		if err := registry[name].Init(ctx); err != nil {
			panic(err)
		}
	}

	// Execute top plugin
	fmt.Println("Executing plugin-a:")
	if err := registry["plugin-a"].Execute(); err != nil {
		panic(err)
	}
}
```