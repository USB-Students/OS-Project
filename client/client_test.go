package client

import "testing"

var (
	host = "localhost"
	port = 8000
)

func TestMakeClient(t *testing.T) {
	MakeClient(host, port)
}
