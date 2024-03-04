BIN_NAME="wisper"
BIN_DIR="bin"

compile:
	@echo "Compiling $(BIN_NAME)"
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go

run: compile
	@./$(BIN_DIR)/$(BIN_NAME)
