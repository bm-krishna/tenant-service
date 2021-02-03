
protoc --go_out=pkg --go_opt=paths=source_relative api/tenant.proto   
protoc --go-grpc_out=pkg --go-grpc_opt=paths=source_relative api/tenant.proto