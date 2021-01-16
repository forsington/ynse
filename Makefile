build:
	go build -o ./bin/ynse

test: 
	gotest `go list ./...`

test-verbose:
	gotest `go list ./...` -coverprofile=coverage.txt
	mkdir coverage
	go tool cover -html=coverage.txt -o coverage/ynse.html
	open coverage/ynse.html
	rm coverage.txt

vet:
	go vet ./...

lint:
	golint ./...

cilint:
	golangci-lint run --fix ./...


ready: cilint test 