.PHONY: build
build: tidy
	@go build -v -ldflags "$(GO_LDFLAGS)" samples/myapp.go

.PHONY: format
format:
	@gofmt -s -w ./

.PHONY: add-copyright
add-copyright:
	@addlicense -v -f $(ROOT_DIR)/scripts/boilerplate.txt $(ROOT_DIR) --skip-dirs=third_party,vendor,$(OUTPUT_DIR)

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: clean
clean:
	@-rm -vrf $(OUTPUT_DIR)