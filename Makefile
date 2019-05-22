target = parselogs

build: clean
	@go build -o $(target) cmd/parselogs/main.go

clean: 
	@rm -f main $(target) 
