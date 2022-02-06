.PHONY: build run debug build-plugin

DIR=$(CURDIR)
PLUGINDIR=$(DIR)/plugins

build:
	CGO_ENABLE=0 go build -o build/redis-dashboard $(DIR)/cmd/redis-dashboard/.

run:
	build/redis-dashboard --plugin-folder=$(PLUGINDIR)

debug: build build-plugin run

build-plugin:
	CGO_ENABLE=0 go build -buildmode=plugin -o $(PLUGINDIR)/traditional_chinese.so $(DIR)/cmd/plugins/traditional_chinese/.