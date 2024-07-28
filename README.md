# user-service

This project utilizes code generation for certain components. If you wish to contribute, you'll need to ensure these code generation dependencies are installed and run properly.

## Contributors Guide

### Step 1: Install Protobuf Dependencies

**For Alpine Linux:**

```sh
apk add protobuf protobuf-dev
```

**Windows:**

If you haven’t installed the compiler, [download the package](https://protobuf.dev/downloads/) and follow the instructions in the README.

### Step 2: Install Proto Dependencies for Go

**For Alpine Linux or Windows:**

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Step 3: Run Code Generation

**Using `go generate`:**

```sh
go generate ./...
```

**Or manually run:**

```sh
go run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0 generate
# little long, fix in WIP
protoc -I $GOPATH/pkg/mod/github.com/!open-!streaming-!solutions/shared/user-service@v0.0.0-20240726232231-b7b2469732b3/ --go_out=./pkg/proto --go_opt=paths=source_relative --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative $GOPATH/pkg/mod/github.com/!open-!streaming-!solutions/shared/user-service@v0.0.0-20240726232231-b7b2469732b3/user-service.proto
```

---

This README provides clear, step-by-step instructions for setting up the necessary dependencies and running code generation.