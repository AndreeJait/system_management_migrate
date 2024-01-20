
.PHONY: migrate-new-system-management
migrate-new-system-management: ## create a new database migration
	@read -p "migration name: " name; \
	CGO_ENABLED="0" go install github.com/rubenv/sql-migrate/...@latest; \
	sql-migrate new -config=system_management_config.yml $${name}

.PHONY: migrate-up-system-management
migrate-up-system-management:
	@read -p "env mode: " mode; \
	go run main.go up --mode=$${env} --module=management_system

.PHONY: migrate-down-system-management
migrate-down-system-management:
	@read -p "env mode: " mode; \
	go run main.go down --mode=$${env} --module=management_system

.PHONY: migrate-fresh-system-management
migrate-fresh-system-management:
	@read -p "env mode: " mode; \
	go run main.go fresh --mode=$${env} --module=management_system
