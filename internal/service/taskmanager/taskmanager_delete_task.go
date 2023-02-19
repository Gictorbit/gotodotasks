package taskmanager

import (
	"context"
	"errors"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteTask removes exists task
func (tts *TodoTaskService) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := tts.ValidateDeleteTodoTask(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	claims, ok := authutil.TokenClaimsFromCtx(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "invalid context")
	}
	if err := tts.dbConn.DeleteTask(ctx, claims.UserID, req.GetId()); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "task not found")
		}
		tts.logger.Error("delete task failed",
			zap.Error(err),
			zap.Uint32("userID", claims.UserID),
			zap.Uint32("taskID", req.GetId()),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &taskpb.DeleteTaskResponse{
		Code: taskpb.ResponseCode_RESPONSE_CODE_SUCCESS,
	}, nil
}
