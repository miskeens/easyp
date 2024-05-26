# define build source
ARG ARTIFACT_SRC=local
ARG GO_TAG=1.22-alpine3.20

# ====================================================== CI BUILD ======================================================
# copy pre-built binaries if running in CI
FROM scratch as build-ci
COPY /easyp /

# ==================================================== SOURCE BUILD ====================================================
# full build if running locally
FROM golang:${GO_TAG} AS build-local

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOPATH=/var/go

RUN apk update --no-cache

WORKDIR /build

ADD go.mod go.sum ./
RUN --mount=type=cache,mode=0777,id=easyp-go-pkgs,target=/var/go/pkg/mod go mod download

COPY cmd internal ./

# TODO: add build verion override with -ldflags -X
RUN go build -o /easyp  ./cmd/easyp

# ==================================================== TARGET BUILD ====================================================
# define artifact source
FROM build-${ARTIFACT_SRC} as artifact-source

FROM alpine:3.20.0

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=artifact-source /easyp /easyp

RUN ln -s /easyp /usr/local/bin/easyp
