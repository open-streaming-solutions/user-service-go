package handler

import (
	"context"
	"github.com/open-streaming-solutions/user-service/internal/errors"
	"github.com/open-streaming-solutions/user-service/internal/logging"
	"github.com/open-streaming-solutions/user-service/internal/service"
	protouser "github.com/open-streaming-solutions/user-service/pkg/proto"
	"github.com/open-streaming-solutions/user-service/pkg/util"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var Module = fx.Module("GrpcHandler",
	fx.Provide(NewGrpcHandler),
)

type GrpcHandler struct {
	protouser.UnimplementedUserServiceServer
	service service.IUserService
	logger  logging.Logger
}

func NewGrpcHandler(logger logging.Logger, service service.IUserService) *GrpcHandler {
	return &GrpcHandler{
		service: service,
		logger:  logger,
	}
}

func RegisterGrpcHandler(g *grpc.Server, grpcHandler *GrpcHandler) {
	protouser.RegisterUserServiceServer(g, grpcHandler)
}

func (h *GrpcHandler) GetUser(ctx context.Context, req *protouser.GetUserRequest) (*protouser.GetUserResponse, error) {
	id, err := h.service.GetUser(ctx, req.GetUsername())
	if err != nil {
		h.logger.Error("Failed to get user", "username", req.GetUsername(), "error", err)
		return nil, errors.ToGrpcError(err)
	}

	return &protouser.GetUserResponse{UUID: util.ConvertUUIDtoString(id)}, nil
}

func (h *GrpcHandler) CreateUser(ctx context.Context, req *protouser.CreateUserRequest) (*emptypb.Empty, error) {
	if err := h.service.CreateUser(ctx, req.GetUUID(), req.GetUsername(), req.GetEmail()); err != nil {
		h.logger.Error("Failed to create user", "username", req.GetUsername(), "error", err)
		return nil, errors.ToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
