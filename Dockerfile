# syntax=docker/dockerfile:1
#
# goreleaser's dockers_v2 cross-compiles the binary before this ever runs and
# hands buildx a temporary context holding one <platform>/gwttr per target
# plus this file -- see rules/go-releases.md's "no go build inside the
# Dockerfile" for why there is nothing left to compile here. A local
# `docker build .` against the full repo works the same way as long as
# `linux/<arch>/gwttr` already exists at the repository root
# (`CGO_ENABLED=0 GOOS=linux GOARCH=$(go env GOARCH) go build -o
# "linux/$(go env GOARCH)/gwttr" ./cmd/gwttr`).
#
# distroless static, not alpine: the binary is CGO_ENABLED=0 and needs no
# libc, so there is nothing for even musl to provide, and the nonroot variant
# already runs as an unprivileged user (65532:65532) with no shell, package
# manager or other attack surface at all.
FROM gcr.io/distroless/static-debian12:nonroot@sha256:afa5c872c891853ca7fcf1f12c3edb23f7eeef36189728842dd51042ff57f7ab

ARG TARGETPLATFORM

COPY $TARGETPLATFORM/gwttr /usr/local/bin/gwttr

USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/gwttr"]
