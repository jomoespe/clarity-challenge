.PHONY: test listhosts parselog logsupplier

hostnames-target   = listhosts
parselog-target    = parselog
logsupplier-target = log-generator

all: clean test listhosts parselog logsupplier

clean: 
	@ go clean ./...
	@ rm -f main $(hostnames-target) $(parselog-target) $(logsupplier-target) 

test:
	@ go vet ./...
	@ go test -cover ./...

listhosts:
	@ go build -o $(hostnames-target) -ldflags "-s -w" cmd/listhosts/main.go

parselog:
	@ go build -o $(parselog-target) -ldflags "-s -w" cmd/parselog/main.go

logsupplier:
	@ go build -o $(logsupplier-target) -ldflags "-s -w" cmd/logsupplier/main.go
