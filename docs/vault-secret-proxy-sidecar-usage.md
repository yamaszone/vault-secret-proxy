# Usage
Vault secret proxy [sidecar](https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/45406.pdf) uses [AWS IAM Auth](https://www.vaultproject.io/docs/auth/aws.html#iam-auth-method) to authenticate and fetch/renew secrets from Vault. During local development from a non-AWS environment, use a Vault secret proxy stub server preconfigured with dummy/pre-prod secrets (see [instructions](https://github.com/yamaszone/vault-secret-proxy/new/master#local-development) below).

## Local Development
- Update [kv-data.json.sample](https://github.com/yamaszone/vault-secret-proxy/blob/master/kv-data.json.sample) using the proper ENV variables (e.g. `API_TOKEN`) used by the Primary App (i.e. the secret consumer) as the Key and corresponding secret path (e.g. `vault-mount/app/token`) to Vault as the Value.
- Launch Vault secret proxy stub server using [`./proxy`](https://github.com/yamaszone/vault-secret-proxy/blob/master/proxy)
  - `./proxy up stub`
  - `curl -sS http://localhost:8888/v1/secrets`
  ```
  {
    "API_TOKEN": "token"
  }
  ```
    - Primary App can now get dummy/pre-prod secrets as above using `http://localhost:8888/v1/secrets` as the Vault URL.

## Production
- Prepare a Kubernetes manifest similar to [`vault-proxy-sidecar.yaml`](https://github.com/yamaszone/vault-secret-proxy/blob/master/deployments/vault-proxy-sidecar.yaml) with the following changes:
  - Update the fields (e.g. ENV values, KV fields, etc.) in `data` under `ConfigMap` object using appropriate values in your context/environment
  - Use a proper `tag` instead of `latest` for the `yamaszone/vault-proxy:latest` sidecar container
- Deploy your stack
  - `kubectl apply -f ./deployments/vault-proxy-sidecar.yaml`
  - Primary App can now get dummy/pre-prod secrets as above using `http://localhost:8888/v1/secrets` as the Vault URL. 
