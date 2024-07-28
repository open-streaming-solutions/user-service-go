# Stage 1: Build the application
FROM golang:alpine AS builder
LABEL authors="Timur Kulakov"

WORKDIR /usr/src/app

COPY ../go.mod ../go.sum ./
RUN go mod download && go mod verify

COPY ../. .

RUN apk add --no-cache \
    curl \
    protobuf \
    protobuf-dev \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Download Atlas binary
RUN curl -sSfL https://release.ariga.io/atlas/atlas-community-linux-amd64-latest -o /usr/local/bin/atlas \
    && chmod +x /usr/local/bin/atlas

RUN go generate ./...
RUN go mod tidy

RUN go build -v -o /usr/local/bin/app ./cmd/user-service/

# Stage 2: Create the final image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

# Copy Atlas binary from builder stage
COPY --from=builder /usr/local/bin/atlas /usr/local/bin/atlas

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

COPY --from=builder /usr/src/app/sql/schema.sql ./sql/

CMD ["app"]