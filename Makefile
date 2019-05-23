.PHONY: test parselog parselogd

parselog-target = parselog
parselogd-target = parselogd

all: clean test parselog parselogd

clean: 
	@ go clean ./...
	@ rm -f main $(parselog-target) $(parselogd-target) 

dependencies:
	@ go mod tidy

test:
	@ go vet ./...
	@ go test -cover ./...

parselog:
	@ go build -o $(parselog-target) cmd/parselog/main.go

parselogd:
	@ go build -o $(parselogd-target) cmd/parselogd/main.go
