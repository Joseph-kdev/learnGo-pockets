//go:generate bash -c "protoc -I=api/proto/ --go_out=api/ --go_opt=paths=source_relative --go-grpc_out=api/ --go-grpc_opt=paths=source_relative api/proto/*.proto"
package api