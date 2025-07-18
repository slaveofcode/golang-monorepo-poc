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

## Prove to Use for V1

You can test those packages by creating a simple golang project with `main.go` file, and write the example code below.

```go
package main

import (
	"fmt"
	"github.com/slaveofcode/golang-monorepo-poc/core"
	"github.com/slaveofcode/golang-monorepo-poc/plugin-a"
	"github.com/slaveofcode/golang-monorepo-poc/plugin-b"
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

## Next: For Example I Added Core V2 with Breaking Changes

A breaking changes for core plugin is added, the old codebase that refers to plugin-a, plugin-b and core v1 should not broken. plugin-c is added to use the core v2 as an implementation, so core now should support bot v1 & v2 codebase.

Tag the core v2 & plugin c, make sure no changes in core v1, plugin-a & plugin-b. If anything changes on them, we should tag them too to maintain consistency.
```
git tag core/v2/v2.0.0
git tag plugin-c/v1.0.0
```

You can test by creating another project, and puth these code to try

```go
package main

import (
	"fmt"
	"github.com/slaveofcode/golang-monorepo-poc/core"
	"github.com/slaveofcode/golang-monorepo-poc/core/v2"
	"github.com/slaveofcode/golang-monorepo-poc/plugin-a"
	"github.com/slaveofcode/golang-monorepo-poc/plugin-b"
	"github.com/slaveofcode/golang-monorepo-poc/plugin-c"
)

func main() {
	// plugin-a & plugin-b which using v1 core is still supported here
	registryV1 := map[string]core.Plugin{
		"plugin-b": &pluginb.PluginB{},
		"plugin-a": &plugina.PluginA{},
	}

	ctxV1 := core.NewContext(registryV1)

	for _, name := range []string{"plugin-b", "plugin-a"} {
		if err := registryV1[name].Init(ctxV1); err != nil {
			panic(err)
		}
	}

	fmt.Println("Executing plugin-a:")
	if err := registryV1["plugin-a"].Execute(); err != nil {
		panic(err)
	}

	// plugin-c is using newest core/v2
	registryV2 := map[string[corev2.Plugin]{
		"plugin-c": &pluginc.PluginC{},
	}}

	ctxV2 := corev2.NewContext(registryV2)
	if err := registryV2["plugin-c"].Init(ctxV2); err != nil {
		panic(err)
	}

	fmt.Println("Executing plugin-c:")
	if err := registryV2["plugin-c"].Execute(); err != nil {
		panic(err)
	}
}
```