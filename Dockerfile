FROM golang:latest

ENV GO111MODULE=on

WORKDIR /go/src/github.com/silalahi/nantoari

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /go/bin/nantoari  ./app/main.go

ENTRYPOINT /go/bin/nantoari

EXPOSE 8080