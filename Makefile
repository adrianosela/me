NAME:=$(shell basename `git rev-parse --show-toplevel`)
SOURCE_MD_FILE:=resume.md

deploy:
	gcloud app deploy

run: build
	FILENAME=$(SOURCE_MD_FILE) ./$(NAME)

build:
	go build -o $(NAME)

auth:
	gcloud auth login
	gcloud config set project resume-275302
