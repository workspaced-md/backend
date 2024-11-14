BIN=main
BUILD_DIR=./tmp
CMD_DIR=./cmd

build:
	@go build -o $(BUILD_DIR)/$(BIN) $(CMD_DIR)

run: build
	@./$(BUILD_DIR)/$(BIN)

clean:
	rm -f $(BUILD_DIR)/$(BIN)
