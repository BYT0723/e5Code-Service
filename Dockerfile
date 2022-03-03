FROM golang:1.17-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH arm
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /app/user service/user/rpc/user.go


FROM alpine

WORKDIR /app
COPY --from=builder /build/zero/service/user/rpc/etc/user.yaml /app/etc/user.yaml
COPY --from=builder /app/user /app/user

EXPOSE 8001

CMD ["./user"]
