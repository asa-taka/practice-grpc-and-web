.PHONY: \
	grpc-all grpc-go \
	swagger \
	rest-all rest-dart rest-go rest-ts

all: grpc-all swagger rest-all

# gRPC Stubs from Protobuf Definition

GOPATH := ${GOPATH}

PROTOC_OPTS := \
	-I./proto \
	-I$(GOPATH)/src \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
	-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis

PROTO_SRC := todo.proto
GRPC_DST := grpc-stubs

grpc-all: grpc-dart grpc-go grpc-ts

grpc-dart:
	mkdir -p $(GRPC_DST)/dart
	protoc $(PROTOC_OPTS) \
		--dart_out=grpc:$(GRPC_DST)/dart \
		$(PROTO_SRC)

grpc-go:
	mkdir -p $(GRPC_DST)/go
	protoc $(PROTOC_OPTS) \
		--go_out=plugins=grpc:$(GRPC_DST)/go \
		$(PROTO_SRC)

grpc-ts:
	mkdir -p $(GRPC_DST)/ts
	protoc $(PROTOC_OPTS) \
    --js_out="import_style=commonjs,binary:$(GRPC_DST)/ts" \
    --ts_out="$(GRPC_DST)/ts" \
		$(PROTO_SRC)

# Swagger Definition

swagger:
	protoc $(PROTOC_OPTS) \
		--swagger_out=logtostderr=true:swagger \
		$(PROTO_SRC)

# REST Stubs fron Swagger Definition

rest-all: swagger rest-dart rest-go rest-ts

SWAGGER_SRC := swagger/todo.swagger.json
REST_DST := rest-stubs

define generate_rest_stub
	swagger-codegen generate -l $(1) -i $(SWAGGER_SRC) -o $(REST_DST)/$(1)
endef

rest-dart:
	$(call generate_rest_stub,dart)

rest-go:
	$(call generate_rest_stub,go)

rest-ts:
	$(call generate_rest_stub,nodejs-server)
	$(call generate_rest_stub,typescript-fetch)

# Miscs

clean:
	rm -r $(GRPC_DST)/*
	rm -r swagger/*
	rm -r $(REST_DST)/*