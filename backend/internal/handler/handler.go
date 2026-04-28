package handler

// AllHandlers 作为一个容器，装载所有的业务 Handler
type AllHandlers struct {
	Auth *AuthHandler
	Task *TaskHandler
	// 以后加了功能直接往这加：
	// Comment *CommentHandler
}

// NewAllHandlers 构造函数
func NewAllHandlers(auth *AuthHandler, task *TaskHandler) *AllHandlers {
	return &AllHandlers{
		Auth: auth,
		Task: task,
	}
}
