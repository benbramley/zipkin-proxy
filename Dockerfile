FROM golang:1.9-alpine as builder
WORKDIR /go/src/github.com/namely/zipkin-proxy
COPY . .
RUN go build -o zipkin-proxy github.com/namely/zipkin-proxy/cmd/zipkinproxy

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/namely/zipkin-proxy/zipkin-proxy .
ENV LISTEN_ON 0.0.0.0:9411
CMD ["./zipkin-proxy", "server"]
