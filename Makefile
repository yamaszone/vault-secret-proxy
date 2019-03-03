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


# just added `gen` and `validate`
.PHONY: install gen validate
