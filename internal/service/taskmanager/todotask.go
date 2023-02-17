package taskmanager

import (
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	taskdb "github.com/Gictorbit/gotodotasks/internal/db/postgres"
	"go.uber.org/zap"
)

type TodoTaskService struct {
	logger *zap.Logger
	dbConn taskdb.TodoTaskDBInterface
	taskpb.UnimplementedTodoTaskServiceServer
}

// NewTodoTaskManager returns new TodoTaskService
func NewTodoTaskManager(logger *zap.Logger, dbConn taskdb.TodoTaskDBInterface) *TodoTaskService {
	return &TodoTaskService{
		logger: logger,
		dbConn: dbConn,
	}
}
