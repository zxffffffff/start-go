all: deps build

.PHONY: deps
deps:
	go get -d -v github.com/dustin/go-broadcast/...

.PHONY: build
build: deps
	go build -o realtime-chat main.go server/srv.go server/rooms.go server/template.go
