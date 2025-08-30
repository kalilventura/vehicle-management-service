.PHONY: dev-up
dev-up:
	docker compose -f .dev/compose.yaml up  --build -d

.PHONY: goose-up
goose-up:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir db/migrations postgres "host=localhost port=5432 user=postgres password=postgres dbname=vehicle_management sslmode=disable" up

.PHONY: migrate
migrate:
	make docker-up
	@sleep 1.5
	make goose-up

.PHONY: build
build:
	make wire
	make swagger
	go mod tidy
	mkdir -p bin && go build -o ./bin -v ./...

.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd cmd; wire

.PHONY: swagger
swagger:
	go install github.com/swaggo/swag/cmd/swag@latest
	swag init -g ./cmd/main.go --parseDependency -o ./cmd/docs

.PHONY: test
test:
	.github/scripts/tests/run_tests.sh

.PHONY: lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run \
		--config ".golangci.yaml" \
		--color "always" \
		--timeout "10m" \
		--print-resources-usage \
		--allow-parallel-runners \
		--max-issues-per-linter 0 \
		--max-same-issues 0 ./...

.PHONY: goimports
goimports:
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -l -w .

.PHONY: gci
gci:
	go install github.com/daixiang0/gci@latest
	gci write --skip-generated -s standard -s default .

.PHONY: lint-fix
lint-fix:
	make goimports
	make gci
	make lint
