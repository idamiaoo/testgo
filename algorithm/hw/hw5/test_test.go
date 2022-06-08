package main

import (
	"testing"

	"google.golang.org/grpc"
)

func TestAtoi(t *testing.T) {
	t.Log(atoi(" -123b"))
	grpc.WithDisableRetry()
}
