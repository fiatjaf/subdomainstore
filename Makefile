subdomainstore: $(shell find . -name "*.go") bindata.go
	go build -ldflags="-s -w" -o ./subdomainstore

bindata.go: static/index.html
	go-bindata -o bindata.go static/...
