package taskmanager

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	taskdb "github.com/Gictorbit/gotodotasks/internal/db/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TasksList returns users tasks based on filters
func (tts *TodoTaskService) TasksList(ctx context.Context, req *taskpb.TasksListRequest) (*taskpb.TasksListResponse, error) {
	userID := uint32(0)
	opts := make([]taskdb.TasksListOption, 0)
	if len(req.Title) > 0 {
		opts = append(opts, taskdb.TasksFilterByTitle(req.Title))
	}
	if req.IsCompleted == taskpb.TasksCompletedFilter_TASKS_COMPLETED_FILTER_DONE {
		opts = append(opts, taskdb.TasksFilterByDone(true))
	} else if req.IsCompleted == taskpb.TasksCompletedFilter_TASKS_COMPLETED_FILTER_UNDONE {
		opts = append(opts, taskdb.TasksFilterByDone(false))
	}
	tasks, err := tts.dbConn.TasksList(ctx, userID, opts...)
	if err != nil {
		tts.logger.Error("failed to get tasks list",
			zap.Error(err),
			zap.Uint32("userID", userID),
		)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &taskpb.TasksListResponse{
		Tasks: tasks,
	}, nil
}
