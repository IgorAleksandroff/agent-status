package agent_status

//go:generate protoc -I=. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/service.proto
//go:generate protoc -I=. --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import proto/auto_assigment.proto
