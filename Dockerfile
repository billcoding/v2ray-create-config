FROM golang:latest as builder
WORKDIR /go/src/app
COPY . .
ENV GOPROXY "https://goproxy.cn,direct"
RUN go install
FROM ubuntu:latest
COPY --from=builder /go/bin/v2ray-create-config .
ENTRYPOINT ["./v2ray-create-config"]