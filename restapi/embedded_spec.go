// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Vault Secrets Proxy Sidecar",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "summary": "server health check",
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/secrets": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "get secrets configured in this sidecar",
        "operationId": "getSecrets",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "description": "list of secrets",
              "type": "object"
            }
          },
          "404": {
            "description": "Secrets not found"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Vault Secrets Proxy Sidecar",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "text/plain"
        ],
        "summary": "server health check",
        "operationId": "getHealth",
        "responses": {
          "200": {
            "description": "OK"
          }
        }
      }
    },
    "/secrets": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "get secrets configured in this sidecar",
        "operationId": "getSecrets",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "description": "list of secrets",
              "type": "object"
            }
          },
          "404": {
            "description": "Secrets not found"
          }
        }
      }
    }
  }
}`))
}
