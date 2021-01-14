.PHONY: build clear

build:
	go build -v ./cmd/apiserver
.DEFAULT_GOAL := build

clear:
	rm -rf apiserver
