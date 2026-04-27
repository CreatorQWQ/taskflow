package service

import (
	"errors"
	"github.com/CreatorQWQ/taskflow/internal/model"
	"github.com/CreatorQWQ/taskflow/internal/repository"
	"github.com/CreatorQWQ/taskflow/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.Repository
	jwtSecret string
}

func NewUserService(repo *repository.Repository, secret string) *UserService {
	return &UserService{repo: repo, jwtSecret: secret}
}

// Register 注册逻辑
func (s *UserService) Register(username, password, email string) error {
	// 1. 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}
	return s.repo.CreateUser(user)
}


// Login 登录逻辑
func (s *UserService) Login(username, password string) (string, error) {
	// 1. 查找用户
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 2. 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("密码错误")
	}

	// 3. 生成 Token
	return utils.GenerateToken(user.ID, s.jwtSecret)
}