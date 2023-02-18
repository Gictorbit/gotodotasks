package postgres

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
)

const creatTodoTaskQuery = `
	INSERT INTO 
	    tasks(id,title,description,user_id,due_date,is_completed)
		VALUES(DEFAULT,$1,$2,$3,$4,$5)
	RETURNING id;
`

// CreateTask creates a Task and returns its id
func (ttd *TodoTaskDataBase) CreateTask(ctx context.Context, userID uint32, task *taskpb.TodoTask) error {
	return ttd.RawConn().QueryRow(ctx, creatTodoTaskQuery,
		task.Title,
		task.Description,
		userID,
		task.Date,
		task.IsCompleted).Scan(&task.Id)
}
