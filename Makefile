test_to_file = go test -coverprofile=coverage.out


coverage:
	$(test_to_file)  ./internal/infrastructure/http/  ./internal/utils
	go tool cover -html=coverage.out

mock:
	mockery --dir ./internal --output ./mocks --all

lint:
	golangci-lint run