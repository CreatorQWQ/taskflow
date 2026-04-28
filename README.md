# TaskFlow - 任务管理应用

一个全栈任务管理应用，使用 Go + Gin 作为后端，Flutter 作为前端，PostgreSQL 作为数据库。

## 📋 项目概述

TaskFlow 是一个现代化的任务管理应用，支持用户注册、登录、任务创建、状态切换和删除等功能。采用前后端分离架构，支持多平台部署（Android、iOS、Web、Windows、Linux、macOS）。

### 主要功能

- ✅ 用户注册与登录（JWT 认证）
- ✅ 任务创建与管理
- ✅ 任务状态切换（完成/未完成）
- ✅ 任务列表查看
- ✅ 任务统计卡片展示
- ✅ 安全的密码存储（bcrypt）
- ✅ 安全的 Token 存储（Flutter Secure Storage）

## 🏗️ 技术栈

### 后端 (Backend)

| 技术 | 说明 |
|------|------|
| Go 1.25.5 | 编程语言 |
| Gin | Web 框架 |
| GORM | ORM 框架 |
| PostgreSQL 15 | 关系型数据库 |
| JWT (golang-jwt) | 身份认证 |
| Viper | 配置管理 |
| bcrypt | 密码加密 |
| Air | 热重载开发工具 |

### 前端 (Frontend)

| 技术 | 说明 |
|------|------|
| Flutter 3.x | UI 框架 |
| Dart 3.11.4 | 编程语言 |
| Riverpod 3.3.1 | 状态管理 |
| Dio 5.9.2 | HTTP 客户端 |
| Flutter Secure Storage | 安全存储 |
| Material 3 | UI 设计系统 |

## 📁 项目结构

```
taskflow/
├── backend/                    # Go 后端服务
│   ├── internal/              # 内部包
│   │   ├── config/           # 配置管理
│   │   │   └── config.go
│   │   ├── handler/          # HTTP 处理器
│   │   │   ├── auth_handler.go   # 认证相关
│   │   │   ├── task_handler.go   # 任务相关
│   │   │   └── handler.go
│   │   ├── middleware/       # 中间件
│   │   │   └── auth_middleware.go  # JWT 认证中间件
│   │   ├── model/            # 数据模型
│   │   │   ├── model.go
│   │   │   ├── user.go
│   │   │   └── task.go
│   │   ├── repository/       # 数据库操作层
│   │   │   ├── db.go
│   │   │   ├── user_repo.go
│   │   │   └── task_repd.go
│   │   ├── router/           # 路由配置
│   │   │   └── router.go
│   │   └── service/          # 业务逻辑层
│   │       ├── user_service.go
│   │       └── task_service.go
│   ├── pkg/                  # 公共包
│   │   ├── logger/          # 日志
│   │   └── utils/           # 工具函数
│   │       └── jwt.go       # JWT 工具
│   ├── .env                 # 环境变量配置
│   ├── .air.toml            # Air 热重载配置
│   ├── Dockerfile           # Docker 构建文件
│   ├── go.mod               # Go 依赖管理
│   └── main.go              # 入口文件
│
├── frontend/                 # Flutter 前端应用
│   ├── lib/
│   │   ├── src/
│   │   │   ├── api/         # API 客户端
│   │   │   │   └── api_client.dart
│   │   │   ├── features/    # 功能模块
│   │   │   │   ├── auth/    # 认证模块
│   │   │   │   │   ├── provider/
│   │   │   │   │   │   ├── auth_provider.dart
│   │   │   │   │   │   └── auth_state.dart
│   │   │   │   │   └── view/
│   │   │   │   │       └── login_screen.dart
│   │   │   │   └── tasks/   # 任务模块
│   │   │   │       ├── model/
│   │   │   │       │   └── task_model.dart
│   │   │   │       ├── provider/
│   │   │   │       │   └── task_provider.dart
│   │   │   │       └── view/
│   │   │   │           ├── widgets/
│   │   │   │           │   └── task_summary_card.dart
│   │   │   │           └── home_screen.dart
│   │   │   └── app.dart     # 应用主组件
│   │   └── main.dart        # 入口文件
│   ├── android/             # Android 平台
│   ├── ios/                 # iOS 平台
│   ├── web/                 # Web 平台
│   ├── windows/             # Windows 平台
│   ├── linux/               # Linux 平台
│   ├── macos/               # macOS 平台
│   └── pubspec.yaml         # Flutter 依赖管理
│
└── docker-compose.yaml       # Docker 编排文件
```

## 🚀 快速开始

### 环境要求

- Go 1.25.5+
- Flutter 3.x
- PostgreSQL 15+（或使用 Docker）
- Docker & Docker Compose（可选，用于容器化部署）

### 方式一：本地开发

#### 1. 启动数据库

```bash
# 使用 Docker 启动 PostgreSQL
docker run -d --name postgres \
  -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=taskflow \
  -p 5432:5432 \
  postgres:15-alpine
```

#### 2. 配置后端

```bash
cd backend

# 复制并编辑环境变量
# .env 文件内容：
# DB_HOST=localhost
# DB_USER=postgres
# DB_PASSWORD=123456
# DB_NAME=taskflow
# DB_PORT=5432
# SERVER_PORT=8080
# JWT_SECRET=your_secret_key

# 安装依赖
go mod download

# 启动服务（开发模式，支持热重载）
air

# 或直接运行
go run main.go
```

#### 3. 配置前端

```bash
cd frontend

# 安装依赖
flutter pub get

# 配置后端地址
# 编辑 lib/src/api/api_client.dart
# 修改 baseUrl 为你的后端地址

# 运行应用（选择设备）
flutter devices  # 查看可用设备
flutter run -d <device-id>

# 热重载：按 r
# 热重启：按 R
# 退出：按 q
```

### 方式二：Docker Compose 部署

```bash
# 启动所有服务（数据库 + 后端）
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down
```

## 📡 API 接口

### 认证模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/api/v1/auth/register` | 用户注册 | ❌ |
| POST | `/api/v1/auth/login` | 用户登录 | ❌ |
| GET | `/api/v1/user/me` | 获取当前用户信息 | ✅ |

### 任务模块

| 方法 | 路径 | 说明 | 认证 |
|------|------|------|------|
| POST | `/api/v1/tasks/` | 创建任务 | ✅ |
| GET | `/api/v1/tasks/` | 获取任务列表 | ✅ |
| PATCH | `/api/v1/tasks/:id/toggle` | 切换任务状态 | ✅ |
| DELETE | `/api/v1/tasks/:id` | 删除任务 | ✅ |

### 请求示例

#### 用户注册

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

#### 用户登录

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
```

响应：
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### 创建任务

```bash
curl -X POST http://localhost:8080/api/v1/tasks/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_token>" \
  -d '{"title":"完成项目文档","content":"编写 README.md"}'
```

#### 获取任务列表

```bash
curl -X GET http://localhost:8080/api/v1/tasks/ \
  -H "Authorization: Bearer <your_token>"
```

#### 切换任务状态

```bash
curl -X PATCH http://localhost:8080/api/v1/tasks/1/toggle \
  -H "Authorization: Bearer <your_token>"
```

## 🗄️ 数据库设计

### 用户表 (users)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| username | string | 用户名（唯一） |
| password | string | 密码（bcrypt 加密） |
| created_at | time | 创建时间 |
| updated_at | time | 更新时间 |

### 任务表 (tasks)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| user_id | uint | 外键（关联用户） |
| title | string | 任务标题 |
| content | string | 任务内容 |
| is_completed | bool | 是否完成 |
| created_at | time | 创建时间 |
| updated_at | time | 更新时间 |

## 🔐 安全说明

- 密码使用 bcrypt 加密存储
- JWT Token 用于身份认证，有效期可配置
- 敏感操作需要 JWT 认证
- Flutter 端使用 Secure Storage 存储 Token
- 生产环境请修改 `.env` 中的 `JWT_SECRET`

## 🛠️ 开发工具

### 后端

- **Air**: 开发时热重载工具
  ```bash
  go install github.com/cosmtrek/air@latest
  ```

### 前端

- **Flutter DevTools**: 调试和性能分析
  ```bash
  flutter pub global activate devtools
  ```

## 📝 开发规范

### 后端架构

采用经典的分层架构：

```
Router → Handler → Service → Repository → Database
```

- **Router**: 路由定义和中间件配置
- **Handler**: 处理 HTTP 请求和响应
- **Service**: 业务逻辑处理
- **Repository**: 数据库操作

### 前端架构

采用 Feature-First + Riverpod 状态管理：

```
View → Provider → API Client → Backend
```

- **View**: UI 组件（ConsumerWidget）
- **Provider**: 状态管理（Notifier/AsyncNotifier）
- **API Client**: HTTP 请求封装（Dio）

## 🐛 常见问题

### 1. 连接超时错误

**问题**: `DioException [connection timeout]`

**解决方案**:
- 确认后端服务已启动
- 检查 `api_client.dart` 中的 `baseUrl` 配置
- Android 模拟器使用 `http://10.0.2.2:8080`
- 真机调试使用局域网 IP，如 `http://192.168.0.x:8080`

### 2. 404 错误

**问题**: API 返回 404

**解决方案**:
- 确认后端服务已重启（添加新路由后）
- 检查请求路径是否正确
- 查看后端启动日志确认路由已注册

### 3. 认证失败

**问题**: 返回 "请求未授权" 或 "无效的令牌"

**解决方案**:
- 确认已登录并获取 Token
- 检查请求头是否包含 `Authorization: Bearer <token>`
- 确认 Token 未过期

## 📄 许可证

本项目仅供学习和个人使用。

## 👨‍💻 作者

- GitHub: [CreatorQWQ](https://github.com/CreatorQWQ)

---

**祝开发愉快！** 🚀
