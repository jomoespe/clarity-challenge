.PHONY: test listhosts parselog

hostnames-target = listhosts
parselog-target  = parselog

all: clean test listhosts parselog

clean: 
	@ go clean ./...
	@ rm -f main $(hostnames-target) $(parselog-target) 

test:
	@ go vet ./...
	@ go test -cover ./...

listhosts:
	@ go build -o $(hostnames-target) cmd/listhosts/main.go

parselog:
	@ go build -o $(parselog-target) cmd/parselog/main.go
