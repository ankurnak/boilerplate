package controllers

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController -> Save: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController -> Save: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskController) GetForUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		tasks, err := c.taskService.GetForUser(user.Id)
		if err != nil {
			log.Printf("TaskController -> GetForUser: %s", err)
			InternalServerError(w, err)
			return
		}

		var tasksDto resources.TasksDto
		tasksDto = tasksDto.DomainToDtoCollection(tasks)
		Success(w, tasksDto)
	}
}

// EXAM

func (c TaskController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			log.Printf("TaskController -> FindById: %s", err)
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		task, err := c.taskService.FindById(id)
		if err != nil {
			log.Printf("TaskController -> FindById: %s", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		var taskDto resources.TaskDto
		taskDto = taskDto.DomainToDto(task)
		Success(w, taskDto)
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		task, err := requests.Bind(r, requests.UpdateTaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		t, err := c.taskService.FindById(task.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			NotFound(w, err)
			return
		}

		t.Title = task.Title
		t.Description = task.Description
		t.Status = task.Status
		t.Deadline = task.Deadline

		task, err = c.taskService.Update(t)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var taskDto resources.TaskDto
		Success(w, taskDto.DomainToDto(task))
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			log.Printf("TaskController: invalid ID format")
			BadRequest(w, err)
			return
		}

		err = c.taskService.Delete(id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}

// EXAM END
