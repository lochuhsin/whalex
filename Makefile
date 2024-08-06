.PHONY: build
build:
	go build -o app cmd/*

.PHONY: run-vanilla
run-vanilla: build
	./app