all:
	go generate -v internal/client/generate.go
	go build -o af

clean:
	rm ./af
	rm internal/client/oas_*
