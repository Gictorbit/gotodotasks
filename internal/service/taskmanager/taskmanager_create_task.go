package taskmanager

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (tts *TodoTaskService) CreateTask(context.Context, *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	tts.logger.Log(zap.InfoLevel, "create task called")
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
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
