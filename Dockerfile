FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod .
COPY cmd/ cmd/
RUN go build -o /out/hash-checker ./cmd/hash-checker && \
    go build -o /out/b64 ./cmd/b64 && \
    go build -o /out/cidr-calc ./cmd/cidr-calc && \
    go build -o /out/dns-lookup ./cmd/dns-lookup && \
    go build -o /out/jwt-decode ./cmd/jwt-decode && \
    go build -o /out/ioc-extract ./cmd/ioc-extract && \
    go build -o /out/log-parser ./cmd/log-parser && \
    go build -o /out/port-scanner ./cmd/port-scanner && \
    go build -o /out/tls-check ./cmd/tls-check && \
    go build -o /out/pwcheck ./cmd/pwcheck && \
    go build -o /out/http-inspect ./cmd/http-inspect && \
    go build -o /out/report-gen ./cmd/report-gen

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/ /usr/local/bin/
