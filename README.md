# healthchecker

Go Main program for Docker Healtcheck.

## Description

This program will call an URL, and return 0 if call succeded (2xx status code), or 1 otherwise.

## Options

- debug : will print to stdout debug infos
- endpoint : force a specific local end point. Must start with `/`. Default to `/heath`
- ip : IP address to call. default to 127.0.0.1
- port : port to call, default to 80
- timeout : timeout for the http call, default to 20s.

Exemples :

```shell
 healthcheck -debug -ip=192.168.0.1 -port=5010 -endpoint=/
 healthcheck -timeout=50s
```

## Dockerfile

May be used in DockerFile as an HEALTH checker.

```Dockerfile
# Create a minimal whoami docker image
# with multi-stage build
# require a "recent" docker version.
FROM golang:1.9.2-stretch as compiler

# Copy file from current dir
WORKDIR /go/src/github.com/jfphilippe/whoami
COPY *.go /go/src/github.com/jfphilippe/whoami/
# Compile
RUN CGO_ENABLED=0 go install --tags netgo -a -ldflags '-s -w' .

# add healthchecker
RUN CGO_ENABLED=0 go get --tags netgo -a -ldflags '-s -w' github.com/jfphilippe/healthchecker

FROM scratch
COPY --from=compiler /go/bin/whoami /whoami
COPY --from=compiler /go/bin/healthchecker /healthchecker
EXPOSE 3000

# Healthcheck in exec mode, because scratch has no shell
HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
        CMD [ "/healthchecker", "-port=3000" ]

ENTRYPOINT ["/whoami"]
```