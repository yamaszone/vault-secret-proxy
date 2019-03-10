// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	//"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	iam_auth "github.com/daveadams/onthelambda"
	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"vault-secret-proxy/restapi/operations"
	"vault-secret-proxy/utils"
)

//go:generate swagger generate server --target ../../vault-secret-proxy --name VaultSecretsProxy --spec ../swagger/swagger.yml

func configureFlags(api *operations.VaultSecretsProxyAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.VaultSecretsProxyAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TxtProducer = runtime.TextProducer()

	api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
		return operations.NewGetHealthOK()
	})
	api.GetSecretsHandler = operations.GetSecretsHandlerFunc(func(params operations.GetSecretsParams) middleware.Responder {
		kv_payload, err := utils.ReadJsonFile("/etc/vault/kv-data.json")
		if err != nil {
			api.Logger("ERROR: Failed to read key-value input data file.")
		}

		// The following is a cheap approach for local development
		if os.Getenv("VAULT_IS_STUB") == "yes" {
			for k, v := range kv_payload {
				url_parts := strings.Split(v.(string), "/")
				key_id := url_parts[len(url_parts)-1]
				// Use last part of the path as dummy secret
				kv_payload[k] = key_id
			}

		} else {
			client, err := iam_auth.VaultClient()
			if err != nil {
				api.Logger("ERROR: %s", err)
			}
			api.Logger("INFO: Successfully authenticated with Vault server.")

			for k, v := range kv_payload {
				path := v.(string)
				api.Logger("INFO: Fetching secrets for path: %s \n", path)
				response, err := client.Logical().Read(path)
				// Commented out for debugging
				//fmt.Println("%s %s", err, response)
				secret := ""
				if err != nil {
					api.Logger("ERROR: %s", err)
				}
				if response != nil {
					url_parts := strings.Split(path, "/")
					key_id := url_parts[len(url_parts)-1]
					secret = response.Data[key_id].(string)
				}
				kv_payload[k] = secret
			}
		}

		return operations.NewGetSecretsOK().WithPayload(kv_payload)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
