default: run

build:
	go build -o dist/pier github.com/leviharrison/pier/cmd/pier

start:
	dist/pier

run: build start
