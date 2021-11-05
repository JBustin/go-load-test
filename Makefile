.DEFAULT_GOAL := help

help: ## Other: Display this help
	@grep -E '^[a-zA-Z1-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN { FS = ":.*?## " }; { printf "\033[36m%-30s\033[0m %s\n", $$1, $$2 }'

install: ## Build binary
	go build

test: ## Start unit test
	go test -cover ./...

install-hook-macos: ## Install precommit git hook (macos)
	@echo $$"#! /bin/bash\nmake test" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit

install-hook: ## Install precommit git hook (linux, win)
	@echo -e "#! /bin/bash\nmake test" > .git/hooks/pre-commit
	@chmod +x .git/hooks/pre-commit

uninstall-hook: ## Uninstall precommit git hook
	@rm -f .git/hooks/pre-commit
	