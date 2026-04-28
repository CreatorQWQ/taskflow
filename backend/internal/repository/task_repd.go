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

// UpdateTaskStatus 更新任务状态 (pending <-> completed)
func (r *Repository) UpdateTaskStatus(taskID uint, userID uint, status string) error {
	// 加上 userID 校验，防止用户 A 修改了用户 B 的任务（越权漏洞防护）
	return r.DB.Model(&model.Task{}).
		Where("id = ? AND user_id = ?", taskID, userID).
		Update("status", status).Error
}

// DeleteTask 删除任务 (GORM 默认会执行软删除)
func (r *Repository) DeleteTask(taskID uint, userID uint) error {
	return r.DB.Where("id = ? AND user_id = ?", taskID, userID).
		Delete(&model.Task{}).Error
}
