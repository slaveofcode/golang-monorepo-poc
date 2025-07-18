package core

type Plugin interface {
	Name() string
	Init(ctx Context) error
	Dependencies() []string
	Execute() error
}

type Context interface {
	GetPlugin(name string) (Plugin, error)
}