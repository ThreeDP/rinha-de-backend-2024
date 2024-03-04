NAME				:= 	go_server
PROJECT_PATH		:= 	./

all:
	go build -o $(NAME) $(PROJECT_PATH)*.go

run:
	./$(NAME)

t: unit coverage benchmark

unit:
	go test ./...

coverage:
	go test ./... -race -coverprofile=coverage.txt -covermode=atomic

benchmark:
	go test ./... -bench=.

.PHONY: all run t unit coverage benchmark 