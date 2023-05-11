package pluginmanager_test

import (
	"testing"

	"github.com/hyperscale-stack/pluginmanager"
	plug "github.com/hyperscale-stack/pluginmanager/demo/pkg/plugin"
	"github.com/stretchr/testify/assert"
)

type Authentification interface {
	Authentificate(username, password string) (bool, error)
}

var _ Authentification = (*MyPlugin)(nil)

// plugin test
type MyPlugin struct {
}

func (p *MyPlugin) Authentificate(username, password string) (bool, error) {
	return true, nil
}

func TestPluginManager(t *testing.T) {
	pm := pluginmanager.New[Authentification]()

	plugin := &pluginmanager.Plugin[Authentification]{
		Name:     "test",
		Instance: &MyPlugin{},
	}

	err := pm.Register(plugin)
	assert.NoError(t, err)

	p, err := pm.Get("test")
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Implements(t, (*Authentification)(nil), p)
	assert.Equal(t, p, plugin.Instance)

	p, err = pm.Get("test2")
	assert.Error(t, err)
	assert.Nil(t, p)

	err = pm.Register(plugin)
	assert.Error(t, err)
}

func TestPluginManagerWithFile(t *testing.T) {
	pm := pluginmanager.New[plug.Authentification]()

	err := pm.RegisterFromFile("test.plugin.so")
	assert.NoError(t, err)

	p, err := pm.Get("test")
	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.Implements(t, (*plug.Authentification)(nil), p)

	p, err = pm.Get("test2")
	assert.Error(t, err)
	assert.Nil(t, p)

	err = pm.RegisterFromFile("test.plugin.so")
	assert.Error(t, err)
}
