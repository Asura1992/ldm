FROM golang:1.18.2

WORKDIR /opt

ADD . /opt

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o api-hello-srv  ./srvs/hello/main.go

CMD ["/opt/api-hello-srv"]
