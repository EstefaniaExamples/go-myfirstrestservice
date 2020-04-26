package main

import "testing"

func TestRun_FailsWithInvalidPort(t *testing.T) {
	err := run(1000000)

	if err == nil {
		t.Fatalf("expect error got nil")
	}
}