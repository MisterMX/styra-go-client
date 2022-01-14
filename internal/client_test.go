package internal

import (
	"testing"

	openAPIClient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mistermx/styra-go-client/pkg/client"
)

// Test if the client code compiles.
func TestClient(t *testing.T) {
	transport := openAPIClient.New("example.com", "/", client.DefaultSchemes)
	client.New(transport, strfmt.Default)
}
