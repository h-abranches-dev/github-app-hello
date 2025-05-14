install: clean
	mkdir bin
	go mod init github.com/h-abranches-dev/github-app-hello

dev2:
	set -a; \
	. .env; \
	set +a; \
	rm -rf bin/service; \
	go build -v -o bin/service; \
	./bin/service

dev:
	docker container rm -f github-app-hello
	docker image rm -f github-app-hello
	docker build -t github-app-hello .
	docker run --env-file .env -p 8093:8093 github-app-hello

clean:
	rm -f go.*
	rm -rf bin

.PHONY: install dev clean