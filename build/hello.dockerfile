FROM golang:1.12-alpine as builder

RUN apk add --no-cache ca-certificates curl git
ENV GO111MODULE=on

WORKDIR /app
COPY . .

WORKDIR /app/cmd/hello

RUN CGO_ENABLED=0 GOOS=`go env GOHOSTOS` GOARCH=`go env GOHOSTARCH` go build -o hello


FROM alpine:3.18.3

WORKDIR /

RUN apk add --no-cache ca-certificates
COPY --from=builder /app/cmd/hello/hello .
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip  /usr/local/go/lib/time/zoneinfo.zip

ENTRYPOINT ["/hello"]
