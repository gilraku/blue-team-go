TOOLS := hash-checker b64 cidr-calc dns-lookup jwt-decode ioc-extract log-parser port-scanner tls-check pwcheck http-inspect report-gen
BIN_DIR := bin

.PHONY: all build clean vet

all: build

build:
	@mkdir -p $(BIN_DIR)
	@for tool in $(TOOLS); do \
		echo "Building $$tool..."; \
		go build -o $(BIN_DIR)/$$tool ./cmd/$$tool; \
	done

vet:
	go vet ./...

clean:
	rm -rf $(BIN_DIR)

install:
	@for tool in $(TOOLS); do \
		go install ./cmd/$$tool; \
	done
