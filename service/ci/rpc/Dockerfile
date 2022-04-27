FROM golang:1.17.7-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY service/ci/rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/ci service/ci/rpc/ci.go


FROM alpine

USER root:root

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/ci /app/ci
COPY --from=builder /app/etc /app/etc

EXPOSE 8003

CMD ["./ci", "-f", "etc/ci.yaml"]
