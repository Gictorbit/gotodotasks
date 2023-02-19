package taskmanager

import (
	"context"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	sqlMaker "github.com/Masterminds/squirrel"
)

type TasksListOption func(s sqlMaker.SelectBuilder) sqlMaker.SelectBuilder

func TasksFilterByTitle(filter string) TasksListOption {
	return func(s sqlMaker.SelectBuilder) sqlMaker.SelectBuilder {
		return s.Where(sqlMaker.ILike{"title": "%" + filter + "%"})
	}
}

func TasksFilterByDone(isCompleted bool) TasksListOption {
	return func(s sqlMaker.SelectBuilder) sqlMaker.SelectBuilder {
		return s.Where("is_completed=?", isCompleted)
	}
}

// TasksList returns user's tasks list based on filters
func (ttd *TodoTaskDataBase) TasksList(ctx context.Context, userID uint32, opts ...TasksListOption) ([]*taskpb.TodoTask, error) {
	taskSQL := ttd.GetSQLBuilder().
		Select(
			"id",
			"title",
			"description",
			"due_date",
			"is_completed",
		).From("tasks").
		Where("user_id=?", userID).
		OrderBy("updated_at", "created_at")

	for _, opt := range opts {
		taskSQL = opt(taskSQL)
	}
	query, params, err := taskSQL.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := ttd.RawConn().Query(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tasks := make([]*taskpb.TodoTask, 0)
	for rows.Next() {
		task := &taskpb.TodoTask{}
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Date, &task.IsCompleted); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
