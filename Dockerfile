
FROM golang:1.18-alpine as builder
RUN apk add make binutils
COPY / /work
WORKDIR /work
RUN make rgw-exporter

FROM alpine:3.16
COPY --from=builder /work/bin/rgw-exporter /rgw-exporter
USER root
ENTRYPOINT ["/rgw-exporter"]

EXPOSE 9080
