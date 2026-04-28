package model

// GetModels 返回项目中所有的数据库模型实例
// 使用 interface{} 切片，方便 GORM 批量处理
func GetModels() []interface{} {
	return []interface{}{
		&User{},
		&Task{},
		// 以后加了新功能，比如 &Comment{}, 只需要在这里加一行
		// 业务逻辑代码不需要再动
	}
}
