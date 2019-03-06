# Vault Secret Proxy
PoC Vault secret proxy to work as a stub server

## Prerequisites
- [Go v1.12](https://golang.org/dl/)
- [dep v0.5.0](https://github.com/golang/dep/releases/tag/v0.5.0)
- [go-swagger v0.18.0](https://github.com/go-swagger/go-swagger/releases/tag/v0.18.0)

## Quickstart
- Checkout the repository
- From project root
  - `make dep` # Installs all necessary dependencies
  - `make install` to install binary as `~/go/bin/vault-secret-proxy`
  - `~/go/bin/vault-secrets-proxy-server --help`

- Launch server
  - `$ make run`

- Play with the server
  - `curl -sS http://127.0.0.1:9999/v1/secrets | jq .`
  ```
  {
    "API_TOKEN": "token",
    "DB_PASSWORD": "password"
  }
  ```
  - Using [httpie](https://github.com/jakubroztocil/httpie)
    - `$ http get http://127.0.0.1:9999/v1/secrets`
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
    - `$ http get http://127.0.0.1:9999/v1/healthz`
    ```
    HTTP/1.1 200 OK
    Content-Length: 0
    Date: Sun, 03 Mar 2019 03:08:54 GMT
    ```
