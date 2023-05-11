package pluginmanager

import (
	"fmt"
	"plugin"
)

type PluginManager[T any] struct {
	plugins map[string]T
}

func New[T any]() *PluginManager[T] {
	return &PluginManager[T]{
		plugins: make(map[string]T),
	}
}

// Register registers a plugin.
func (pm *PluginManager[T]) Register(plugin *Plugin[T]) error {
	name := plugin.ID()

	if _, exists := pm.plugins[name]; exists {
		return fmt.Errorf("plugin %s already exists", name)
	}

	pm.plugins[name] = plugin.Load()

	return nil
}

// RegisterFromFile registers a plugin from a file.
func (pm *PluginManager[T]) RegisterFromFile(path string) error {
	p, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("could not open plugin %s: %w", path, err)
	}

	register, err := p.Lookup("Register")
	if err != nil {
		return fmt.Errorf("could not find Register function in plugin %s: %w", path, err)
	}

	pi, err := register.(func() (*Plugin[T], error))()
	if err != nil {
		return fmt.Errorf("could not register plugin %s: %w", path, err)
	}

	return pm.Register(pi)
}

// Get returns a plugin by name.
func (pm *PluginManager[T]) Get(name string) (p T, err error) {
	if plugin, exists := pm.plugins[name]; exists {
		return plugin, nil
	}

	err = fmt.Errorf("plugin %s does not exist", name)

	return
}
