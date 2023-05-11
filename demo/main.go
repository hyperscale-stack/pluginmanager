package main

import (
	"fmt"

	"github.com/hyperscale-stack/pluginmanager"
	"github.com/hyperscale-stack/pluginmanager/demo/pkg/plugin"
)

func main() {
	pm := pluginmanager.New[plugin.Authentification]()

	err := pm.RegisterFromFile("./test.plugin.so")
	if err != nil {
		panic(err)
	}

	p, err := pm.Get("test")
	if err != nil {
		panic(err)
	}

	if ok, err := p.Authentificate("test", "test"); !ok || err != nil {
		panic("Authentification failed")
	}

	fmt.Println("Authentification success")
}
