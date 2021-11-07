FROM golang:latest

WORKDIR $GOPATH/src/test
COPY . $GOPATH/src/test
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
RUN  go mod tidy&&go build -o webapp
  
EXPOSE 8081
HEALTHCHECK --interval=30s --timeout=3s \
    CMD curl -f http://localhost:8081/ || exit 1
ENTRYPOINT ["./webapp"]