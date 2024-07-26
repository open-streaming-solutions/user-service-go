package handler

import (
	"context"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/service"
	usergen "github.com/Open-Streaming-Solutions/user-service/pkg/proto"
	"github.com/Open-Streaming-Solutions/user-service/pkg/util"
	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var Module = fx.Module("GrpcHandler",
	fx.Provide(NewGrpcHandler),
)

type GrpcHandler struct {
	usergen.UnimplementedUserServiceServer
	service service.IUserService
	logger  logging.Logger
}

func NewGrpcHandler(service service.IUserService) *GrpcHandler {
	return &GrpcHandler{service: service}
}

func RegisterGrpcHandler(g *grpc.Server, grpcHandler *GrpcHandler) {
	usergen.RegisterUserServiceServer(g, grpcHandler)
}

func (h *GrpcHandler) GetUser(ctx context.Context, req *usergen.GetUserRequest) (*usergen.GetUserResponse, error) {
	id, err := h.service.GetUser(ctx, req.GetUsername())
	if err != nil {
		h.logger.Error("Failed to get user", "username", req.GetUsername(), "error", err)
		return nil, err
	}

	return &usergen.GetUserResponse{UUID: util.ConvertUUIDtoString(id)}, nil
}

func (h *GrpcHandler) CreateUser(ctx context.Context, req *usergen.CreateUserRequest) (*emptypb.Empty, error) {
	uuid := &pgtype.UUID{}
	err := uuid.Scan(req.GetUUID())
	if err != nil {
		h.logger.Error("Failed to scan UUID", "error", err)
		return nil, err
	}

	if err := h.service.CreateUser(ctx, *uuid, req.GetUsername(), req.GetEmail()); err != nil {
		h.logger.Error("Failed to create user", "username", req.GetUsername(), "error", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
