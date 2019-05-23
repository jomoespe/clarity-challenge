.PHONY: test listhosts parselog logsupplier

hostnames-target   = listhosts
parselog-target    = parselog
logsupplier-target = logsupplier

all: clean test listhosts parselog logsupplier

clean: 
	@ go clean ./...
	@ rm -f main $(hostnames-target) $(parselog-target) $(logsupplier-target) 

test:
	@ go vet ./...
	@ go test -cover ./...

listhosts:
	@ go build -o $(hostnames-target) cmd/listhosts/main.go

parselog:
	@ go build -o $(parselog-target) cmd/parselog/main.go

logsupplier:
	@ go build -o $(logsupplier-target) cmd/logsupplier/main.go
