package service

import (
	"github.com/CreatorQWQ/taskflow/internal/model"
	"github.com/CreatorQWQ/taskflow/internal/repository"
)

type TaskService struct {
	repo *repository.Repository
}

func NewTaskService(repo *repository.Repository) *TaskService {
	return &TaskService{repo: repo}
}

// CreateTask 业务：创建任务
func (s *TaskService) CreateTask(title, content string, userID uint) error {
	task := &model.Task{
		Title:   title,
		Content: content,
		UserID:  userID,
		Status:  "pending",
	}
	return s.repo.CreateTask(task)
}

// GetUserTasks 业务：获取用户任务列表
func (s *TaskService) GetUserTasks(userID uint) ([]model.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}