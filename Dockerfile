FROM golang:1.22 AS builder

LABEL stage=gobuilder

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

ARG AppDir

WORKDIR /build

ADD ${AppDir}/go.mod .
ADD ${AppDir}/go.sum .
ADD core .

RUN go mod tidy && go mod download

COPY service/chat/api .
COPY ${AppDir}/etc /app/etc

RUN cd ${AppDir} && go build -ldflags="-s -w" -o /app/app

FROM scratch

ARG AppDir
LABEL iotdreamcatcher.net.service=${AppDir}
LABEL iotdreamcatcher.net.slslog=true

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/etc /app/etc

CMD ["./app", "-f", "etc/config.yaml"]
