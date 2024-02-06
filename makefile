.PHONY: default run build test docs clean

# Variables
APP_NAME=go-jobs
RESET = \033[0m
GREEN = \033[32m
RED = \033[31m
CHECK_MARK = ✔️
CROSS_MARK = ❌

# Tasks
default: run-with-docs

run:
	@echo "$(GREEN)$(CHECK_MARK) Running $(APP_NAME)$(RESET)"
	@go run main.go
	@echo "$(GREEN)$(CHECK_MARK) $(APP_NAME) is running$(RESET)"

run-with-docs:
	@echo "$(GREEN)$(CHECK_MARK) Running $(APP_NAME) with documentation$(RESET)"
	@swag init
	@go run main.go
	@echo "$(GREEN)$(CHECK_MARK) $(APP_NAME) is running$(RESET)"

build:
	@echo "$(GREEN)$(CHECK_MARK) Building $(APP_NAME)$(RESET)"
	@go build -o $(APP_NAME) main.go
	@echo "$(GREEN)$(CHECK_MARK) $(APP_NAME) is built$(RESET)"

test:
	@echo "$(GREEN)$(CHECK_MARK) Running tests$(RESET)"
	@go test -v ./...
	@echo "$(GREEN)$(CHECK_MARK) Tests passed$(RESET)"

docs:
	@echo "$(GREEN)$(CHECK_MARK) Generating swagger documentation$(RESET)"
	@swag init
	@echo "$(GREEN)$(CHECK_MARK) Swagger documentation generated$(RESET)"

clean:
	@echo "$(RED)$(CROSS_MARK) Cleaning up$(RESET)"
	@rm -f $(APP_NAME)
	@rm -f ./docs
	@echo "$(GREEN)$(CHECK_MARK) Cleaned up$(RESET)"
