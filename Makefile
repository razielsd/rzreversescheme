test:
	go test -count=1 ./pkg/...

run:
	go clean -testcache && go run main.go --proto=http
