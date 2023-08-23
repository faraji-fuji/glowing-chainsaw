shell:=/bin/bash

dev:
	go run .

tidy:
	go mod tidy

push:
	git push origin dev

pr:
	gh pr create --web -B staging