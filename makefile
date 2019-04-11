build: build-server build-agent

build-server:
	go get -v ./server/...
	go build -o ./server/.bin/server ./server/main.go

build-agent:
	go get -v ./agent/...
	go build -o ./agent/.bin/agent ./agent/main.go