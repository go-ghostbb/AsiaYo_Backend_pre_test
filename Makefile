build:
	@echo "開始編譯..."
	@go mod tidy
	@go build -o bin/test main.go

docker:
	@echo "開始打包 Docker Image - AsiaYo_Backend_pre_test"
	@docker build -t="pre_test" .