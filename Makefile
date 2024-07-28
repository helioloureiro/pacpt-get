SOURCE := main.go list.go utils.go
BINARY := pact-get
pact-get: $(SOURCE)
	go build -o $(BINARY) ./...
