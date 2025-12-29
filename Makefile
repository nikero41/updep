# Strip debug info
GO_FLAGS += -ldflags="-w -s"
# Avoid embedding the build path in the executable for more reproducible builds
GO_FLAGS += -trimpath

PROJECT_FILES = cmd/$(APP_NAME)/*.go pkg/*/*.go internal/*/*.go go.mod go.sum
GENERATED_FILES = debug.log

OS = darwin linux win32
ARCH = arm64 x64

PLATFORMS = $(foreach os, $(OS), $(foreach arch, $(ARCH), $(os)-$(arch)))
PLATFORM_DIRS = $(foreach platform, $(PLATFORMS), npm/platforms/$(platform))
PLATFORM_FILES = $(foreach platform, $(PLATFORMS), npm/platforms/$(platform)/bin/$(call binaryname,$(platform)))

APP_NAME = updep
binaryname = $(if $(findstring win32, $(1)),$(APP_NAME).exe,$(APP_NAME))

os = $(subst win32,windows,$(firstword $(subst -, ,$(1))))
arch = $(subst x64,amd64,$(lastword $(subst -, ,$(1))))

.PHONY: all
all: $(PLATFORM_FILES)

.PHONY: format
format:
	@go fmt ./...

npm/platforms/%/bin/$(APP_NAME) npm/platforms/%/bin/$(APP_NAME).exe: $(PROJECT_FILES)
	$(info Building $*)
	@GOOS=$(call os,$*) GOARCH=$(call arch,$*) go build $(GO_FLAGS) -o $@ ./cmd/$(APP_NAME)

.PHONY: run
run:
	@DEBUG=1 go run ./cmd/$(APP_NAME)

.PHONY: dev-install
dev-install:
	@go install ./cmd/$(APP_NAME)

.PHONY: install
install: $(PROJECT_FILES)
	@go install $(GO_FLAGS) ./cmd/$(APP_NAME)

.PHONY: check
check: 
	echo todo

.PHONY: clean
clean:
	@rm -rf $(PLATFORM_FILES) $(GENERATED_FILES)
