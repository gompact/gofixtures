FROM golang

RUN apt-get update && apt-get install -y netcat

WORKDIR /go/src/github.com/schehata/gofixtures

COPY go.mod go.sum ./

RUN go mod download

COPY . .

CMD until nc -z $GOFIXTURES_TEST_DB_HOST $GOFIXTURES_TEST_DB_PORT; do sleep 1; done &&\
    go test -v ./...

