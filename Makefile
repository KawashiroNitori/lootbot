all: build

generate:
	go generate ./internal/... ./ent/

build:
	go build -o bin/lootbot main.go
