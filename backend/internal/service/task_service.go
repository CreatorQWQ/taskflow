package service

import (
	"errors"

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

// ToggleTaskStatus 切换状态逻辑
func (s *TaskService) ToggleTaskStatus(taskID uint, userID uint) error {
	// 1. 先找到这个任务
	var task model.Task
	err := s.repo.DB.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error
	if err != nil {
		return errors.New("任务不存在或无权操作")
	}

	// 2. 切换状态
	newStatus := "completed"
	if task.Status == "completed" {
		newStatus = "pending"
	}

	// 3. 更新
	return s.repo.DB.Model(&task).Update("status", newStatus).Error
}

// DeleteTask 删除逻辑
func (s *TaskService) DeleteTask(taskID uint, userID uint) error {
	return s.repo.DB.Where("id = ? AND user_id = ?", taskID, userID).Delete(&model.Task{}).Error
}
