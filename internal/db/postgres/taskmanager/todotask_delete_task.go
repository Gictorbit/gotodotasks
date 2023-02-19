package taskmanager

import (
	"context"
)

const deleteTodoTaskQuery = `
	DELETE FROM tasks WHERE user_id=$1 AND id=$2 RETURNING id;
`

// DeleteTask removes and exists task
func (ttd *TodoTaskDataBase) DeleteTask(ctx context.Context, userID uint32, taskID uint32) error {
	return ttd.RawConn().QueryRow(ctx, deleteTodoTaskQuery, userID, taskID).Scan(&taskID)
}
