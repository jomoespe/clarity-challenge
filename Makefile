parselog-target = parselog
parselogd-target = parselogd

build: clean
	@go build -o $(parselog-target) cmd/parselog/main.go

clean: 
	@rm -f main $(parselog-target) 
