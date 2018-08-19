# Practice to generate gRPC and REST Stubs from Protobuf

This project aims to practice to generate Stubs from a Protobuf definition through followign flows.

- `.proto` → gRPC Stubs
- `.proto` → Swagger definition → REST Stub

## Requirements

- `protoc-gen-{go,dart,ts,swagger}`
- `swagger-codegen`

## Generate gRPC and REST Stubs

Put `make` then you'll get

```
$ tree -L 2
.
├── Makefile
├── README.md
├── grpc-stubs
│   ├── dart
│   ├── go
│   └── ts
├── proto
│   └── todo.proto
├── rest-stubs
│   ├── dart
│   ├── go
│   ├── nodejs-server
│   └── typescript-fetch
└── swagger
    └── todo.swagger.json
```