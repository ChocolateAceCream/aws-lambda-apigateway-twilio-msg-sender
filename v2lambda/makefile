.PHONY: build
build:
	GOOS=linux go build -tags lambda.norpc -o ./bootstrap main.go;\
	zip bootstrap.zip bootstrap