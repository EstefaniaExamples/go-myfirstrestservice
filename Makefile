# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
COVERAGE=$(GOTOOL) cover
GOFORMAT=$(GOCMD) fmt
BUILD_DIR=build
BINARY_NAME=$(BUILD_DIR)/go-myfirstrestservice
APP_PATH="./internal/app"
default: build

build: clean test
	$(GOBUILD) -o $(BINARY_NAME) -v $(APP_PATH)
test:
	$(GOTEST) -v -cover -coverprofile=coverage.out -covermode=atomic $(APP_PATH)/...
coverage: test
	$(COVERAGE) -html=coverage.out
clean:
	$(GOCLEAN) $(APP_PATH)
	rm -rf $(BUILD_DIR)
format:
	$(GOFORMAT) $(APP_PATH)/...
run: build
	./$(BINARY_NAME)
