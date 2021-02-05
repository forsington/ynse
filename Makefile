build:
	go build -o ./bin/ynse

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/ynse``-darwin
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/ynse-amd64-linux
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./bin/ynse-x86-linux
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./bin/ynse-amd64.exe
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./bin/ynse-x86.exe
test:
	gotest `go list ./...`

test-verbose:
	gotest `go list ./...` -coverprofile=coverage.txt
	mkdir -p coverage
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