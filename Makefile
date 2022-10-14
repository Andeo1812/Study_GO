.PHONY: run_tests check_coverage

run_tests:
	go test ./... -v -race

check_coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
