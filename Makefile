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
	sudo cp kv-data.json.sample /etc/kv-data.json
	~/go/bin/vault-secrets-proxy-server --host=0.0.0.0 --port=9999

all: dep run

.PHONY: dep
