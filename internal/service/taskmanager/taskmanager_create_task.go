package taskmanager

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateTask creates a new task
func (tts *TodoTaskService) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if err := tts.ValidateCreateTodoTask(req); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	userID := uint32(0)
	if err := tts.dbConn.CreateTask(ctx, userID, req.Task); err != nil {
		tts.logger.Error("create task failed",
			zap.Error(err),
			zap.Uint32("userID", userID),
		)
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	return &taskpb.CreateTaskResponse{
		TaskId: req.Task.GetId(),
	}, nil
}
