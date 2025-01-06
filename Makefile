SRC_SERVER = cmd/gophkeeper/main.go
TEST_DIR = ./internal/...


.PHONY: run-server test-server

test-run:
	go run ${SRC_SERVER} --dbURL mongodb://localhost:3000

run-server:
	go run $(SRC_SERVER)

run:
	docker-compose up

test-server:
	go test -count=1 -coverprofile c.out -coverpkg ./... $(TEST_DIR) &&  go tool cover -html c.out -o index.html
