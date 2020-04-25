NAME:=$(shell basename `git rev-parse --show-toplevel`)

deploy:
	gcloud app deploy

run: build
	./$(NAME)

build:
	go build -o $(NAME)
