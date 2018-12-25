all: clean
	go install -v ./vendor/github.com/golang/protobuf/protoc-gen-go
	go install -v ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go install -v ./vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install -v ./vendor/github.com/go-swagger/go-swagger/cmd/swagger

	protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --go_out=plugins=grpc:.

	protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --grpc-gateway_out=logtostderr=true:.

	protoc --proto_path=. \
			--proto_path=vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			api/*.proto --swagger_out=logtostderr=true:.

	find api -name '*.swagger.json' -print -exec swagger flatten --with-flatten=expand --with-flatten=remove-unused '{}' --output='{}' \;

	swagger mixin api/swagger/info.json api/*.swagger.json > api/swagger/swagger.json
	swagger validate api/swagger/swagger.json
	rm -f api/*.swagger.json

	swagger generate client --spec=api/swagger/swagger.json --name=issue --target=api/swagger

	go install -v ./...

clean:
	rm -f api/*.pb.go api/*.pb.gw.go api/*.swagger.json
	rm -fr api/swagger/swagger.json api/swagger/client api/swagger/models
