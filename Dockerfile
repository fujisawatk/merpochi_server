FROM golang:1.14.2 as builder

WORKDIR /go/src/server

RUN groupadd -g 10001 merpochi \
    && useradd -u 10001 -g merpochi merpochi

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server

FROM alpine:3.12

COPY --from=builder /go/bin/server /go/bin/server
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 8080

USER merpochi

ENTRYPOINT ["/go/bin/server"]