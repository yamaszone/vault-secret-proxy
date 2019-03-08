FROM golang:1.12-alpine3.9 as builder

RUN apk add --update curl git

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/local/bin/dep
RUN chmod +x /usr/local/bin/dep

WORKDIR $GOPATH/src/vault-secret-proxy
COPY . .
RUN dep ensure
RUN go install ./cmd/vault-secrets-proxy-server/


# NOTE: Use alpine:3.9 below once its vulnerabilities are patched
FROM alpine:3.8

COPY kv-data.json.sample /etc/kv-data.json
COPY --from=builder /go/bin/vault-secrets-proxy-server /vault-secret-proxy

EXPOSE 8888
# TODO: Run as non-root
CMD ["/vault-secret-proxy", "--host", "0.0.0.0", "--port", "8888"]
