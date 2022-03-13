ARG GOLANG_VERSION=1.17

# Building custom health checker
FROM golang:$GOLANG_VERSION as health-build-env

# Copying source
WORKDIR /go/src/app
COPY ./healthcheck /go/src/app

# Installing dependencies
RUN go get -d -v ./...

# Compiling
RUN go build -o /go/bin/healthchecker

# Building bouncer
FROM golang:$GOLANG_VERSION as build-env

# Copying source
WORKDIR /go/src/app
COPY . /go/src/app

# Installing dependencies
RUN go get -d -v ./...

# Compiling
RUN go build -o /go/bin/cs-mikrotik-bouncer

FROM gcr.io/distroless/base:nonroot
COPY --from=health-build-env --chown=nonroot:nonroot /go/bin/healthchecker /
COPY --from=build-env --chown=nonroot:nonroot /go/bin/cs-mikrotik-bouncer /

# Run as a non root user.
USER nonroot

# Using custom health checker
# HEALTHCHECK --interval=10s --timeout=5s --retries=2\
#   CMD ["/healthchecker"]

# Run app
CMD ["/cs-mikrotik-bouncer"]