FROM golang:1.14.4-alpine3.12 as builder

ARG SSM_AGENT_VERSION=2.3.1205.0

RUN apk add --no-cache \
         'make~=4.3-r0' \
         'git~=2.26.2-r0' \
         'gcc~=9.3.0-r2' \
         'libc-dev~=0.7.2-r3' \
         'bash~=5.0.17-r0'

RUN wget -q https://github.com/aws/amazon-ssm-agent/archive/${SSM_AGENT_VERSION}.tar.gz && \
    mkdir -p /go/src/github.com && \
    tar xzf ${SSM_AGENT_VERSION}.tar.gz && \
    mv amazon-ssm-agent-${SSM_AGENT_VERSION} /go/src/github.com/amazon-ssm-agent && \
    echo ${SSM_AGENT_VERSION} > /go/src/github.com/amazon-ssm-agent/VERSION

WORKDIR /go/src/github.com/amazon-ssm-agent

RUN gofmt -w agent && make checkstyle || ./Tools/bin/goimports -w agent && \
    make build-linux

WORKDIR /go/src/server

COPY go.mod go.sum ./

RUN go mod download

RUN go get bitbucket.org/liamstask/goose/cmd/goose

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/server

FROM alpine:3.12 as prod

RUN apk add --no-cache \
      'aws-cli~=1.18.55-r0' \
      'sudo~=1.9.0-r0' \
      'mysql-client~=10.4.15-r0'

RUN adduser -D ssm-user && \
    echo "Set disable_coredump false" >> /etc/sudo.conf && \
    echo "ssm-user ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/ssm-agent-users && \
    mkdir -p /etc/amazon/ssm

COPY --from=builder /go/src/github.com/amazon-ssm-agent/bin/linux_amd64/ /usr/bin
COPY --from=builder /go/src/github.com/amazon-ssm-agent/bin/amazon-ssm-agent.json.template /etc/amazon/ssm/amazon-ssm-agent.json
COPY --from=builder /go/src/github.com/amazon-ssm-agent/bin/seelog_unix.xml /etc/amazon/ssm/seelog.xml
COPY --from=builder /go/bin/server /go/bin/server
COPY --from=builder /go/bin/goose /go/bin/goose

RUN mkdir -p /go/bin/db
COPY ./db/dbconf.yml /go/bin/db/dbconf.yml
COPY ./db/migrations /go/bin/db/migrations

COPY ./db/mysql/init /docker-entrypoint-initdb.d

EXPOSE 8080

COPY ./aws/docker-entrypoint.sh /

CMD ["sh", "/docker-entrypoint.sh"]