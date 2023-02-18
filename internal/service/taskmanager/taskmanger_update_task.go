package taskmanager

import (
	"context"
	"errors"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateTask updates an exists task
func (tts *TodoTaskService) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if err := tts.ValidateUpdateTodoTask(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	userID := uint32(0)
	if err := tts.dbConn.UpdateTask(ctx, userID, req.Task); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "task not found")
		}
		tts.logger.Error("update task failed",
			zap.Error(err),
			zap.Uint32("userID", userID),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &taskpb.UpdateTaskResponse{
		Code: taskpb.ResponseCode_RESPONSE_CODE_SUCCESS,
	}, nil
}
