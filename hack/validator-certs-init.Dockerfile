FROM ubuntu:latest AS install

RUN set -e; \
  export DEBIAN_FRONTEND=noninteractive; \
  apt-get update; \
  apt-get install -y --no-install-recommends ca-certificates

FROM ubuntu:latest

COPY --from=install /usr/bin/openssl /usr/bin/openssl
COPY --from=install /usr/sbin/update-ca-certificates /usr/sbin/update-ca-certificates