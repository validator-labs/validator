# Build the manager binary
FROM --platform=$TARGETPLATFORM golang:alpine3.19 AS builder
ARG TARGETOS
ARG TARGETARCH

RUN apk add --no-cache curl

WORKDIR /workspace

# Get Helm
RUN curl -s https://get.helm.sh/helm-v3.10.1-linux-amd64.tar.gz | tar -xzf - && \
    mv linux-amd64/helm . && rm -rf linux-amd64

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/main.go cmd/main.go
COPY api/ api/
COPY internal/ internal/
COPY pkg/ pkg/

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o manager cmd/main.go

RUN chmod 777 /etc /etc/ssl && chmod -R 777 /etc/ssl/certs && \
    mkdir /.cache /charts && chmod -R 777 /.cache /charts

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM --platform=$TARGETPLATFORM gcr.io/distroless/static:nonroot AS production
WORKDIR /

COPY --from=builder /workspace/manager .
COPY --from=builder /workspace/helm .
COPY --from=builder --chown=65532:65532 /etc /etc
COPY --from=builder --chown=65532:65532 /.cache /.cache
COPY --from=builder --chown=65532:65532 /charts /charts

USER 65532:65532

ENTRYPOINT ["/manager"]
