FROM golang:1.17-alpine as builder
RUN apk --no-cache add git
WORKDIR /go/src/build
COPY . .
RUN export CGO_ENABLED=0 \
    && mkdir -p dist \
    && go mod vendor
RUN go build -o dist/vault-backup .

FROM alpine:3.15
COPY --from=builder /go/src/build/dist/ /usr/local/bin/