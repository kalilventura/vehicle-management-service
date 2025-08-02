.PHONY: build
build:
	make wire
	go mod tidy
	mkdir -p bin && go build -o ./bin -v ./...

.PHONY: wire
wire:
	go install github.com/google/wire/cmd/wire@latest
	cd cmd; wire

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
