package handler

import (
	"github.com/CreatorQWQ/taskflow/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskHandler struct {
	svc *service.TaskService
}

func NewTaskHandler(svc *service.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

// CreateTask 处理 POST /api/v1/tasks
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "标题不能为空"})
		return
	}

	// 从中间件中取出当前用户的 ID
	userID, _ := c.Get("current_user_id")

	if err := h.svc.CreateTask(input.Title, input.Content, userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建任务失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// ListTasks 处理 GET /api/v1/tasks
func (h *TaskHandler) ListTasks(c *gin.Context) {
	userID, _ := c.Get("current_user_id")

	tasks, err := h.svc.GetUserTasks(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取任务列表失败"})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// ToggleStatus 切换状态
func (h *TaskHandler) ToggleStatus(c *gin.Context) {
	// 将字符串 ID 转换为 uint
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID, _ := c.Get("current_user_id")
	if err := h.svc.ToggleTaskStatus(uint(taskID), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// Delete 删除
func (h *TaskHandler) Delete(c *gin.Context) {
	taskIDStr := c.Param("id")
	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID"})
		return
	}

	userID, _ := c.Get("current_user_id")
	if err := h.svc.DeleteTask(uint(taskID), userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}