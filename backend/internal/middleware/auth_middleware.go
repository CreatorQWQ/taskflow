package middleware

import (
	"net/http"
	"strings"

	"github.com/CreatorQWQ/taskflow/pkg/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 身份验证中间件
func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 从 Header 中获取 Authorization
		// 标准格式是: Authorization: Bearer <TOKEN>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求未授权"})
			c.Abort() // 终止后续操作
			return
		}

		// 2. 解析 Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "认证格式错误"})
			c.Abort()
			return
		}

		// 3. 校验 Token
		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
			c.Abort()
			return
		}

		// 4. 【核心点】将解析出的 UserID 存入上下文
		// 这样后续的 Handler 就能直接知道是谁在操作了
		c.Set("current_user_id", claims.UserID)
		c.Next() // 通过校验，继续执行后续逻辑
	}
}