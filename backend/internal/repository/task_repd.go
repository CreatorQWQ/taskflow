package repository

import (
	"github.com/CreatorQWQ/taskflow/internal/model"
)

// CreateTask 创建新任务
func (r *Repository) CreateTask(task *model.Task) error {
	return r.DB.Create(task).Error
}

// GetTasksByUserID 获取某个用户的所有任务
func (r *Repository) GetTasksByUserID(userID uint) ([]model.Task, error) {
	var tasks []model.Task
	// 按照创建时间倒序排列
	err := r.DB.Where("user_id = ?", userID).Order("created_at desc").Find(&tasks).Error
	return tasks, err
}
