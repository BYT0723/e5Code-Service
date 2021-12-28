FROM arm32v7/golang:1.17-alpine

WORKDIR /go/src/

RUN apk add git

RUN git clone https://github.com/BYT0723/e5Code-Service.git

ENV GOPROXY="https://goproxy.io"

WORKDIR /go/src/e5Code-Service/service/user/rpc/

RUN go build -o userRpc user.go

EXPOSE 8001

ENTRYPOINT [ "./userRpc","-f","etc/user.yaml"]
