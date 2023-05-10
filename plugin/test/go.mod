module github.com/hyperscale-stack/pluginmanager/plugin/test

go 1.20

require (
	github.com/hyperscale-stack/pluginmanager v0.0.0-20190821144910-3b4c3b2b1b0b
	github.com/hyperscale-stack/pluginmanager/demo v0.0.0-20190821144910-3b4c3b2b1b0b
)

replace github.com/hyperscale-stack/pluginmanager => ../../

replace github.com/hyperscale-stack/pluginmanager/demo => ../../demo
