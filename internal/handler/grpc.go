package handler

import (
	"context"
	xerrors "errors" // I'm so sorry =(
	"github.com/Open-Streaming-Solutions/user-service/internal/errors"
	"github.com/Open-Streaming-Solutions/user-service/internal/logging"
	"github.com/Open-Streaming-Solutions/user-service/internal/service"
	protouser "github.com/Open-Streaming-Solutions/user-service/pkg/proto"
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
	uuid := &pgtype.UUID{}
	err := uuid.Scan(req.GetUUID())
	if err != nil {
		h.logger.Error("Failed to scan UUID", "error", err)
		return nil, xerrors.Join(errors.ErrInvalidUUID, err)
	}

	if err := h.service.CreateUser(ctx, *uuid, req.GetUsername(), req.GetEmail()); err != nil {
		h.logger.Error("Failed to create user", "username", req.GetUsername(), "error", err)
		return nil, errors.ToGrpcError(err)
	}

	return &emptypb.Empty{}, nil
}
