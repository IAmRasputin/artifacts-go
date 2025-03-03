all:
	go build -o af
	go generate -v internal/client/generate.go

clean:
	rm ./af
