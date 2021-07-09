FROM golang:latest
WORKDIR /go/src/app
COPY . .
ENV GOPROXY "https://goproxy.cn,direct"
RUN go install
ENTRYPOINT ["v2ray-create-config", "-u"]