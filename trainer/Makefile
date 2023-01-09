# Dev mode usage only

run:
	cd ./cmd/main && go run . && cd ../..
run_tests:
	go test -v ./...
run_single_test:
	go test -v ./internal/server -count=1 -run $(test)