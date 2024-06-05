package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline"`
}

// EXAM

type UpdateTaskRequest struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline"`
	Status      string `json:"status"`
}

// EXAM END

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	var deadline *time.Time
	if r.Deadline != 0 {
		dl := time.Unix(r.Deadline, 0)
		deadline = &dl
	}
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
	}, nil
}

// EXAM

func (r UpdateTaskRequest) ToDomainModel() (interface{}, error) {
	var deadline *time.Time
	if r.Deadline != 0 {
		dl := time.Unix(r.Deadline, 0)
		deadline = &dl
	}
	return domain.Task{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
		Status:      domain.TaskStatus(r.Status),
	}, nil
}

// EXAM END
