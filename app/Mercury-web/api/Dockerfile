FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.io,direct

WORKDIR /build/zero
COPY . .
RUN sh -c "[ -f go.mod ]" || exit
COPY app/Mercury-web/api/etc /app/etc
RUN go build -ldflags="-s -w" -o /Mercury-web-api app/Mercury-web/api/Mercury.go


FROM alpine

RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories
RUN cat /etc/apk/repositories
RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /Mercury-web-api /app/Mercury-web-api
COPY --from=builder /app/etc /app/etc

CMD ["./Mercury-web-api", "-f", "etc/Mercury-api.yaml"]