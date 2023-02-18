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
func (tts *TodoTaskService) TasksList(context.Context, *taskpb.TasksListRequest) (*taskpb.TasksListResponse, error) {
	tts.logger.Log(zap.InfoLevel, "task list called")
	return nil, status.Errorf(codes.Unimplemented, "method TasksList not implemented")
}
func (tts *TodoTaskService) DeleteTask(context.Context, *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	tts.logger.Log(zap.InfoLevel, "delete task called")
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (tts *TodoTaskService) UpdateTask(context.Context, *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	tts.logger.Log(zap.InfoLevel, "update task called")
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
