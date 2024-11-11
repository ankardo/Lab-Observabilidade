FROM golang:1.23.3 AS golang

ARG SERVICE
ARG SERVICE_PORT

ENV SERVICE=${SERVICE}
ENV SERVICE_PORT=${SERVICE_PORT}

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

FROM golang AS builder
WORKDIR /build
COPY . .
RUN go build -o main ./${SERVICE}/cmd/app/main.go && \
  chmod +x main

FROM alpine:3.19 AS upx
RUN apk add --no-cache upx=4.2.1-r0
COPY --from=builder /build/main /upx/main
RUN upx --best --lzma /upx/main -o /upx/main_compressed

FROM scratch AS main
WORKDIR /app
COPY --from=upx /upx/main_compressed /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT [ "./main" ]
EXPOSE ${SERVICE_PORT}
