# Builder
FROM golang:1.14 as builder

WORKDIR /src
COPY . ./
RUN make build

# Builder
FROM alpine

COPY --from=builder /src/app /
ENTRYPOINT ["/app"]
