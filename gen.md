```
进入 Mercury-web/api
goctl api go -api Mercury-api.api -dir ./ -style gozero 

进入 Mercury-web/rpc
goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./ --style=goZero

```

