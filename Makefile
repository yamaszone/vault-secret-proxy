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
		--exclude-main \
		--name=vault-secrets-proxy

run: install
	cp kv-data.json.sample /tmp/kv-data.json
	~/go/bin/vault-secrets-proxy-server

# just added `gen` and `validate`
.PHONY: install gen validate
