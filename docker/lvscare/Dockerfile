FROM --platform=$BUILDPLATFORM alpine:3.16.2 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
ARG TARGETOS
ARG TARGETARCH
ARG VERSION
ARG PROJECT=sealos
ARG UNAME=labring

ADD https://github.com/${UNAME}/${PROJECT}/releases/download/v${VERSION}/sealos_${VERSION}_${TARGETOS}_${TARGETARCH}.tar.gz sealos.tar.gz
RUN tar -zxvf sealos.tar.gz lvscare && chmod +x lvscare

FROM alpine:3.16.2

RUN apk add --no-cache ipset iptables

WORKDIR /usr/bin/
COPY --from=builder /workspace/lvscare /usr/bin/
# nosemgrep: dockerfile.security.missing-user.missing-user
ENTRYPOINT ["/usr/bin/lvscare"]
# nosemgrep: dockerfile.security.missing-user.missing-user
CMD ["--help"]
