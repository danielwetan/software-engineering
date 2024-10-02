install:
	go mod tidy && go mod vendor

run-server:
	go run cmd/http/main.go
	
test:
	go test ./... -count=1

coverage:
	go test -coverprofile=coverage.out ./... -count=1 && go tool cover -html=coverage.out 

mock:
	mockgen -source=internal/store/store.go -destination=internal/store/mockstore/mockstore.go --package=mockstore


