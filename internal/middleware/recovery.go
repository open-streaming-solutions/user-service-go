package middleware

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"runtime/debug"
)

// RecoveryHandlerFunc is a gRPC server interceptor for recovering from panics.
func RecoveryHandlerFunc(p any) (err error) {
	slog.Error("Recovered from panic: %v", p)
	slog.Error("Stack trace: %s", debug.Stack())
	return status.Errorf(codes.Internal, "Internal server error")
}
