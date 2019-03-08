#!/usr/bin/env bats

load helper

@test "BATS: BATS configured properly." {
	run bats --version
	[ "$output" = "Bats 0.4.0" ]
}

@test "BATS: bats-assert library configured properly." {
	run bats --help
	assert_contains "$output" "Usage:"
}

