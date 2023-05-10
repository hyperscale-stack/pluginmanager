module github.com/hyperscale-stack/pluginmanager

go 1.20

require github.com/stretchr/testify v1.8.2

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/hyperscale-stack/pluginmanager/demo v1.8.2

replace github.com/hyperscale-stack/pluginmanager/demo => ./demo
