package taskmanager

import (
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	"github.com/Gictorbit/gotodotasks/internal/authutil"
	taskdb "github.com/Gictorbit/gotodotasks/internal/db/postgres/taskmanager"
	"go.uber.org/zap"
)

type TodoTaskService struct {
	logger *zap.Logger
	dbConn taskdb.TodoTaskDBInterface
	auth   *authutil.AuthManager
	taskpb.UnimplementedTodoTaskServiceServer
}

// NewTodoTaskManager returns new TodoTaskService
func NewTodoTaskManager(logger *zap.Logger, dbConn taskdb.TodoTaskDBInterface, auth *authutil.AuthManager) *TodoTaskService {
	return &TodoTaskService{
		logger: logger,
		dbConn: dbConn,
		auth:   auth,
	}
}

func (tts *TodoTaskService) GetAuthManager() *authutil.AuthManager {
	return tts.auth
}
