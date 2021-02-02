default: run

build:
	go build -o dist/pier github.com/leviharrison/pier/cmd/pier

start:
	./pier

run: build start
