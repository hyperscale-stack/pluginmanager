package main

import (
	"github.com/hyperscale-stack/pluginmanager"
	"github.com/hyperscale-stack/pluginmanager/demo/pkg/plugin"
)

var _ plugin.Authentification = (*MyPlugin)(nil)

type MyPlugin struct {
}

func (p *MyPlugin) Authentificate(username, password string) (bool, error) {
	return true, nil
}

func Register() (*pluginmanager.Plugin[*MyPlugin], error) {

	return &pluginmanager.Plugin[*MyPlugin]{
		Name:     "test",
		Instance: &MyPlugin{},
	}, nil
}
