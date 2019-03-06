package utils

import (
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	kvdata, err := ReadJsonFile("/etc/kv-data.json")
	if err != nil {
		t.Error("Error reading KV file.")
	}

	is_invalid_api_token := kvdata["API_TOKEN"] != "vault-mount/app/token"
	is_invalid_db_pass := kvdata["DB_PASSWORD"] != "vault-mount/app/db/password"
	if is_invalid_api_token || is_invalid_db_pass {
		t.Error("Error parsing KV file.")
	}

}
