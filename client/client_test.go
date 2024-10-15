package client

import "testing"

var (
	host = "localhost"
	port = 2000
)

func TestMakeClient(t *testing.T) {
	MakeClient(host, port)
}
