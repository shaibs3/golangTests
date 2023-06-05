TEST_COVERAGE_FILE := cover.out
BIN_DIR := .tools/bin
GOLANGCI_LINT := $(BIN_DIR)/golangci-lint
GOLANGCI_LINT_VERSION := 1.52.2
GO := go
ifdef GO_BIN
	GO = $(GO_BIN)
endif

build:
	$(GO) build ./...

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

test:
	$(GO) test -v -race ./... -coverprofile ${TEST_COVERAGE_FILE}

test_phase_2:
	$(GO) test -v -race ./phase_2 -coverprofile ${TEST_COVERAGE_FILE}

test_phase_3:
	$(GO) test -v -race ./phase_3 -coverprofile ${TEST_COVERAGE_FILE}

lint: $(GOLANGCI_LINT)
	$(GOLANGCI_LINT) run --fix --fast --enable-all --disable nosnakecase --disable ifshort --disable gomnd --disable wsl --disable nlreturn --disable scopelint --timeout 120s

$(GOLANGCI_LINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(BIN_DIR) v$(GOLANGCI_LINT_VERSION)

test_with_coverage: test
	$(GO) tool cover -html=${TEST_COVERAGE_FILE}

mockery:
	mockery --name=S3Downloader --with-expecter --dir ./phase_2
	mockery --name=Locker --with-expecter --dir ./phase_3

