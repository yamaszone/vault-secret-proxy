#!/usr/bin/env bats

load helper

# NOTE: Rewrite this using https://github.com/kubernetes/client-go with better
# control loops

@test "VAULT_PROXY: Primary app can retrieve secrets from Vault proxy stub." {
	local manifest=tests/bats/fixtures/vault-proxy-stub.yaml

	run kubectl apply -f $manifest
	[ "$status" -eq 0 ]

	sleep 15
	pod_id=$(kubectl get pods | grep "^vault-canary-stub" | grep "Running" | awk '{print $1}')
	# wait is an experimental feature. Commented out for now.
	#kubectl wait --for=condition=Ready pod/${pod_id}


	run kubectl exec $pod_id -c primary-app -- sh -c "curl -sS http://localhost:8888/v1/secrets"
	[ "$status" -eq 0 ]
	assert_contains "$output" "token"
	assert_contains "$output" "password"

	kubectl delete -f $manifest
}
