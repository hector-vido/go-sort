# Build the sort binary
build:
	go build ./cmd/sort

# Execute unit tests present in internal/functions/
test:
	cd internal/functions/ && go test -v

# Remove the sort binary
clean:
	rm sort
