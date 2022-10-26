# microService_demo
a micro-service demo based go-microv4.dev

### 开发
[download protoc](https://github.com/protocolbuffers/protobuf/releases)


download go-micro plugin for protoc, see [github generator](https://github.com/go-micro/generator) for detail:
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
```

generate pb file:
```
protoc --go_out=. --micro_out=. proto/*.proto
```
