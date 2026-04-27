package repository

import (
		"github.com/CreatorQWQ/taskflow/internal/model"

)

// CreateUser 创建用户
func (r *Repository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}


func (r *Repository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}