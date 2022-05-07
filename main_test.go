package main

import (
	"testing"
)

func TestExpressDelivery(t *testing.T) {
	var a = []int{1, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0, 1}
	t.Log(minSender(a))
}
