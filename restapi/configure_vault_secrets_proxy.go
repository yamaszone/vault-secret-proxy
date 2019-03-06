// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	//"fmt"
	"log"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	//iam_auth "github.com/daveadams/onthelambda"
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
		kv_payload, err := utils.ReadJsonFile("/etc/kv-data.json")
		if err != nil {
			api.Logger("Error reading key-value data file.")
		}
		/*
			client, err := iam_auth.VaultClient()
			if err != nil {
				log.Fatalf("ERROR: %s", err)
			}
		*/
		for k, v := range kv_payload {
			/*
				resp, err := client.Logical().Read(v)
				if err != nil {
					api.Logger("ERROR: %s", err)
				}
			*/
			url_parts := strings.Split(v.(string), "/")
			key_id := url_parts[len(url_parts)-1]
			//kv_payload[k] = resp.Data[key_id].(string)
			// below is a placeholder
			kv_payload[k] = key_id
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
