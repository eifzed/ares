
generate:
	protoc -I pb --go_out=./pb --go_opt=module=github.com/eifzed/ares/pb --go-grpc_out=./pb --go-grpc_opt=module=github.com/eifzed/ares/pb pb/antre.proto

generate-dart:
	protoc -I pb --dart_out=./pb pb/ares.proto google/protobuf/timestamp.proto

	