package pluginmanager

type Plugin[T any] struct {
	Name     string
	Instance T
}

func (p *Plugin[T]) ID() string {
	return p.Name
}

func (p *Plugin[T]) Load() T {
	return p.Instance
}
