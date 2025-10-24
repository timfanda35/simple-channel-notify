# Builder
FROM golang:1.25 AS builder

WORKDIR /src
COPY . ./
RUN make build

# Builder
FROM alpine

COPY --from=builder /src/app /

ENTRYPOINT ["/app"]
