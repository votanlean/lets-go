```bash
go run main.go

curl localhost:8090/hello
curl -H "Authorization: Bearer token123" -H "X-Api-Key: apikey456" localhost:8090/headers

```