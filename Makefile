SRC_SERVER = cmd/gophkeeper/main.go
TEST_DIR = ./internal/...


.PHONY: run-server test-server

run-server:
	go run $(SRC_SERVER)

test-server:
	go test -coverprofile c.out -coverpkg ./... $(TEST_DIR) &&  go tool cover -html c.out -o index.html
