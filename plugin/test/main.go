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

func Register() (*pluginmanager.Plugin[plugin.Authentification], error) {

	return &pluginmanager.Plugin[plugin.Authentification]{
		Name:     "test",
		Instance: &MyPlugin{},
	}, nil
}
