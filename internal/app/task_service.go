package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	GetForUser(uId uint64) ([]domain.Task, error)

	// EXAM
	FindById(id uint64) (domain.Task, error)
	Update(t domain.Task) (domain.Task, error)
	Delete(id uint64) error
	// EXAM END
}

type taskService struct {
	taskRepo database.TaskRepository
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService -> Save: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) GetForUser(uId uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.GetByUserId(uId)
	if err != nil {
		log.Printf("TaskService -> GetForUser: %s", err)
		return nil, err
	}
	return tasks, nil
}

// EXAM

func (s taskService) FindById(id uint64) (domain.Task, error) {
	task, err := s.taskRepo.FindById(id)
	if err != nil {
		log.Printf("TaskService -> FindById: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Update(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Update(t)
	if err != nil {
		log.Printf("TaskService -> Update: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Delete(id uint64) error {
	err := s.taskRepo.Delete(id)
	if err != nil {
		log.Printf("TaskService -> Delete: %s", err)
		return err
	}
	return nil
}

// EXAM END
