# styra-go-client

Inofficial Go client for Styra REST API primarily used for [crossplane-contrib/provider-styra](https://github.com/crossplane-contrib/provider-styra).

## Code Generation

The generated code is based on the [official OpenAPI v2 spec](https://docs.styra.com/reference/api/openapi) by Styra (with some manual adjustments). [go-swagger](https://github.com/go-swagger/go-swagger) is used to generate the client code.

To generate the Go code run

```bash
make clean generate
```
