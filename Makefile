NAME:=$(shell basename `git rev-parse --show-toplevel`)

deploy:
	gcloud app deploy

run: build
	./$(NAME)

build:
	go build -o $(NAME)

auth:
	gcloud auth login
	gcloud config set project resume-275302
