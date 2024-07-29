package user_service

import _ "github.com/open-streaming-solutions/shared"

// gen.go used for more convenient and flexible generation

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0 generate
//go:generate protoc -I $GOPATH/pkg/mod/github.com/open-streaming-solutions/shared@v0.1.2/user-service/ --go_out=./pkg/proto --go_opt=paths=source_relative --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative $GOPATH/pkg/mod/github.com/open-streaming-solutions/shared@v0.1.2/user-service/user-service.proto
