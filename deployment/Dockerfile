FROM golang:alpine
LABEL authors="Timur Kulakov"

WORKDIR /usr/src/app

COPY ../go.mod ../go.sum ./
RUN go mod download && go mod verify

COPY ../. .
RUN go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

RUN apk add --no-cache \
    protobuf \
    protobuf-dev \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN protoc -I ./api/proto/ --go_out=./pkg/proto --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/proto --go-grpc_opt=paths=source_relative api/proto/user.proto


RUN go build -v -o /usr/local/bin/app ./cmd/user-service/

CMD ["app"]