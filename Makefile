dist:
    @mkdir -p ./bin
    @rm -f ./bin/*
    GOFIXTURES=darwin GOARCH=amd64 go build -o ./bin/gofixtures ./cmd/gofixtures
    GOFIXTURES=linux GOARCH=amd64 go build -o ./bin/gofixtures ./cmd/gofixtures
    GOFIXTURES=linux GOARCH=386 go build -o ./bin/gofixtures ./cmd/gofixtures
