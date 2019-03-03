# Vault Secret Proxy
PoC Vault secret proxy to work as a stub server

## Prerequisites
- [dep v0.5.0](https://github.com/golang/dep)
- [go-swagger v0.18.0](https://github.com/go-swagger/go-swagger)

## Quickstart
- Checkout the repository
- From project root
  - `make dep` # Installs all necessary dependencies
  - `make install` to install binary as `~/go/bin/vault-secret-proxy`
  - `~/go/bin/vault-secrets-proxy-server --help`
  ```
  Usage: vault-secrets-proxy-server [--port PORT]

  Options:
    --port PORT, -p PORT   port to listen to [default: 9999]
    --help, -h             display this help and exit
  ```
  - `$ ~/go/bin/vault-secrets-proxy-server`
  - `curl -sS http://127.0.0.1:9999/secrets | jq .`
  ```
  {
    "key1": "val1",
    "key2": "val2"
  }
  ```
  - Using [httpie](https://github.com/jakubroztocil/httpie)
    - `$ http get http://127.0.0.1:9999/secrets`
    ```
    HTTP/1.1 200 OK
    Content-Length: 30
    Content-Type: application/json
    Date: Sun, 03 Mar 2019 03:08:12 GMT

    {
        "key1": "val1",
        "key2": "val2"
    }
    ```
    - `$ http get http://127.0.0.1:9999/healthz`
    ```
    HTTP/1.1 200 OK
    Content-Length: 0
    Date: Sun, 03 Mar 2019 03:08:54 GMT
    ```
