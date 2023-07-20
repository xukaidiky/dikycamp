export tag=v1.0
root:
	export ROOT=github.com/dikycamp/golang

build:
	echo "building httpserver binary"
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .

release: build
	echo "building httpserver container"
	docker build -t xkdiky/dikycamp/httpserver:${tag} .

push: release
	echo "pushing xkdiky/dikycamp/httpserver"
	docker push xkdiky/dikycamp/httpserver:v1.0
