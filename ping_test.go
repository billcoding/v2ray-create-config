package main

import (
	"testing"
)

func TestTelnet(t *testing.T) {
	t.Log(telnet("64.64.248.156", 20001))
}
