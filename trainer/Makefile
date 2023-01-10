# Dev mode usage only

run:
	cd ./cmd/main && go run . && cd ../..
build:
	go build -o bin/main github.com/ivar-mahhonin/food-delivery-classifier/trainer/cmd/main
run_bin:
	cd bin && ./main && cd ..
remove_model:
	rm -rf ../model_files
run_tests:
	go test -v ./...
run_single_test:
	go test -v ./internal/server -count=1 -run $(test)
