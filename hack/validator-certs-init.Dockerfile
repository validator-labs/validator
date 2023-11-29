FROM --platform=$TARGETPLATFORM ubuntu:latest AS install

RUN set -e; \
  export DEBIAN_FRONTEND=noninteractive; \
  apt-get update; \
  apt-get install -y --no-install-recommends ca-certificates

# finalize, keeping only required artifacts
FROM --platform=$TARGETPLATFORM ubuntu:latest

COPY --from=install /usr/bin/openssl /usr/bin/openssl
COPY --from=install /usr/sbin/update-ca-certificates /usr/sbin/update-ca-certificates
COPY --from=install /usr/share/ca-certificates /usr/share/ca-certificates
COPY --from=install /etc/ca-certificates.conf /etc/ca-certificates.conf

USER 65532:65532