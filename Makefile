# Пути
PROTO_DIR := proto
OUT_DIR := internal/pb

# Компилятор
PROTOC := protoc

# Все proto файлы
PROTOS := $(shell find $(PROTO_DIR) -name "*.proto")

# Цель по умолчанию
all: install generate

# Установка необходимых плагинов
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Генерация Go protobuf и gRPC кода
generate:
	@echo "Creating output directory: $(OUT_DIR)"
	@mkdir -p $(OUT_DIR)
	@echo "Generating Go protobuf and gRPC code..."
	@$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTOS)
	@echo "Generation completed! Files are in $(OUT_DIR)"
