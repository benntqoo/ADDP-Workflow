# AI Terminal Proxy Makefile
# 支援 TDD 開發流程的構建和測試命令

.PHONY: help build test test-verbose test-race test-coverage clean install deps lint fmt vet benchmark

# 默認目標
help: ## 顯示所有可用命令
	@echo "AI Terminal Proxy - Makefile Commands"
	@echo "====================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# 構建相關
build: ## 構建主程序
	@echo "🔨 構建 claude-proxy..."
	go build -o bin/claude-proxy ./cmd/proxy

build-all: ## 構建所有平台的二進制文件
	@echo "🔨 構建所有平台..."
	GOOS=windows GOARCH=amd64 go build -o bin/claude-proxy-windows-amd64.exe ./cmd/proxy
	GOOS=linux GOARCH=amd64 go build -o bin/claude-proxy-linux-amd64 ./cmd/proxy
	GOOS=darwin GOARCH=amd64 go build -o bin/claude-proxy-darwin-amd64 ./cmd/proxy
	GOOS=darwin GOARCH=arm64 go build -o bin/claude-proxy-darwin-arm64 ./cmd/proxy

# 測試相關 (TDD 核心)
test: ## 運行所有測試 (TDD)
	@echo "🧪 運行單元測試..."
	go test ./...

test-verbose: ## 運行詳細模式測試
	@echo "🧪 運行詳細測試..."
	go test -v ./...

test-race: ## 運行競態檢測測試
	@echo "🧪 運行競態檢測測試..."
	go test -race ./...

test-coverage: ## 生成測試覆蓋率報告
	@echo "📊 生成測試覆蓋率報告..."
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "覆蓋率報告生成在 coverage.html"

test-benchmark: ## 運行性能基準測試
	@echo "⚡ 運行性能基準測試..."
	go test -bench=. ./...

# TDD 開發循環
tdd: ## TDD 開發循環 (監控文件變化並自動測試)
	@echo "🔄 啟動 TDD 監控模式..."
	@which fswatch > /dev/null || (echo "請安裝 fswatch: brew install fswatch" && exit 1)
	fswatch -o . | xargs -n1 -I{} make test

red: ## TDD Red 階段 - 運行測試（應該失敗）
	@echo "🔴 TDD RED 階段 - 運行測試..."
	@go test ./... || echo "✅ 測試失敗是預期的（RED 階段）"

green: ## TDD Green 階段 - 運行測試（應該通過）
	@echo "🟢 TDD GREEN 階段 - 運行測試..."
	go test ./...

refactor: ## TDD Refactor 階段 - 運行所有檢查
	@echo "🔵 TDD REFACTOR 階段 - 完整檢查..."
	make fmt
	make vet
	make lint
	make test
	make test-race

# 代碼質量
deps: ## 安裝依賴
	@echo "📦 安裝依賴..."
	go mod download
	go mod tidy

fmt: ## 格式化代碼
	@echo "✨ 格式化代碼..."
	go fmt ./...

vet: ## 運行 go vet 檢查
	@echo "🔍 運行 go vet..."
	go vet ./...

lint: ## 運行 golangci-lint 檢查
	@echo "🔍 運行 golangci-lint..."
	@which golangci-lint > /dev/null || (echo "請安裝 golangci-lint: https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

# 工具相關
install: build ## 安裝到 $GOPATH/bin
	@echo "📦 安裝 claude-proxy..."
	go install ./cmd/proxy

clean: ## 清理構建產物
	@echo "🧹 清理構建產物..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# 開發輔助
dev: ## 開發模式 - 構建並運行
	@echo "🚀 開發模式啟動..."
	go run ./cmd/proxy

demo-terminals: ## 演示終端列表功能
	@echo "📺 演示終端列表..."
	go run ./cmd/proxy list-terminals

version: ## 顯示版本信息
	@echo "📋 版本信息..."
	go run ./cmd/proxy version

# CI/CD 相關
ci: ## CI 流程 - 完整檢查
	@echo "🔄 CI 流程執行..."
	make deps
	make fmt
	make vet
	make lint
	make test
	make test-race
	make test-coverage
	make build

# 項目初始化
init: ## 初始化項目環境
	@echo "🎯 初始化項目環境..."
	go mod download
	@echo "安裝開發工具..."
	@which golangci-lint > /dev/null || (echo "正在安裝 golangci-lint..." && curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin)
	@echo "✅ 項目環境初始化完成"

# Docker 相關（可選）
docker-build: ## 構建 Docker 鏡像
	@echo "🐳 構建 Docker 鏡像..."
	docker build -t claude-proxy:latest .

docker-run: ## 運行 Docker 容器
	@echo "🐳 運行 Docker 容器..."
	docker run -it --rm claude-proxy:latest

# 文檔生成
docs: ## 生成 Go 文檔
	@echo "📚 生成文檔..."
	godoc -http=:6060 &
	@echo "文檔服務啟動在 http://localhost:6060"