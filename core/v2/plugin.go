package v2

type Plugin interface {
	Name() string
	Init(ctx Context) error
	Dependencies() []string
	Execute() error
	Version() string // breaking change: new required method
}

type Context interface {
	GetPlugin(name string) (Plugin, error)
}