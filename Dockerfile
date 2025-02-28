
FROM golang:1.24 as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

COPY / /work
WORKDIR /work

RUN make rgw-exporter

FROM scratch
COPY --from=builder /work/bin/rgw-exporter /rgw-exporter
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

USER 999
ENTRYPOINT ["/rgw-exporter"]

EXPOSE 9080
