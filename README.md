# user-service

> **Note:** This project utilizes code generation for certain components.
> If you wish to contribute, you'll need to ensure these code generation dependencies are installed and run properly. 
> Don't forget to set these environment variables before running the application.

## Environment Variables

To configure the application, set the following environment variables:

| Variable                    | Description              | Example Value    |
|-----------------------------|--------------------------|------------------|
| `USER_SERVICE_DB_USERNAME`  | Database username        | `db_user`        |
| `USER_SERVICE_DB_PASSWORD`  | Database password        | `db_password`    |
| `USER_SERVICE_DB_NAME`      | Database name            | `user_service_db`|
| `USER_SERVICE_DB_HOST`      | Database host            | `localhost`      |
| `USER_SERVICE_DB_PORT`      | Database port            | `5432`           |
| `USER_SERVICE_PORT`         | Application port         | `8080`           |

## Local build

**For building docker image:**

```sh
docker build -t app -f ./deployment/Dockerfile . 
```

**For simple building app by Go:**

```sh
go run ./cmd/user-service/main.go
```

**For simple building app by exec:**

```sh
exec ./<user-service-binary>
```

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