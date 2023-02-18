package taskmanager

import (
	"errors"
	taskpb "github.com/Gictorbit/gotodotasks/api/gen/proto/todotask/v1"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

func (tts *TodoTaskService) ValidateCreateTodoTask(req *taskpb.CreateTaskRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Task, validation.Required, validation.By(func(value interface{}) error {
			return ValidateTodoTask(req.Task, false)
		})),
	)
}

func (tts *TodoTaskService) ValidateUpdateTodoTask(req *taskpb.UpdateTaskRequest) error {
	return validation.ValidateStruct(req,
		validation.Field(&req.Task, validation.Required, validation.By(func(value interface{}) error {
			return ValidateTodoTask(req.Task, true)
		})),
	)
}

func ValidateTodoTask(task *taskpb.TodoTask, isUpdate bool) error {
	dateValidation := func(value interface{}) error {
		now := time.Now().UTC().Unix()
		date := value.(uint64)
		if int64(date) < now {
			return errors.New("date must be greater than current time")
		}
		return nil
	}
	return validation.ValidateStruct(task,
		validation.Field(&task.Id, validation.When(isUpdate, validation.Required).Else(validation.Empty)),
		validation.Field(&task.IsCompleted, validation.When(!isUpdate, validation.Empty)),
		validation.Field(&task.Title, validation.Required),
		validation.Field(&task.Date, validation.When(task.Date != 0, validation.By(dateValidation))),
	)
}
