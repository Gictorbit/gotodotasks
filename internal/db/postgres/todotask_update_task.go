package postgres

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
)

const updateTodoTaskQuery = `
	UPDATE tasks SET
		title=$1,
		description=$2,
		due_date=$3,
		is_completed=$4,
		updated_at=now()
	WHERE 
	    id=$5 AND user_id=$6
	RETURNING id;
`

// UpdateTask updates an existing task
func (ttd *TodoTaskDataBase) UpdateTask(ctx context.Context, userID uint32, task *taskpb.TodoTask) error {
	return ttd.RawConn().QueryRow(ctx, updateTodoTaskQuery,
		task.Title,
		task.Description,
		task.Date,
		task.IsCompleted,
		task.Id, userID).Scan(&task.Id)
}
