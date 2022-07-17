.PHONY: run
run: # билдить проект с помощью команды make
	go run ./cmd/apiserver

.DEFAULT_GOAL := run