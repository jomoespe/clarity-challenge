parselog-target = parselog
parselogd-target = parselogd

build: clean
	@go build -o $(parselog-target) cmd/parselog/main.go
	@go build -o $(parselogd-target) cmd/parselogd/main.go

clean: 
	@rm -f main $(parselog-target) 
