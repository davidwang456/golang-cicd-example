# Go 项目示例

这是一个 Go 项目示例，展示了完整的项目结构、构建流程和 CI/CD 配置。

## 项目结构

```
.
├── cmd/                    # 主程序入口
│   └── main.go            # 主程序
├── internal/              # 内部包
│   └── app/              # 应用逻辑
│       └── app.go        # 应用实现
├── pkg/                   # 公共包
│   └── utils/            # 工具函数
│       └── utils.go      # 工具实现
├── test/                  # 测试文件
│   └── main_test.go      # 主程序测试
├── .gitlab-ci.yml        # GitLab CI/CD 配置
├── .gitignore            # Git 忽略文件
├── go.mod                # Go 模块定义
├── go.sum                # Go 模块校验和
├── Makefile              # 构建脚本
└── README.md             # 项目文档
```

## 环境要求

- Go 1.21 或更高版本
- GitLab CI/CD 环境
- Nexus 仓库（用于存储构建产物）

## 私有仓库配置

项目使用私有 Go 模块仓库，需要配置以下环境：

1. 配置 GOPRIVATE 环境变量：
```bash
go env -w GOPRIVATE=gitlab.example.com
```

2. 配置 Git 凭证：
```bash
git config --global url."https://oauth2:${GITLAB_TOKEN}@gitlab.example.com/".insteadOf "https://gitlab.example.com/"
```

3. 配置 Go 代理（可选）：
```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

## 构建和测试

### 本地构建

1. 编译项目：
```bash
make build
```

2. 运行测试：
```bash
make test
```

3. 清理构建文件：
```bash
make clean
```

### CI/CD 流程

项目使用 GitLab CI/CD 进行自动化构建和发布，主要包含以下阶段：

1. **构建阶段**
   - 使用 Go 1.21 镜像
   - 配置私有仓库访问
   - 编译项目
   - 生成构建产物

2. **测试阶段**
   - 运行单元测试
   - 生成测试报告

3. **发布阶段**
   - 将构建产物发布到 Nexus 仓库
   - 使用版本号标记发布

## 开发指南

1. 创建新功能分支：
```bash
git checkout -b feature/your-feature-name
```

2. 提交代码：
```bash
git add .
git commit -m "feat: your feature description"
```

3. 推送到远程：
```bash
git push origin feature/your-feature-name
```

## 版本管理

项目使用语义化版本（Semantic Versioning）：

- 主版本号：不兼容的 API 修改
- 次版本号：向下兼容的功能性新增
- 修订号：向下兼容的问题修正

## 贡献指南

1. Fork 项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

[MIT License](LICENSE) 