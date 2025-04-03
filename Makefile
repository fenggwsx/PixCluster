BUILD_DIR := build
APPS := kmeans summarize text2image
WEB_DIR := web
WEB_BUILD_DIR := $(WEB_DIR)/out

APP_BUILD_TARGETS = $(addprefix build-,$(APPS))
APP_CLEAN_TARGETS = $(addprefix clean-,$(APPS))

GOOS        ?= linux
GOARCH      ?= amd64
CGO_ENABLED ?= 0
GOFLAGS     :=
TAGS        :=
LDFLAGS     := -w -s

GIT_COMMIT = $(shell git rev-parse HEAD)
GIT_SHA    = $(shell git rev-parse --short HEAD)
GIT_TAG    = $(shell git describe --tags --abbrev=0 --exact-match 2>/dev/null)
GIT_DIRTY  = $(shell test -n "`git status --porcelain`" && echo "dirty" || echo "clean")

.PHONY: build
build: build-apps build-web

.PHONY: build-apps
build-apps: $(APP_BUILD_TARGETS)
	@echo "Building apps ..."

.PHONY: $(APP_BUILD_TARGETS)
$(APP_BUILD_TARGETS): build-%:
	@echo "Building $* ..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
		go build $(GOFLAGS) -trimpath \
		-tags '$(TAGS)' \
		-ldflags '$(LDFLAGS)' \
		-o cmd/$*/main \
		cmd/$*/main.go
	@mkdir -p $(BUILD_DIR)
	@echo "Packaging $* to $(BUILD_DIR)/$*.zip ..."
	@zip -jq $(BUILD_DIR)/$*.zip cmd/$*/main

.PHONY: build-web
build-web:
	@echo "Building web ..."
	$(MAKE) -C $(WEB_DIR) build
	@mkdir -p $(BUILD_DIR)
	@echo "Packaging web to $(BUILD_DIR)/web.zip ..."
	@cd $(WEB_DIR) && zip -qr ../$(BUILD_DIR)/web.zip out nginx.conf

.PHONY: clean
clean: clean-apps clean-web 
	@echo "Cleaning $(BUILD_DIR) ..."
	rm -rf $(BUILD_DIR)

.PHONY: clean-apps
clean-apps: $(APP_CLEAN_TARGETS)

.PHONY: $(APP_CLEAN_TARGETS)
$(APP_CLEAN_TARGETS): clean-%:
	@echo "Cleaning $* ..."
	rm -f cmd/$*/main
	rm -f $(BUILD_DIR)/$*.zip

.PHONY: clean-web
clean-web:
	@echo "Cleaning web ..."
	$(MAKE) -C $(WEB_DIR) clean

.PHONY: info
info:
	@echo "Git Tag:           ${GIT_TAG}"
	@echo "Git Commit:        ${GIT_COMMIT}"
	@echo "Git Tree State:    ${GIT_DIRTY}"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build      - Build everything (default)"
	@echo "  build-web  - Build frontend only"
	@echo "  build-apps - Build all Go apps"
	@echo "  clean      - Remove all build artifacts"
	@echo "  clean-web  - Remove frontend build only"
	@echo "  clean-apps - Remove all Go apps"
	@echo "  info       - Show git information"
	@echo "  help       - Show this help message"
