# This Makefile used only during local development on macOS. Prefer Docker.
dep:
	dep ensure -v

install:
	go install -v ./cmd/vault-secrets-proxy-server/

validate:
	swagger validate ./swagger/swagger.yml

gen: validate
	swagger generate server \
		--target=. \
		--spec=./swagger/swagger.yml \
		--name=vault-secrets-proxy

run: install
	sudo mkdir -p /etc/vault
	sudo cp kv-data.json.sample /etc/vault/kv-data.json
	~/go/bin/vault-secrets-proxy-server --host=0.0.0.0 --port=9999

test:
	go test ./... -coverprofile=cover.out

all: dep run

.PHONY: dep
