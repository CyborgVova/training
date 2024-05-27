.PHONY: all
all:
	go run ./cmd/main.go

.PHONY: goose-create
goose-create:
	@goose -dir ./pkg/db/migrations create new_user_table sql

.PHONY: goose-up
goose-up:
	@goose -dir ./pkg/db/migrations postgres postgres://test:test@localhost:5432/test up

.PHONY: goose-down
goose-down:
	@goose -dir ./pkg/db/migrations postgres postgres://test:test@localhost:5432/test down

.PHONY: compose-up
compose-up:
	docker-compose up -d

.PHONY: compose-down
compose-down:
	docker-compose down
