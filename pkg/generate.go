package pkg

import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
)

//go:generate go run -tags generate github.com/go-swagger/go-swagger/cmd/swagger generate client -f ../doc/styra-swagger.json -c client -m models --default-scheme=https -q
