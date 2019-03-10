# Vault Secret Proxy Sidecar
Vault secret proxy [sidecar](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/45406.pdf) using [AWS IAM Auth](https://www.vaultproject.io/docs/auth/aws.html#iam-auth-method).

## Usage
Vault secret proxy sidecar can be used by any service (__Primary App__) running within a Kubernetes Pod in AWS. It can be used as a stub server during local development. See the following links for usage instructions:
- [Local Development](https://github.com/yamaszone/vault-secret-proxy/blob/master/docs/vault-secret-proxy-sidecar-usage.md#local-development)
- [Kubernetes Cluster in AWS](https://github.com/yamaszone/vault-secret-proxy/blob/master/docs/vault-secret-proxy-sidecar-usage.md#local-development)

## Development
#### Prerequisites
- [Go v1.12](https://golang.org/dl/)
- [dep v0.5.0](https://github.com/golang/dep/releases/tag/v0.5.0)
- [go-swagger v0.18.0](https://github.com/go-swagger/go-swagger/releases/tag/v0.18.0)

#### Quickstart
- Checkout the repository
- From project root
  - `make dep` # Installs all necessary dependencies
  - `make install` to install binary as `~/go/bin/vault-secret-proxy`
  - `~/go/bin/vault-secrets-proxy-server --help`

- Launch server
  - `$ make run`

- Play with the server
  - `curl -sS http://localhost:8888/v1/secrets | jq .`
  ```
  {
    "API_TOKEN": "token",
    "DB_PASSWORD": "password"
  }
  ```
  - Using [httpie](https://github.com/jakubroztocil/httpie)
    - `$ http get http://localhost:8888/v1/secrets`
    ```
    HTTP/1.1 200 OK
    Content-Length: 30
    Content-Type: application/json
    Date: Sun, 03 Mar 2019 03:08:12 GMT

    {
      "API_TOKEN": "token",
      "DB_PASSWORD": "password"
    }
    ```
    - `$ http get http://localhost:8888/v1/healthz`
    ```
    HTTP/1.1 200 OK
    Content-Length: 0
    Date: Sun, 03 Mar 2019 03:08:54 GMT
    ```
#### Credits
- [github.com/daveadams/onthelambda](github.com/daveadams/onthelambda)
