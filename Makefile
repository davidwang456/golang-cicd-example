.PHONY: build test clean

# 设置 Go 编译参数
LDFLAGS=-ldflags "-s -w"
BINARY_NAME=go-example

# 默认目标
all: test build

# 构建应用
build:
	go build $(LDFLAGS) -o $(BINARY_NAME) main.go

# 运行测试
test:
	go test -v ./...

# 清理构建产物
clean:
	go clean
	rm -f $(BINARY_NAME)

# 安装依赖
deps:
	go mod download

# 运行应用
run: build
	./$(BINARY_NAME) 