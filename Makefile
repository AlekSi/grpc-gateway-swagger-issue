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

	# --with-flatten=remove-unused
	# swagger flatten --with-flatten=expand api/*.swagger.json > api/flattened.swagger.json
	swagger mixin api/api.json api/*.swagger.json > api/swagger.json
	swagger validate api/swagger.json

	swagger generate client --spec=api/swagger.json --target=api/swagger

	go install -v ./...

clean:
	rm -f api/*.pb.go api/*.pb.gw.go api/*.swagger.json api/swagger.json
	rm -fr api/swagger
	mkdir api/swagger
