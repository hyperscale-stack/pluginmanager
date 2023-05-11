.PHONY: all clean test cover travis lint

release:
	@echo "Release v$(version)"
	@git pull
	@git checkout master
	@git pull
	@git checkout develop
	@git flow release start $(version)
	@git flow release finish $(version) -p -m "Release v$(version)"
	@git checkout develop
	@echo "Release v$(version) finished."

all: test

clean:
	@go clean -i ./...

test: test.plugin.so $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go test -timeout 30s ./...

lint:
	@golangci-lint run

test.plugin.so: $(shell find . -type f -print | grep -v vendor | grep "\.go")
	@go build -buildmode=plugin -o ./test.plugin.so ./plugin/test
