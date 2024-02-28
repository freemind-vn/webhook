package plugin

type Plugin interface {
	Name() string
	Description() string
	Version() string

	Init(config map[string]any)
	Run(payload map[string]any) error
}
