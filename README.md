```shell
# install generator
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# generate stubs
for c in server client types; do
  oapi-codegen -config api/$c.cfg.yaml api/swagger.yaml
done
```