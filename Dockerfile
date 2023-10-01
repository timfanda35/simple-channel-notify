# Builder
FROM golang:1.21 as builder

WORKDIR /src
COPY . ./
RUN make build

# Builder
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app /

ENTRYPOINT ["/app"]
