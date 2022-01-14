package pkg

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger" //nolint:typecheck
)

//go:generate go run -tags generate github.com/go-swagger/go-swagger/cmd/swagger generate client -f ../api/styra-openapi-v2.json -c client -m models --default-scheme=https -q
